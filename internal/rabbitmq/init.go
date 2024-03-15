package rabbitmq

import (
	"github.com/lancer2672/DandelionServer_Noti/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func CreateQueue(ch *amqp.Channel, name string) amqp.Queue {
	q, err := ch.QueueDeclare(
		name,  // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	utils.FailOnError(err, "Failed to open a channel")
	return q
}
func CreateChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	return ch
}
func ConnectRabbitMQ(connString string) *amqp.Connection {
	conn, err := amqp.Dial(connString)
	utils.FailOnError(err, "Failed to connect to rabbitmq")
	return conn
}
