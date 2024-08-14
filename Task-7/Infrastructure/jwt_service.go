package infrastructure

import (
	domain "Task-7/Domain"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)



func (i *Infrastruct) JWT_Auth(existingUser *domain.User, user *domain.User) (string, error){
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
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