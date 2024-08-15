package repositories

import (
	domain "Task-8/Domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	Collection *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{
		Collection: db.Collection("users"),
	}
}

func (ur *MongoUserRepository) NoUsers(ctx context.Context) (bool, error) {
	count, err := ur.Collection.CountDocuments(context.TODO(), bson.D{{}})
	if err != nil {
		return false, fmt.Errorf("internal server error")
	}
	if count != 0 {
		return false, nil
	}
	return true, nil
}

func (ur MongoUserRepository) Register(ctx context.Context, user *domain.User) (domain.User, error) {
	if user.Username == "" || user.Password == "" {
		return domain.User{}, fmt.Errorf("missing required fields")
	}

	existingUser := ur.Collection.FindOne(context.TODO(), bson.D{{Key: "username", Value: user.Username}})
	if existingUser.Err() != mongo.ErrNoDocuments {
		return domain.User{}, fmt.Errorf("username already in use")
	}

	insertResult, err := ur.Collection.InsertOne(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("unable to register")
	}

	user.ID = insertResult.InsertedID.(primitive.ObjectID)
	return *user, nil
}

func (uc *MongoUserRepository) Login(ctx context.Context, username string) (*domain.User, error) {
	var existingUser domain.User
	err := uc.Collection.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}}).Decode(&existingUser)
	if err != nil {
		return &domain.User{}, fmt.Errorf("invalid username or password")
	}
	return &existingUser, nil
}

func (uc *MongoUserRepository) PromoteUser(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.M{"$set": bson.M{"role": "admin"}}

	result, err := uc.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("internal server error")
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("no user found")

	}
	return nil
}
