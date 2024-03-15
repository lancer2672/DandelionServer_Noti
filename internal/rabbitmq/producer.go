package rabbitmq

import (
	"context"
	"log"
	"time"

	"github.com/lancer2672/DandelionServer_Noti/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func NewProducer(ch *amqp.Channel, q amqp.Queue) *Producer {
	return &Producer{
		Channel: ch,
		Queue:   q,
	}
}

func (p *Producer) Publish(body string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := p.Channel.PublishWithContext(ctx,
		"",           // exchange
		p.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	utils.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
