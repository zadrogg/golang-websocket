package routes

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"golang.org/x/net/websocket"
	"net/http"
	ws "websocket/handlers"
)

func ApiRoutes() {
	http.Handle("/notify", websocket.Handler(ws.Echo))
	http.HandleFunc("/docs/", httpSwagger.WrapHandler)
}
