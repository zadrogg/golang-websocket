package routes

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"proxy-service/controllers"
)

func ApiRoutes() {
	http.HandleFunc("/", controllers.ProxyRequest)
	http.HandleFunc("/docs/", httpSwagger.WrapHandler)
}
