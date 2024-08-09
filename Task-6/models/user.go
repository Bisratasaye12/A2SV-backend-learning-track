package models

import (
	"fmt"
	"time"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type User struct {
    ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Username  string             `json:"username" bson:"username" validate:"required,min=3,max=30"`
    Email     string             `json:"email" bson:"email" validate:"email"`
    Password  string             `json:"password" bson:"password" validate:"required,min=8"`
    Role      string             `json:"role" bson:"role"`
    CreatedAt time.Time          `json:"created_at" bson:"created_at"`
    UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}


func (u *User) Validate() error {
    validate := validator.New()
    en := en.New()
    uni := ut.New(en, en)
    trans, _ := uni.GetTranslator("en")
    en_translations.RegisterDefaultTranslations(validate, trans)

    err := validate.Struct(u)
    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            return fmt.Errorf(err.Translate(trans))
        }
    }
    return nil
}
