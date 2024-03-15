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
			// // Introduce an artificial error
			// if string(d.Body) == "error" {
			// 	log.Printf("Failed to process message: %s", d.Body)
			// if err := d.Nack(false, false); err != nil {
			// 		log.Printf("Failed to nack message: %s", err)
			// 	}
			// 	continue
			// }

			if err := d.Ack(false); err != nil {
				log.Printf("Failed to acknowledge message: %s", err)
			}
		}
	}()

	go func() {
		for msg := range c.Channel.NotifyReturn(make(chan amqp.Return)) {
			log.Printf("Failed to deliver message: %s", msg.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
