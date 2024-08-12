package usecases

import (
	"Task-7/Domain"
	repositories "Task-7/Repositories"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type UserUseCase interface {
    Register(ctx context.Context, user *domain.User) error
    Login(ctx context.Context, user *domain.User) (string, error)
    PromoteUser(ctx context.Context, id primitive.ObjectID) error
    GetUserByID(ctx context.Context, id primitive.ObjectID) (domain.User, error)
    GetUserByEmail(ctx context.Context, email string) (domain.User, error)
}


type userUseCase struct {
    UserRepo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) *userUseCase {
    return &userUseCase{
        UserRepo: userRepo,
    }
}   