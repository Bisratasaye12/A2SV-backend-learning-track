package usecases

import (
	"Task-8/Domain"
	repositories "Task-8/Repositories"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type UserUseCase interface {
    Register(ctx context.Context, user *domain.User) (domain.User, error)
    Login(ctx context.Context, user *domain.User) (string, error)
    PromoteUser(ctx context.Context, id primitive.ObjectID) error
}


type userUseCase struct {
    UserRepo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) *userUseCase {
    return &userUseCase{
        UserRepo: userRepo,
    }
}   



func (uc *userUseCase) Register(ctx context.Context, user *domain.User) (domain.User, error){
    return uc.UserRepo.Register(ctx, user)
}


func (uc *userUseCase) Login(ctx context.Context, user *domain.User) (string, error){
    return uc.UserRepo.Login(ctx, user)
}



func (uc *userUseCase) PromoteUser(ctx context.Context, id primitive.ObjectID) error{
    return uc.UserRepo.PromoteUser(ctx, id)
}