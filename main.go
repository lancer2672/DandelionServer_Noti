package main

import (
	"log"

	"github.com/lancer2672/DandelionServer_Noti/internal/rabbitmq"
	"github.com/lancer2672/DandelionServer_Noti/utils"
)

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config file", err)
	}
	conn, err := rabbitmq.ConnectRabbitMQ(config.RABBITMQ_CONN)
	if err != nil {
		log.Fatal("cannot connect to rabbitmq", err)
	}
	_ = conn
}
