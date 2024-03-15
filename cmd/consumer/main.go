// cmd/consumer/main.go
package main

import (
	"log"

	"github.com/lancer2672/DandelionServer_Noti/internal/rabbitmq"
	"github.com/lancer2672/DandelionServer_Noti/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	utils.FailOnError(err, "cannot load config file")
	conn := rabbitmq.ConnectRabbitMQ(config.RABBITMQ_CONN)

	ch, err := conn.Channel()
	utils.FailOnError(err, "failed to open a channel")

	q := rabbitmq.CreateQueue(ch, "hello")

	consumer := rabbitmq.NewConsumer(ch, q)
	consumer.Consume()
	defer func() {
		ch.Close()
		conn.Close()
		log.Println("RabbitMQ connection closed")
	}()
}