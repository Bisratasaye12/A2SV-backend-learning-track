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

func (suite *UserControllerSuite) TestRegister_Negative() {
	newUser := domain.User{
		Username: "testuser",
		Password: "password123",
	}

	suite.usecase.On("Register", mock.Anything, &newUser).Return(nil, fmt.Errorf("registration error"))

	userJSON, _ := json.Marshal(newUser)
	response, err := http.Post(fmt.Sprintf("%s/register", suite.testingServer.URL), "application/json", bytes.NewBuffer(userJSON))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	var responseBody map[string]interface{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	suite.Equal(http.StatusInternalServerError, response.StatusCode)
	suite.Equal("registration error", responseBody["error"])
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *UserControllerSuite) TestLogin_Positive() {
	user := domain.User{
		Username: "testuser",
		Password: "password123",
	}

	suite.usecase.On("Login", mock.Anything, user.Username, user.Password).Return("token123", nil)

	loginData := map[string]string{
		"username": user.Username,
		"password": user.Password,
	}
	loginJSON, _ := json.Marshal(loginData)
	response, err := http.Post(fmt.Sprintf("%s/login", suite.testingServer.URL), "application/json", bytes.NewBuffer(loginJSON))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	var responseBody map[string]string
	json.NewDecoder(response.Body).Decode(&responseBody)
	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal("token123", responseBody["token"])
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *UserControllerSuite) TestPromoteUser_Positive() {
	userID := primitive.NewObjectID()

	suite.usecase.On("PromoteUser", mock.Anything, userID).Return(nil)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/users/promote/%s", suite.testingServer.URL, userID.Hex()), nil)
	suite.NoError(err, "no error when creating the request")

	resp := httptest.NewRecorder()
	suite.controller.PromoteUser(&gin.Context{Request: req})

	suite.Equal(http.StatusOK, resp.Code)

	var responseBody map[string]string
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	suite.NoError(err, "no error when decoding response body")
	suite.Equal("User promoted to admin", responseBody["message"])

	suite.usecase.AssertExpectations(suite.T())
}


func (suite *UserControllerSuite) TestPromoteUser_Negative() {
	userID := primitive.NewObjectID()

	suite.usecase.On("PromoteUser", mock.Anything, userID).Return(fmt.Errorf("promotion error"))

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/users/promote/%s", suite.testingServer.URL, userID.Hex()), nil)
	suite.NoError(err, "no error when creating the request")

	resp := httptest.NewRecorder()
	suite.controller.PromoteUser(&gin.Context{Request: req})
	suite.Equal(http.StatusInternalServerError, resp.Code)

	var responseBody map[string]string
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	suite.NoError(err, "no error when decoding response body")
	suite.Equal("promotion error", responseBody["error"])


	suite.usecase.AssertExpectations(suite.T())
}


func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(TaskControllerSuite))
}
