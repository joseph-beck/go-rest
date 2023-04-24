package firestore

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func newConn(path string) *firestore.Client {
	file, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("json file exists")
	// defer the closing of our jsonFile so that we can parse it later on
	defer file.Close()

	ctx := context.Background()
	sa := option.WithCredentialsFile(path)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}
