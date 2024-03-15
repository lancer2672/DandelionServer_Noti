package main

import (
	"log"

	"github.com/lancer2672/DandelionServer_Noti/internal/rabbitmq"
	"github.com/lancer2672/DandelionServer_Noti/utils"
)

func main() {

	config, err := utils.LoadConfig(".")
	utils.FailOnError(err, "cannot load config file")
	conn, err := rabbitmq.ConnectRabbitMQ(config.RABBITMQ_CONN)
	utils.FailOnError(err, "cannot connect to rabbitmq")
	_ = conn
	defer func() {
		conn.Close()
		log.Println("RabbitMQ connection closed")
	}()
}
