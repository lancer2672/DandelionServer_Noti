package rabbitmq

import (
	"github.com/lancer2672/DandelionServer_Noti/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func CreateQueue(ch *amqp.Channel, queueName string) amqp.Queue {
	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	utils.FailOnError(err, "Failed to open a channel")
	return q
}

func CreateQueueWithTTL(channel *amqp.Channel, queueName string, msgTTL int64, dlxName string, routingKey string) amqp.Queue {
	args := amqp.Table{}
	args["x-message-ttl"] = msgTTL
	args["x-dead-letter-exchange"] = dlxName // associate a DLX to the queue
	args["x-dead-letter-routing-key"] = routingKey
	q, err := channel.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		args,      // arguments
	)
	utils.FailOnError(err, "Failed to create a queue")
	return q
}

func CreateQueueWithDLX(channel *amqp.Channel, queueName string) amqp.Queue {
	args := amqp.Table{}
	// args["x-death"] = amqp.Table{
	// 	"count": 3,
	// }
	q, err := channel.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		args,      // arguments
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
