package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/MuApp")
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
}
