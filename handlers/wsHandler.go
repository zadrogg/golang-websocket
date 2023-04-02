package handlers

import (
	"golang.org/x/net/websocket"
	"gorm.io/gorm"
	"websocket/models"
)

type Message struct {
	Text string `json:"message"`
}
type Hub struct {
	clients          map[string]*websocket.Conn
	addClientChan    chan *websocket.Conn
	removeClientChan chan *websocket.Conn
	broadcastChan    chan Message
	db               *gorm.DB
}

func NewHub() *Hub {
	return &Hub{
		clients:          make(map[string]*websocket.Conn),
		addClientChan:    make(chan *websocket.Conn),
		removeClientChan: make(chan *websocket.Conn),
		broadcastChan:    make(chan Message),
	}
}

func Create(ws *websocket.Conn, h *Hub) {
	go h.run()

	h.addClientChan <- ws

	for {
		var m Message
		err := websocket.JSON.Receive(ws, &m)
		if err != nil {
			h.broadcastChan <- Message{err.Error()}
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
			h.broadcast(m)
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

func (h *Hub) broadcast(m Message) {

	models.OnMessage(h.db, "")
}
