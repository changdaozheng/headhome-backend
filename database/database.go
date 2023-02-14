package database

import (
	"fmt"
	"log"
	"context"
	
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"cloud.google.com/go/firestore"
)

var Ctx context.Context
var Client *firestore.Client

var Members *firestore.CollectionRef

func InitDB(){
	Ctx = context.Background()
	conf := &firebase.Config{ProjectID: "gsc23-12e94"}
	opt := option.WithCredentialsFile("./database/firebase_config.json")
	app, err := firebase.NewApp(Ctx, conf ,opt)
	if err != nil {
	  log.Fatalln(err)
	}
	
	Client, err = app.Firestore(Ctx)
	if err != nil {
	  log.Fatalln(err)
	}
	
	Members = Client.Collection("members")
}

func CloseDB(){
	fmt.Print(Client)
	Client.Close()
}