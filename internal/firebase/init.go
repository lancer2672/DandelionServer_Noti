package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/lancer2672/DandelionServer_Noti/utils"
	"google.golang.org/api/option"
)

var app *firebase.App
var messagingClient *messaging.Client

func InitializeApp() {
	var err error
	opt := option.WithCredentialsFile("service-key.json")

	app, err = firebase.NewApp(context.Background(), nil, opt)
	utils.FailOnError(err, "error initializing app")

	messagingClient, err = app.Messaging(context.Background())
	utils.FailOnError(err, "error creating messaging client")

	log.Println("Connected to firebase")
}

func SendNotification() {
	messagingClient.Send(context.Background(), &messaging.Message{

		Notification: &messaging.Notification{
			Title: "Title",
			Body:  "Body",
		},
		Token: "hehe",
	})
}
