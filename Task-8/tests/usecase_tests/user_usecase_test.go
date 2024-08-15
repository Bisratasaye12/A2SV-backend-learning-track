package tests

import (
	domain "Task-8/Domain"
	mocks "Task-8/Mocks"
	usecases "Task-8/UseCases"
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



type userUseCaseSuite struct {
	suite.Suite
	repository 		*mocks.UserRepository
	usecase    		*usecases.Userusecase
	infrastructure  *mocks.Infrastructure
}


func (suite *userUseCaseSuite) SetupTest(){
	suite.repository = new(mocks.UserRepository)
	suite.infrastructure = &mocks.Infrastructure{}
	suite.usecase = &usecases.Userusecase{UserRepo: suite.repository, Infra: suite.infrastructure }
}



func (suite *userUseCaseSuite) TestRegister_Positive(){
	test_user := &domain.User{
		Username: "test username",
		Email: "test@email.com",
		Password: "test-password",
	}

	hashedPassword := "hashed-password"
	suite.infrastructure.On("EncryptPassword", test_user.Password).Return(hashedPassword, nil)

	no_user := true
	suite.repository.On("NoUsers", context.TODO()).Return(no_user, nil)
	suite.repository.On("Register", context.TODO(), test_user).Return(*test_user, nil)

	
	ret, err := suite.usecase.Register(context.TODO(), test_user)

	suite.NoError(err)
	suite.NotEmpty(ret)
	suite.NotZero(ret.CreatedAt)
	suite.Equal("admin", ret.Role)
	suite.Equal(hashedPassword, ret.Password)

	suite.infrastructure.AssertExpectations(suite.T())
	suite.repository.AssertExpectations(suite.T())
}

func (suite *userUseCaseSuite) TestRegister_EncryptionFailure() {
    test_user := &domain.User{
        Username: "test username",
        Email:    "test@email.com",
        Password: "test-password",
    }

    suite.infrastructure.On("EncryptPassword", test_user.Password).Return("", fmt.Errorf("encryption error"))

    result, err := suite.usecase.Register(context.TODO(), test_user)


    suite.Error(err)
    suite.Equal(domain.User{}, result) 

    suite.infrastructure.AssertExpectations(suite.T())
}

func (suite *userUseCaseSuite) TestRegister_NoUsersCheckFailure() {
    test_user := &domain.User{
        Username: "test username",
        Email:    "test@email.com",
        Password: "test-password",
    }

    hashedPassword := "hashed-password"
    suite.infrastructure.On("EncryptPassword", test_user.Password).Return(hashedPassword, nil)

    suite.repository.On("NoUsers", context.TODO()).Return(false, fmt.Errorf("repository error"))

    result, err := suite.usecase.Register(context.TODO(), test_user)

    suite.Error(err)
    suite.Equal(domain.User{}, result) 

    suite.infrastructure.AssertExpectations(suite.T())
    suite.repository.AssertExpectations(suite.T())
}

func (suite *userUseCaseSuite) TestLogin_Positive() {
    test_user := &domain.User{
        Username: "test-username",
        Email:    "test-email@example.com",
        Password: "test-password",
    }
    
    existingUser := &domain.User{
        Username: "test-username",
        Email:    "test-email@example.com",
        Password: "hashed-password",
    }

    jwtToken := "jwt-token"
    
    suite.repository.On("Login", context.TODO(), test_user.Username).Return(existingUser, nil)
    suite.infrastructure.On("JWT_Auth", existingUser, test_user).Return(jwtToken, nil)

    token, err := suite.usecase.Login(context.TODO(), test_user)
	log.Println("the token", token)

    suite.NoError(err)
    suite.Equal(jwtToken, token)

    suite.repository.AssertExpectations(suite.T())
    suite.infrastructure.AssertExpectations(suite.T())
}


func (suite *userUseCaseSuite) TestLogin_RepositoryLoginFailure() {
    test_user := &domain.User{
        Username: "test-username",
        Password: "test-password",
    }

    suite.repository.On("Login", context.TODO(), test_user.Username).Return(&domain.User{}, fmt.Errorf("repository error"))

    token, err := suite.usecase.Login(context.TODO(), test_user)

    suite.Error(err)
    suite.Empty(token)

    suite.repository.AssertExpectations(suite.T())
}

func (suite *userUseCaseSuite) TestLogin_JWTAuthFailure() {
    test_user := &domain.User{
        Username: "test-username",
        Password: "test-password",
    }

    existingUser := &domain.User{
        Username: "test-username",
        Password: "hashed-password",
    }

    suite.repository.On("Login", context.TODO(), test_user.Username).Return(existingUser, nil)
    suite.infrastructure.On("JWT_Auth", existingUser, test_user).Return("", fmt.Errorf("jwt auth error"))

    token, err := suite.usecase.Login(context.TODO(), test_user)

    suite.Error(err)
    suite.Empty(token)

    suite.repository.AssertExpectations(suite.T())
    suite.infrastructure.AssertExpectations(suite.T())
}


func (suite *userUseCaseSuite) TestPromoteUser_Positive() {
    userID := primitive.NewObjectID()

    suite.repository.On("PromoteUser", context.TODO(), userID).Return(nil)

    err := suite.usecase.PromoteUser(context.TODO(), userID)

    suite.NoError(err)

    suite.repository.AssertExpectations(suite.T())
}

func (suite *userUseCaseSuite) TestPromoteUser_Negative() {
    userID := primitive.NewObjectID()

    suite.repository.On("PromoteUser", context.TODO(), userID).Return(fmt.Errorf("repository error"))

    err := suite.usecase.PromoteUser(context.TODO(), userID)

    suite.Error(err)

    suite.repository.AssertExpectations(suite.T())
}


func TestUserUsecase(t *testing.T) {
	suite.Run(t, new(userUseCaseSuite))
}