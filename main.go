package main

import (
	"log"

	"github.com/lancer2672/DandelionServer_Noti/server"
	"github.com/lancer2672/DandelionServer_Noti/services"
	"github.com/lancer2672/DandelionServer_Noti/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}
	notiService := services.GetService()
	server.ConfigServer(config, notiService)
	server.StartServer(config.SERVER_ADDRESS)
}
