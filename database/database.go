package database

import (
	"log"
	"context"
	
	"cloud.google.com/go/firestore"

	"github.com/changdaozheng/headhome-backend/firebase_app"
)

var FBCtx context.Context
var Client *firestore.Client

func init(){
	var err error
	FBCtx = context.Background()
	Client, err = firebase_app.App.Firestore(FBCtx)
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
	Client.Close()
}