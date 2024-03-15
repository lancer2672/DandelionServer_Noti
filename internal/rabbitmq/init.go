package rabbitmq

import (
	"context"
	"log"
	"time"

	"github.com/lancer2672/DandelionServer_Noti/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func createQueue(ch *amqp.Channel) {
	q, err := ch.QueueDeclare(
		"hello", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	utils.FailOnError(err, "Failed to open a channel")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	utils.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)

}
func createChannel(conn *amqp.Connection) {
	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()
}
func ConnectRabbitMQ(connString string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(connString)
	if err != nil {
		return nil, err
	}
	createChannel(conn)
	return conn, nil
}
