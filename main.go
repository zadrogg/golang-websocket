package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"net/http"
	"websocket/config"
	"websocket/handlers"
	"websocket/routes"
)

func init() {
	//load values from .env
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

// @title           Документация сервиса уведомлений
// @version         1.0
// @description     Сервис уведомлений

// @contact.name   API Support

// @host      localhost
// @BasePath  /
// @schemes http
func main() {
	conf := config.GetConfig()

	InitDb(conf)

	log.Info("run server")

	// роуты
	routes.ApiRoutes()

	// запускаем сервер и пишем логи о ошибках
	handlers.Error(http.ListenAndServe(conf.Server.Url, nil))
}
