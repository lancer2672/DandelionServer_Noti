// cmd/consumer/main.go
package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/lancer2672/DandelionServer_Noti/constants"
	"github.com/lancer2672/DandelionServer_Noti/db"
	"github.com/lancer2672/DandelionServer_Noti/internal/firebase"
	"github.com/lancer2672/DandelionServer_Noti/internal/rabbitmq"
	"github.com/lancer2672/DandelionServer_Noti/utils"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	config, err := utils.LoadConfig(".")
	utils.FailOnError(err, "cannot load config file")

	conn := rabbitmq.ConnectRabbitMQ(config.RABBITMQ_CONN)
	db.Init(config.DB_SOURCE)
	firebase.InitializeApp()

	ch, err := conn.Channel()
	utils.FailOnError(err, "failed to open a channel")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	utils.FailOnError(err, "failed to set QoS")

	// Tạo và cấu hình queue cho notification
	queue_TTL := createNotificationQueue(ch)

	// Tạo và cấu hình queue cho DLX
	queue_DLX := createDLXQueue(ch)

	// Khởi tạo consumer cho notification queue và truyền callback
	consumer := rabbitmq.NewConsumer(ch, queue_TTL)
	go consumer.Consume(notificationCallback)

	// Khởi tạo consumer cho DLX queue và truyền callback
	dlx_consumer := rabbitmq.NewConsumer(ch, queue_DLX)
	go dlx_consumer.Consume(dlxCallback)

	// Keep the application running
	<-make(chan bool)
}

func createNotificationQueue(ch *amqp091.Channel) amqp091.Queue {
	// Tạo queue với TTL
	queue_TTL := rabbitmq.CreateQueueWithTTL(ch, constants.NOTI_QUEUE_NAME, constants.TTL_VALUE, constants.DLX_EX_NAME, constants.DLX_ROUTING_KEY)

	// Bind exchange DLX với queue
	err := ch.QueueBind(
		queue_TTL.Name,            // queue name
		constants.DLX_ROUTING_KEY, // routing key
		constants.DLX_EX_NAME,     // exchange
		false,
		nil)
	utils.FailOnError(err, "bind exchange dlx failed")

	return queue_TTL
}

func createDLXQueue(ch *amqp091.Channel) amqp091.Queue {
	// Tạo DLX queue
	queue_DLX := rabbitmq.CreateQueueWithDLX(ch, constants.DLX_QUEUE_NAME)

	return queue_DLX
}

func notificationCallback(d amqp091.Delivery, queueName string) {
	// Xử lý thông điệp từ notification queue
	log.Printf("Received message from notification queue (%s): %s", queueName, d.Body)
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	if err := d.Ack(false); err != nil {
		log.Printf("Failed to acknowledge message: %s", err)
	}
	// Thực hiện xử lý dữ liệu và gửi đến Firebase, DB, vv.
}

func dlxCallback(d amqp091.Delivery, queueName string) {
	// Xử lý thông điệp từ DLX queue
	log.Printf("Received message from DLX queue (%s): %s", queueName, d.Body)
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	if err := d.Ack(false); err != nil {
		log.Printf("Failed to acknowledge message: %s", err)
	}
	// Thực hiện các thao tác xử lý khác cho DLX, ví dụ như ghi log, cảnh báo, vv.
}
