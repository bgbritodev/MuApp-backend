package config

import (
	"context"
	"fmt"
	"log"

	//"os"

	//"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	// Load .env file from the parent directory
	mongoTestClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://localhost:27017"))

	// Check the connection.
	err = mongoTestClient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to mongoDB!!!")
	}

	// Assign the client to the global variable
	log.Println("MongoDB client assigned to global variable")
	return mongoTestClient

}
