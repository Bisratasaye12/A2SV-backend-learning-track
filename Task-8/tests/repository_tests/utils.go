package tests

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)


type DB_Cleanup struct {
	db *mongo.Database
}

func InitCleanupDB(db *mongo.Database, collection_name string) *DB_Cleanup{
	return &DB_Cleanup{
		db,
	}
}

func (cleaner *DB_Cleanup) CleanUp(collection_name string){
	cleaner.db.Collection(collection_name).Drop(context.TODO())
}



