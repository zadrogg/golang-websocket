package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"time"
	"websocket/handlers"
	"websocket/models"
)

type ResponseWsToken struct {
	Token string `json:"token"`
}

//var db *gorm.DB

func WsToken(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var input models.UserConnection
		userParams := &input

		err := json.NewDecoder(r.Body).Decode(userParams)
		if err != nil {
			handlers.JSONError(w, err, http.StatusBadRequest)
			return
		}

		token := md5.Sum([]byte(userParams.UserIdentifier + time.Now().String()))

		createRow := models.DB.Create(&models.UserConnection{
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			Token:          hex.EncodeToString(token[:]),
			UserIdentifier: input.UserIdentifier,
		})

		if createRow.RowsAffected < 1 {
			handlers.JSONError(w, createRow.Error, http.StatusInternalServerError)
			return
		}

		handlers.JSONSetHeaders(w, http.StatusOK)
		_ = json.NewEncoder(w).Encode(&ResponseWsToken{Token: hex.EncodeToString(token[:])})
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	return
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
