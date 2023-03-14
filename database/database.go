package database

import (
	"fmt"
	"log"
	"context"
	
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"cloud.google.com/go/firestore"
)

var FBCtx context.Context
var Client *firestore.Client


func InitDB(){
	FBCtx = context.Background()
	conf := &firebase.Config{ProjectID: "gsc23-12e94"}
	opt := option.WithCredentialsFile("./database/firebase_config.json")
	app, err := firebase.NewApp(FBCtx, conf ,opt)
	if err != nil {
	  log.Fatalln(err)
	}
	
	Client, err = app.Firestore(FBCtx)
	if err != nil {
	  log.Fatalln(err)
	}

	//Init collections 
	InitVolunteers()
	InitCareGiver()
	InitCareReceiver()
	InitSosLog()
	InitTravelLog()
}

func CloseDB(){
	fmt.Print(Client)
	Client.Close()
}