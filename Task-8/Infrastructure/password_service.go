package infrastructure

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)


func EncryptPassword(password string) (string, error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "",fmt.Errorf("internal server error")
	}
	return string(hashedPassword), err
}