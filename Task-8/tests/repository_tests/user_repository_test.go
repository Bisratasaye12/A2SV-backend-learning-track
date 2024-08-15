package tests

import (
	domain "Task-8/Domain"
	infrastructure_pack "Task-8/Infrastructure"
	repositories "Task-8/Repositories"
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepositorySuite struct{
	suite.Suite

	repository 		*repositories.MongoUserRepository
	db_cleaner 		*DB_Cleanup
}


func (suite *UserRepositorySuite) SetupTest(){
	err := godotenv.Load("../../.env")
	suite.NoError(err) 
	
    
    uri := os.Getenv("MONGODB_URI")
	db := infrastructure_pack.InitDB(uri)
	
	cleaner := InitCleanupDB(db, "test-users")

	suite.repository = &repositories.MongoUserRepository{Collection: db.Collection("test-users")}
	suite.db_cleaner = cleaner
}


func (suite *UserRepositorySuite) TearDownTest(){
	suite.db_cleaner.CleanUp("test-users")
}

func (suite *UserRepositorySuite) TestNoUsers_Positive(){
	noUsers, err := suite.repository.NoUsers(context.TODO())
	suite.NoError(err, "no error on valid input")
	suite.True(noUsers, "no users in the database")
}

func (suite *UserRepositorySuite) TestRegister_Positive(){
	test_user := &domain.User{
		Username: "test-user",
		Email: "test-user@gmail.com",
		Password: "12345678",
	}

	regUser, err := suite.repository.Register(context.TODO(), test_user)
	suite.NoError(err, "no error on valid input")
	suite.NotEmpty(regUser.ID, "user ID should not be empty")

	suite.Equal(test_user.Username, regUser.Username)
}



func (suite *UserRepositorySuite) TestRegister_Negative(){
	test_user := &domain.User{
		Username: "test-user",
		Email: "test-user@gmail.com",
		Password: "12345678",
	}

	test_user2 := &domain.User{
		Username: "test-user",
		Email: "test@gmail.com",
		Password: "12345678",
	}

	test_user3 := &domain.User{
		Username: "test-user3",
		Email: "test@gmail.com",
		Password: "",
	}

	regUser, err := suite.repository.Register(context.TODO(), test_user)
	suite.NoError(err, "no error on valid input")

	regUser2, err := suite.repository.Register(context.TODO(), test_user2)
	suite.Error(err, "error due to existing username")

	regUser3, err := suite.repository.Register(context.TODO(), test_user3)
	suite.Error(err, "error due to missing field password")

	suite.NotEmpty(regUser)
	suite.Empty(regUser2)
	suite.Empty(regUser3)
}


func (suite *UserRepositorySuite) TestLogin_Positive(){
	test_user := &domain.User{
		Username: "test-user",
		Email: "test-user@gmail.com",
		Password: "12345678",
	}

	regUser, err := suite.repository.Register(context.TODO(), test_user)
	suite.NoError(err, "no error on valid input")

	user, err := suite.repository.Login(context.TODO(), regUser.Username)
	suite.NoError(err, "no error in logging a valid user")
	suite.NotEmpty(user)
	suite.Equal(regUser.Username, user.Username)

}

func (suite *UserRepositorySuite) TestLogin_Negative() {
	
	nonExistentUsername := "non-existent-user"
	user, err := suite.repository.Login(context.TODO(), nonExistentUsername)
	suite.Error(err, "error expected when logging in with a non-existent username")
	suite.Empty(user, "returned user should be empty")

	test_user := &domain.User{
		Username: "test-user",
		Email:    "test-user@gmail.com",
		Password: "12345678",
	}

	_, err = suite.repository.Register(context.TODO(), test_user)
	suite.NoError(err, "no error on registering a valid user")

	incorrectPasswordUser := &domain.User{
		Username: "test-user",
		Password: "wrong-password",
	}

	incorrectUserNameUser := &domain.User{
		Username: "wrong-username",
		Password: "12345678",
	}

	user, err = suite.repository.Login(context.TODO(), incorrectPasswordUser.Username)
	suite.NoError(err, "error expected when logging in with incorrect password")
	suite.NotEqual(incorrectPasswordUser.Password, user.Password)
	
	user, err = suite.repository.Login(context.TODO(), incorrectUserNameUser.Username)
	suite.Error(err, "error expected when logging in with incorrect username")
	suite.Empty(user, "returned user should be empty")
}


func (suite *UserRepositorySuite) TestPromoteUser_Positive(){
	test_user := &domain.User{
		Username: "test-user",
		Email: "test-user@gmail.com",
		Password: "12345678",
	}

	regUser, err := suite.repository.Register(context.TODO(), test_user)
	suite.NoError(err, "no error on valid input")

	err = suite.repository.PromoteUser(context.TODO(), regUser.ID)
	suite.NoError(err, "no error on promoting a user to admin")
}

func (suite *UserRepositorySuite) TestPromoteUser_Negative(){
	test_user := &domain.User{
		Username: "test-user",
		Email: "test-user@gmail.com",
		Password: "12345678",
	}

	regUser, err := suite.repository.Register(context.TODO(), test_user)
	suite.NoError(err, "no error on valid input")
	
	invalidID := primitive.ObjectID{}
	err = suite.repository.PromoteUser(context.TODO(), invalidID)
	suite.Error(err, "error expected when promoting a user with an invalid ID")

	err = suite.repository.PromoteUser(context.TODO(), regUser.ID)
	suite.NoError(err, "no error on promoting a user to admin")
	suite.NotEqual("user", regUser.Role)
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}

