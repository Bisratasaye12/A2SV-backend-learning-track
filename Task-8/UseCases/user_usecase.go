package usecases

import (
	"Task-8/Domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)





type Userusecase struct {
    UserRepo domain.UserRepository
    Infra    domain.Infrastructure
}

func NewUserUseCase(userRepo domain.UserRepository, infra domain.Infrastructure) *Userusecase {
    return &Userusecase{
        UserRepo: userRepo,
        Infra:   infra,
        
    }
}   



func (uc *Userusecase) Register(ctx context.Context, user *domain.User) (domain.User, error){
    var err error
	user.Password, err = uc.Infra.EncryptPassword(user.Password)
    
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

    ret_user, err := uc.UserRepo.Register(ctx, user)

    if err != nil{
        return domain.User{}, err
    }
    ret_user.CreatedAt = user.CreatedAt
    ret_user.Role = user.Role
    ret_user.Password = user.Password
    return ret_user, nil
}


func (uc *Userusecase) Login(ctx context.Context, user *domain.User) (string, error){
    exsUser, err := uc.UserRepo.Login(ctx, user.Username)
    if err != nil{
        return "", err
    }
    ret_token, err := uc.Infra.JWT_Auth(exsUser, user)

    return ret_token, err

}



func (uc *Userusecase) PromoteUser(ctx context.Context, id primitive.ObjectID) error{
    return uc.UserRepo.PromoteUser(ctx, id)
}