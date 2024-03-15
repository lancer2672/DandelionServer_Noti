// cmd/consumer/main.go
package main

import (
	"log"

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

	err = ch.ExchangeDeclare(
		constants.DLX_EX_NAME, // name
		"fanout",              // type
		true,                  // durable
		false,                 // auto-deleted
		false,                 // internal
		false,                 // no-wait
		nil,                   // arguments
	)
	utils.FailOnError(err, "create exchange failed")

	//  notification queue with TTL set, would be bind with queue_DLX through "dlx_exchange"
	queue_TTL := rabbitmq.CreateQueueWithTTL(ch, constants.NOTI_QUEUE_NAME, constants.TTL_VALUE, constants.DLX_EX_NAME)
	// dlx queue "dlx_exchange"
	queue_DLX := rabbitmq.CreateQueueWithDLX(ch, constants.DLX_QUEUE_NAME, constants.DLX_EX_NAME)
	err = ch.QueueBind(
		queue_DLX.Name, // queue name
		// constants.DLX_ROUTING_KEY, // routing key
		"",                    // routing key
		constants.DLX_EX_NAME, // exchange
		false,
		nil)
	utils.FailOnError(err, "bind exchange dlx failed")

	producer_TTL := rabbitmq.NewProducer(ch, queue_TTL)

	producer_DLX := rabbitmq.NewProducer(ch, queue_DLX)
	producer_TTL.Publish("error")
	_ = producer_DLX
	defer func() {
		ch.Close()
		conn.Close()
		log.Println("RabbitMQ connection closed")
	}()
}
