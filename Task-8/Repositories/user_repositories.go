package repositories

import (
	"Task-8/Domain"
	infrastructure "Task-8/Infrastructure"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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


func (ur mongoUserRepository) Register(ctx context.Context, user *domain.User) (domain.User, error){
	if user.Username == "" || user.Password == "" {
		return domain.User{},fmt.Errorf("missing required fields")
	}

	existingUser := ur.collection.FindOne(context.TODO(), bson.D{{Key: "username", Value: user.Username}})
	if existingUser.Err() != mongo.ErrNoDocuments {
		return domain.User{},fmt.Errorf("username already in use")
	}
	
	var err error
	user.Password, err = infrastructure.EncryptPassword(user.Password)
	if err != nil{
		return domain.User{}, err
	}
	user.CreatedAt = time.Now()

	count, err := ur.collection.CountDocuments(context.TODO(), bson.D{{}})
	if err != nil{
		return domain.User{},fmt.Errorf("internal server error")
	}

	if count == 0{
		user.Role = "admin"
	} else{
		user.Role = "user"
	}

	insertResult, err := ur.collection.InsertOne(ctx, user)
	if err != nil{
		return domain.User{}, fmt.Errorf("unable to register")
	}
	
	user.ID = insertResult.InsertedID.(primitive.ObjectID)
	return *user, nil
}


func (uc *mongoUserRepository) Login(ctx context.Context, user *domain.User) (string, error){
	var existingUser domain.User
	err := uc.collection.FindOne(context.TODO(), bson.D{{Key: "username", Value: user.Username}}).Decode(&existingUser)
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}
	return infrastructure.JWT_Auth(&existingUser, user)
}


func (uc *mongoUserRepository) PromoteUser(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.M{"$set": bson.M{"role": "admin"}}

	_, err := uc.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil{
		return fmt.Errorf("internal server error")
	}
	return nil
}