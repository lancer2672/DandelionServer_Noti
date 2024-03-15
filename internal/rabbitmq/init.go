package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectRabbitMQ(connString string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(connString)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
