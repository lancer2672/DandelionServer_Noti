// consumer.go
package rabbitmq

import (
	"log"

	"github.com/lancer2672/DandelionServer_Noti/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func NewConsumer(ch *amqp.Channel, q amqp.Queue) *Consumer {
	return &Consumer{
		Channel: ch,
		Queue:   q,
	}
}

type MessageHandler func(d amqp.Delivery, queueName string)

func (c *Consumer) Consume(callback MessageHandler) {
	msgs, err := c.Channel.Consume(
		c.Queue.Name, // queue
		"",           // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			go callback(d, c.Queue.Name)
		}
	}()
	log.Printf(c.Queue.Name + " is waiting for messages...")
	// Keep the consumer alive
	select {}
}
