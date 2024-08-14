package repositories

import (
	domain "Task-8/Domain"
	infrastructure "Task-8/Infrastructure"
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type UserRepositorySuite struct{
	suite.Suite

	repository *mongoUserRepository
	db_cleaner *DB_Cleanup
}


func (suite *UserRepositorySuite) SetupSuite(){
	err := godotenv.Load("../.env")
	suite.NoError(err) 
	
    
    uri := os.Getenv("MONGODB_URI")
	
	infrastructure.InitDB(uri)
	db := infrastructure.Database
	cleaner := InitCleanupDB(db, "test-users")

	suite.repository = &mongoUserRepository{db.Collection("test-users")}
	suite.db_cleaner = cleaner
}

func (suite *UserRepositorySuite) TearDownSuite(){
	suite.db_cleaner.CleanUp("test-users")
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
	suite.NoError(err, "error due to missing field password")

	suite.NotEmpty(regUser)
	suite.Empty(regUser2)
	suite.Empty(regUser3)
}


func (suite *UserRepositorySuite) Login(){
	
}

