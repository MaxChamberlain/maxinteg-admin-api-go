package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var app *firebase.App
var err error

func InitFirebase() {
	context := context.Background()

	instance := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_CREDENTIALS")))
	app, err = firebase.NewApp(context, nil, instance)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Firebase Initialized")
	fmt.Println("vars: " + strings.Join(os.Environ(), "\n"))
}

func GetFirebase() *firebase.App {
	return app
}
