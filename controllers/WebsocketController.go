package controllers

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"websocket/handlers"
)

func WsToken(w http.ResponseWriter, r *http.Request) {

}

func ServerCreate(ws *websocket.Conn) {
	hub := handlers.NewHub()
	handlers.Create(ws, hub)
}

func Connect(ws *websocket.Conn) {
	var err error

	fmt.Println(ws)

	for {
		var reply string

		fmt.Println(reply)

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}
