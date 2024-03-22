package main

import (
	"log"

	"github.com/lancer2672/DandelionServer_Noti/db"
	"github.com/lancer2672/DandelionServer_Noti/db/repository"
	"github.com/lancer2672/DandelionServer_Noti/server"
	"github.com/lancer2672/DandelionServer_Noti/services"
	"github.com/lancer2672/DandelionServer_Noti/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}
	conn := db.Init(config.DB_SOURCE)
	repo := repository.NewNotificationRepo(conn)
	notiService := services.NewNotificationService(repo)
	server.ConfigServer(config, notiService)
	server.StartServer(config.SERVER_ADDRESS)

}
