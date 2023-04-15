package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"golang.org/x/net/websocket"
	"net/http"
	"time"
	"websocket/Validator"
	"websocket/handlers"
	"websocket/models"
	"websocket/requests"
)

type ResponseWsToken struct {
	Token string `json:"token"`
}

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

func SendMessage(w http.ResponseWriter, r *http.Request) {
	Validator.RuleMethod(w, r, http.MethodPost)

	var input requests.SendMessageRequest
	params := &input

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		handlers.JSONError(w, err, http.StatusBadRequest)
		return
	}

	handlers.WsChannel.Broadcast(
		handlers.Message{
			Message: params.Message,
			UserId:  params.UserIdentifier,
		})
}
