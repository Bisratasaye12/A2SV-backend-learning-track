package infrastructure

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
    MongoClient *mongo.Client
	Database *mongo.Database
)


func InitDB(uri string) {
    clientOptions := options.Client().ApplyURI(uri)

    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal("Failed to connect to MongoDB:", err)
    }

    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal("Failed to ping MongoDB:", err)
    }

    MongoClient = client

	// ping
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}
	fmt.Println("Connected to MongoDB!")
	Database = client.Database("task_management")
}
