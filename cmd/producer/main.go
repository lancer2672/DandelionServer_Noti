// cmd/consumer/main.go
package main

import (
	"fmt"
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
		"direct",              // type
		true,                  // durable
		false,                 // auto-deleted
		false,                 // internal
		false,                 // no-wait
		nil,                   // arguments
	)
	utils.FailOnError(err, "create exchange failed")
	err = ch.ExchangeDeclare(
		constants.NOTIFICATION_EX_NAME, // name
		"direct",                       // type
		true,                           // durable
		false,                          // auto-deleted
		false,                          // internal
		false,                          // no-wait
		nil,                            // arguments
	)
	utils.FailOnError(err, "create exchange failed")

	//  notification queue with TTL set, would be bind with queue_DLX through "dlx_exchange"
	queue_TTL := rabbitmq.CreateQueueWithTTL(ch, constants.NOTI_QUEUE_NAME, constants.TTL_VALUE, constants.DLX_EX_NAME, constants.DLX_ROUTING_KEY)
	err = ch.QueueBind(
		queue_TTL.Name,                     // queue name
		constants.NOTIFICATION_ROUTING_KEY, // routing key
		constants.NOTIFICATION_EX_NAME,     // exchange
		false,
		nil)
	utils.FailOnError(err, "bind exchange dlx failed")

	producer_TTL := rabbitmq.NewProducer(ch, queue_TTL)

	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Message %d", i+1)
		producer_TTL.Publish(message, constants.NOTIFICATION_EX_NAME, constants.NOTIFICATION_ROUTING_KEY)
	}
	defer func() {
		ch.Close()
		conn.Close()
		log.Println("RabbitMQ connection closed")
	}()
}
