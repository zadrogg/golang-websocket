package handlers

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

func OnOpen(write http.ResponseWriter, request *http.Request) {

}

func OnMessage(write http.ResponseWriter, request *http.Request) {

}

func Echo(ws *websocket.Conn) {
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
