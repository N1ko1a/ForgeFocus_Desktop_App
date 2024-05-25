package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client // Global variable to hold the MongoDB client

func ConnectToDB() {
	// Find .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Get value from .env
	MONGO_URI := os.Getenv("MONGO_URI")

	// Connect to the database.
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	mongoClient, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection.
	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to db")
}

func connectToCollection(collectionName string) (*mongo.Collection, context.Context, error) {
	if mongoClient == nil {
		return nil, nil, fmt.Errorf("mongoClient is nil, not connected to MongoDB server")
	}

	ctx := context.Background()
	Database := mongoClient.Database("Productivity")
	Collection := Database.Collection(collectionName)

	err := mongoClient.Ping(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	return Collection, ctx, nil
}
