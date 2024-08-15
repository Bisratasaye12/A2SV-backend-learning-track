package tests

import (
	controllers "Task-8/Delivery/Controllers"
	mocks "Task-8/Mocks"
	"net/http/httptest"

	"github.com/stretchr/testify/suite"
)




type TaskControllerSuite struct {
	suite.Suite

	usecase *mocks.TaskUseCase
	controller *controllers.TaskController
	testingServer   *httptest.Server
}



func (suite *TaskControllerSuite) SetupTest(){
	suite.usecase = new(mocks.TaskUseCase)
	suite.controller = controllers.NewTaskController(suite.usecase)
}

