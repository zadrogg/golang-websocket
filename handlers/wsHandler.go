package handlers

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
	"gorm.io/gorm"
	"websocket/models"
)

type Message struct {
	Message string `json:"message"`
	UserId  string `json:"user_id"`
}

type Hub struct {
	clients          map[string]*websocket.Conn
	addClientChan    chan *websocket.Conn
	removeClientChan chan *websocket.Conn
	broadcastChan    chan Message
	db               *gorm.DB
}

var WsChannel *Hub

func NewHub() *Hub {
	WsChannel = &Hub{
		clients:          make(map[string]*websocket.Conn),
		addClientChan:    make(chan *websocket.Conn),
		removeClientChan: make(chan *websocket.Conn),
		broadcastChan:    make(chan Message),
	}

	return WsChannel
}

func Create(ws *websocket.Conn, h *Hub) {
	go h.run()

	h.addClientChan <- ws

	for {
		var m Message
		err := websocket.JSON.Receive(ws, &m)
		if err != nil {
			h.broadcastChan <- Message{Message: err.Error()}
			h.removeClient(ws)
			return
		}
		h.broadcastChan <- m
	}
}

func (h *Hub) run() {
	for {
		select {
		case conn := <-h.addClientChan:
			h.addClient(conn)
		case conn := <-h.removeClientChan:
			h.removeClient(conn)
		case m := <-h.broadcastChan:
			h.Broadcast(m)
		}
	}
}

func (h *Hub) addClient(conn *websocket.Conn) {
	request := conn.Request().URL.Query()

	token := request.Get("token")
	userId := request.Get("identifier")
	socket := conn.RemoteAddr().String()

	h.clients[socket] = conn
	models.OnOpen(h.db, token, userId, socket)
}

func (h *Hub) removeClient(conn *websocket.Conn) {
	socket := conn.LocalAddr().String()
	delete(h.clients, socket)
	models.OnClose(h.db, socket)
}

func (h *Hub) Broadcast(m Message) {
	socket := models.OnMessage(h.db, m.UserId)
	conn := h.clients[socket]

	err := websocket.JSON.Send(conn, m.Message)
	if err != nil {
		log.Info("Error broadcasting message: ", err)
		return
	}
}
