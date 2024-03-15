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

func NewConsumer(ch *amqp.Channel) *Consumer {
	q, err := ch.QueueDeclare(
		"hello", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	return &Consumer{
		Channel: ch,
		Queue:   q,
	}
}

func (c *Consumer) Consume() {
	msgs, err := c.Channel.Consume(
		c.Queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
