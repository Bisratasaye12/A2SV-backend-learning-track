package usecases

import (
	"Task-7/Domain"
	infrastructure "Task-7/Infrastructure"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)





type userUseCase struct {
    UserRepo domain.UserRepository
    infra    *infrastructure.Infrastruct
}

func NewUserUseCase(userRepo domain.UserRepository, infra *infrastructure.Infrastruct) *userUseCase {
    return &userUseCase{
        UserRepo: userRepo,
        infra:   infra,
    }
}   



func (uc *userUseCase) Register(ctx context.Context, user *domain.User) (domain.User, error){
    var err error
    log.Println("FROM REGISTER USER_USECASES: ",user.Password)
	user.Password, err = uc.infra.EncryptPassword(user.Password)
    
	if err != nil{
		return domain.User{}, err
	}
	user.CreatedAt = time.Now()

	noUser, err := uc.UserRepo.NoUsers(ctx)
    if err != nil{
        return domain.User{}, err
    }

    if noUser{
        user.Role = "admin"
    } else{
        user.Role = "user"
    }

    return uc.UserRepo.Register(ctx, user)
}


func (uc *userUseCase) Login(ctx context.Context, user *domain.User) (string, error){
    exsUser, err := uc.UserRepo.Login(ctx, user.Username)
    if err != nil{
        return "", err
    }
    return uc.infra.JWT_Auth(exsUser, user)
}



func (uc *userUseCase) PromoteUser(ctx context.Context, id primitive.ObjectID) error{
    return uc.UserRepo.PromoteUser(ctx, id)
}