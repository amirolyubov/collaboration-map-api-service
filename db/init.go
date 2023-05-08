package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func GetClientOptions() *options.ClientOptions {
	dburi := "mongodb+srv://ghousedev:12W2r3w4r568@ghouse.4azrsyy.mongodb.net/?retryWrites=true&w=majority"

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(dburi).
		SetServerAPIOptions(serverAPIOptions)

	return clientOptions
}


func InitDb() {
	clientOptions := GetClientOptions()

	newClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	} else {
		client = newClient
	}
}

func GetMongoClient() mongo.Client {
	return *client
}
