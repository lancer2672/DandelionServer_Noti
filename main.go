package main

import (
	"log"

	"github.com/lancer2672/DandelionServer_Noti/server"
	"github.com/lancer2672/DandelionServer_Noti/services"
	"github.com/lancer2672/DandelionServer_Noti/utils"
	_ "github.com/swaggo/http-swagger"
)

// @title Swagger Dandelion Notification API
// @version 1.0
// @description This is Dandelion Notification server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config", err)
	}
	notiService := services.GetService()
	mux := server.ConfigServer(config, notiService)
	server.StartServer(config.SERVER_ADDRESS, mux)
}
