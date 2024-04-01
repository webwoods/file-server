// mongo_singleton.go

package database

import (
	"context"
	"os"
	"sync"

	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	mongoOnce   sync.Once
)

// GetMongoClient returns a singleton instance of the MongoDB client.
func GetMongoClient() (*mongo.Client, error) {
	var err error

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// Initialize the MongoDB client only once
	mongoOnce.Do(func() {
		// Set the server API options
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		uri := os.Getenv("ds")

		opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

		// Create a new client and connect to the server
		mongoClient, err = mongo.Connect(context.TODO(), opts)
		if err != nil {
			println("error connecting to server: " + err.Error())
			panic(err)
		}

		// Print a message indicating the connection is established
		println("MongoDB client connected successfully.")

		// // Defer the disconnection of the client
		// // Note: This will disconnect the client when the program exits
		// defer func() {
		// 	if err = mongoClient.Disconnect(context.TODO()); err != nil {
		// 		println("error occured. disconneting: " + err.Error())
		// 		panic(err)
		// 	}
		// }()

		// // Ping the deployment to confirm a successful connection
		// if err := mongoClient.Database("reware").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		// 	panic(err)
		// }
	})

	return mongoClient, err
}

// DisconnectMongoClient disconnects the MongoDB client.
func DisconnectMongoClient() {
	if mongoClient != nil {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			println("Error occurred while disconnecting MongoDB client:", err.Error())
		} else {
			println("MongoDB client disconnected successfully.")
		}
	}
}
