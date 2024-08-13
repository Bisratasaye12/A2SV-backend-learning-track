package usecases

import (
	"Task-8/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)





type userUseCase struct {
    UserRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) *userUseCase {
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