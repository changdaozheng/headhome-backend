package fcm

import (
	"log"
	"context"

    "firebase.google.com/go/messaging"

	"github.com/changdaozheng/headhome-backend/firebase_app"
)

var FCMContext context.Context
var FCMClient *messaging.Client

func init(){
	var err error

	FCMContext = context.Background()
	FCMClient, err = firebase_app.App.Messaging(FCMContext)
	if err != nil {
		log.Fatalln(err)
	}
}

func TopicSend(body map[string]string, topic string) (error){
	message := &messaging.Message{
        Data: body, 
        Topic: topic,
	}

	// Send a message to the devices subscribed to the provided topic.
	_, err := FCMClient.Send(FCMContext, message)
	if err != nil {
		return err
	}

	return nil
}