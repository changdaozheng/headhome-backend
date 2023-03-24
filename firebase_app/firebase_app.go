package firebase_app

import (
    "os"
	"log"
	"context"

    "firebase.google.com/go"
    "google.golang.org/api/option"
    "github.com/joho/godotenv"
)

var App *firebase.App

func init() {

    err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := &firebase.Config{ProjectID: "gsc23-12e94"}
	opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_ADMIN_PRIVATE_KEY")))
	App, err = firebase.NewApp(context.Background(), conf ,opt)
	if err != nil {
	  	log.Fatalln(err)
	}
}