package data

import (
	"Task-6/models"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var (
    users    *mongo.Collection
)

func init() {
    users = dbClient.Database("Task_management").Collection("users")
}


func Register(user *models.User) error {
	
	if user.Username == "" || user.Password == "" {
		return fmt.Errorf("missing required fields")
	}

	existingUser := users.FindOne(context.TODO(), bson.D{{"username", user.Username}})
	if existingUser.Err() != mongo.ErrNoDocuments {
		return fmt.Errorf("Username already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("internal server error")
	}

	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()

	count, err := users.CountDocuments(context.TODO(), bson.D{{}})
	if err != nil{
		return fmt.Errorf("internal server error")
	}

	if count == 0{
		user.Role = "admin"
	} else{
		user.Role = "user"
	}

	_, err = users.InsertOne(context.TODO(), user)
	if err != nil{
		return fmt.Errorf("unable to register user")
	}
	return nil	
}

func Login(user *models.User) (string, error) {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

	var existingUser models.User
	err := users.FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(&existingUser)
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": existingUser.ID,
		"username":   existingUser.Username,
		"role":    existingUser.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("internal server error")
	}

	return jwtToken, nil
}

func PromoteUser(id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	update := bson.M{"$set": bson.M{"role": "admin"}}

	_, err := users.UpdateOne(context.TODO(), filter, update)
	if err != nil{
		return fmt.Errorf("internal server error")
	}
	return nil
}


func GetUsers() []models.User{
	var usrs []models.User
	var val models.User
	cursor, err := users.Find(context.TODO(), bson.D{{}})
	if err != nil{
		return nil
	}
	for cursor.Next(context.TODO()){
		cursor.Decode(&val)
		usrs = append(usrs, val)
	}
	defer cursor.Close(context.TODO())

	return usrs
}


