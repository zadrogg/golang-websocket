package routes

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"golang.org/x/net/websocket"
	"net/http"
	"websocket/controllers"
)

func ApiRoutes() {
	http.Handle("/notify", websocket.Handler(controllers.ServerCreate))

	http.HandleFunc("/send", controllers.SendMessage)
	http.HandleFunc("/ws-token", controllers.WsToken)
	http.HandleFunc("/docs/", httpSwagger.WrapHandler)
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
}
