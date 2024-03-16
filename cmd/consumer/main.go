// cmd/consumer/main.go
package main

import (
	"github.com/lancer2672/DandelionServer_Noti/constants"
	"github.com/lancer2672/DandelionServer_Noti/internal/rabbitmq"
	"github.com/lancer2672/DandelionServer_Noti/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	utils.FailOnError(err, "cannot load config file")
	conn := rabbitmq.ConnectRabbitMQ(config.RABBITMQ_CONN)

	ch, err := conn.Channel()
	utils.FailOnError(err, "failed to open a channel")

	queue_TTL := rabbitmq.CreateQueueWithTTL(ch, constants.NOTI_QUEUE_NAME, constants.TTL_VALUE, constants.DLX_EX_NAME, constants.DLX_ROUTING_KEY)

	queue_DLX := rabbitmq.CreateQueue(ch, constants.DLX_QUEUE_NAME)
	err = ch.QueueBind(
		queue_DLX.Name,            // queue name
		constants.DLX_ROUTING_KEY, // routing key
		constants.DLX_EX_NAME,     // exchange
		false,
		nil)
	utils.FailOnError(err, "bind exchange dlx failed")

	consumer := rabbitmq.NewConsumer(ch, queue_TTL)
	dlx_consumer := rabbitmq.NewConsumer(ch, queue_DLX)

	forever := make(chan bool)
	go consumer.Consume()
	go dlx_consumer.Consume()
	<-forever
}
