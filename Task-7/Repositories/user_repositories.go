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
    CreateUser(ctx context.Context, user *domain.User) (domain.User, error)
    GetUserByID(ctx context.Context, id primitive.ObjectID) (domain.User, error)
    GetUserByEmail(ctx context.Context, email string) (domain.User, error)
    UpdateUser(ctx context.Context, id primitive.ObjectID, updatedUser *domain.User) (domain.User, error)
    DeleteUser(ctx context.Context, id primitive.ObjectID) error
}
