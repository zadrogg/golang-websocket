package routes

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	ws "websocket/handlers"
)

func ApiRoutes() {
	http.HandleFunc("/", ws.OnMessage)
	http.HandleFunc("/docs/", httpSwagger.WrapHandler)
}
