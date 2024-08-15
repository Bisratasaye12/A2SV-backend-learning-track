package tests

import (
	controllers "Task-8/Delivery/Controllers"
	domain "Task-8/Domain"
	mocks "Task-8/Mocks"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)




type UserControllerSuite struct {
	suite.Suite

	usecase 		*mocks.UserUseCase
	controller 		*controllers.UserController
	testingServer   *httptest.Server
}



func (suite *UserControllerSuite) SetupTest(){
	suite.usecase = new(mocks.UserUseCase)
	suite.controller = controllers.NewUserController(suite.usecase)

	r := gin.Default()
	// Auth routes group
	authRoutes := r.Group("/")
	{
		authRoutes.POST("/register", suite.controller.Register)
		authRoutes.POST("/login", suite.controller.Login)
	}

	// User routes group
	userRoutes := r.Group("/users")
	{
		userRoutes.PUT("/promote/:id", suite.controller.PromoteUser)
	}

	suite.testingServer = httptest.NewServer(r)
}



func (suite *UserControllerSuite) TearDownTest(){
	suite.testingServer.Close()
}


func (suite *UserControllerSuite) TestRegister_Positive() {
	newUser := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "testuser",
		Password: "password123",
	}

	suite.usecase.On("Register", mock.Anything, &newUser).Return(&newUser, nil)

	userJSON, _ := json.Marshal(newUser)
	response, err := http.Post(fmt.Sprintf("%s/register", suite.testingServer.URL), "application/json", bytes.NewBuffer(userJSON))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := map[string]domain.User{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal(newUser, responseBody["user"])
	suite.usecase.AssertExpectations(suite.T())
}




func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(TaskControllerSuite))
}
