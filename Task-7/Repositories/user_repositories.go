package repositories

import (
	"Task-7/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database) *mongoUserRepository {
	return &mongoUserRepository{
		collection: db.Collection("users"),
	}
}

type UserRepository interface {
    Register(ctx context.Context, user *domain.User) (domain.User, error)
    GetUserByID(ctx context.Context, id primitive.ObjectID) (domain.User, error)
    Login(ctx context.Context) error
}


