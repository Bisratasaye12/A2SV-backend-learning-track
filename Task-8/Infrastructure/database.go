package infrastructure

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Infrastruct struct{
	MongoClient 	*mongo.Client
	Database    	*mongo.Database
}


func NewInfrastructure() *Infrastruct {
	return &Infrastruct{
		MongoClient: nil,
		Database: nil,
	}
}

func (i *Infrastruct) InitDB(uri string) {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	i.MongoClient = client

	// ping
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}
	log.Println("Connected to MongoDB!")
	i.Database = client.Database("Task_management")
}
