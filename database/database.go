package database

import (
	"os"
	"fmt"
	"log"
	"context"
	
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"cloud.google.com/go/firestore"

	"github.com/joho/godotenv"
)

var FBCtx context.Context
var Client *firestore.Client


func InitDB(){

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	FBCtx = context.Background()
	conf := &firebase.Config{ProjectID: "gsc23-12e94"}
	opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_ADMIN_PRIVATE_KEY")))
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