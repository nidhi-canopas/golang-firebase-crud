package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

func main() {

	ctx := context.Background()

	// configure database URL
	conf := &firebase.Config{
		DatabaseURL: "https://fir-realtime-db-demo-xxxxx-default-rtdb.asia-southeast1.firebasedatabase.app",
	}

	// fetch service account key
	opt := option.WithCredentialsFile("fir-realtime-db-demo-xxxxx-firebase-adminsdk-2zjoi-9b40ce1e42.json")

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("error in creating firebase DB client: ", err)
	}

	// add/update data to firebase DB
	SaveDataToFirebaseDB(client)

	// retrieve data from firebase DB
	GetDataFromFirebaseDB(client)

	// delete data from firebase DB
	DeleteDataFromFirebaseDB(client)
}

func SaveDataToFirebaseDB(client *db.Client) {
	// create ref at path user_scores/:userId
	ref := client.NewRef("user_scores/" + fmt.Sprint(1))

	if err := ref.Set(context.TODO(), map[string]interface{}{"score": 40}); err != nil {
		log.Fatal(err)
	}

	fmt.Println("score added/updated successfully!")
}

func GetDataFromFirebaseDB(client *db.Client) {
	type UserScore struct {
		Score int `json:"score"`
	}

	// get database reference to user score
	ref := client.NewRef("user_scores/1")

	// read from user_scores using ref
	var s UserScore
	if err := ref.Get(context.TODO(), &s); err != nil {
		log.Fatalln("error in reading from firebase DB: ", err)
	}
	fmt.Println("retrieved user's score is: ", s.Score)
}

func DeleteDataFromFirebaseDB(client *db.Client) {
	ref := client.NewRef("user_scores/1")

	if err := ref.Delete(context.TODO()); err != nil {
		log.Fatalln("error in deleting ref: ", err)
	}
	fmt.Println("user's score deleted successfully:)")
}
