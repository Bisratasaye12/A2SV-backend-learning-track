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




type TaskControllerSuite struct {
	suite.Suite

	usecase *mocks.TaskUseCase
	controller *controllers.TaskController
	testingServer   *httptest.Server
}



func (suite *TaskControllerSuite) SetupTest(){
	suite.usecase = new(mocks.TaskUseCase)
	suite.controller = controllers.NewTaskController(suite.usecase)

	r := gin.Default()
	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.GET("/", suite.controller.GetAllTasks)
		taskRoutes.GET("/:id", suite.controller.GetTaskByID)
		taskRoutes.POST("/",  suite.controller.AddTask)
		taskRoutes.PUT("/:id",  suite.controller.UpdateTask)
		taskRoutes.DELETE("/:id",  suite.controller.DeleteTask)
	}

	suite.testingServer = httptest.NewServer(r)
}



func (suite *TaskControllerSuite) TearDownTest(){
	suite.testingServer.Close()
}

func (suite *TaskControllerSuite) TestGetAllTasks_Positive() {
	tasks := []domain.Task{
		{
			ID:          primitive.NewObjectID(),
			Title:       "Task 1",
			Description: "Description 1",
			Status:      "Pending",
		},
		{
			ID:          primitive.NewObjectID(),
			Title:       "Task 2",
			Description: "Description 2",
			Status:      "Completed",
		},
	}

	suite.usecase.On("GetAllTasks", mock.Anything).Return(tasks, nil)

	response, err := http.Get(fmt.Sprintf("%s/tasks", suite.testingServer.URL))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusOK, response.StatusCode)

	var responseBody map[string][]domain.Task
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err, "error decoding response body")

	tasksFromResponse, ok := responseBody["tasks"]
	suite.True(ok, "response body should contain 'tasks' key")

	expectedJSON, err := json.Marshal(tasks)
	suite.NoError(err, "error marshalling expected tasks")

	actualJSON, err := json.Marshal(tasksFromResponse)
	suite.NoError(err, "error marshalling response body tasks")

	suite.JSONEq(string(expectedJSON), string(actualJSON), "response body matches expected tasks")
	suite.usecase.AssertExpectations(suite.T())
}


func (suite *TaskControllerSuite) TestGetAllTasks_Empty() {
	var tasks []domain.Task

	suite.usecase.On("GetAllTasks", mock.Anything).Return(tasks, nil)

	response, err := http.Get(fmt.Sprintf("%s/tasks", suite.testingServer.URL))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusOK, response.StatusCode)

	var responseBody map[string][]domain.Task
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err, "error decoding response body")

	tasksFromResponse, ok := responseBody["tasks"]
	suite.True(ok, "response body should contain 'tasks' key")

	suite.Empty(tasksFromResponse, "tasks list should be empty")
	suite.usecase.AssertExpectations(suite.T())
}


func (suite *TaskControllerSuite) TestGetTaskByID_Positive() {
	taskID := primitive.NewObjectID()
	task := domain.Task{
		ID:          taskID,
		Title:       "Task 1",
		Description: "Description 1",
		Status:      "Pending",
	}

	suite.usecase.On("GetTaskByID", mock.Anything, taskID).Return(task, nil)

	response, err := http.Get(fmt.Sprintf("%s/tasks/%s", suite.testingServer.URL, taskID.Hex()))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusOK, response.StatusCode)

	var responseBody map[string]domain.Task
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err, "error decoding response body")

	taskFromResponse, ok := responseBody["task"]
	suite.True(ok, "response body should contain 'task' key")

	suite.Equal(task, taskFromResponse)
	suite.usecase.AssertExpectations(suite.T())
}


func (suite *TaskControllerSuite) TestGetTaskByID_Negative() {
	invalidID := "invalid_id"

	response, err := http.Get(fmt.Sprintf("%s/tasks/%s", suite.testingServer.URL, invalidID))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusBadRequest, response.StatusCode)


	var responseBody map[string]string
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err, "error decoding response body")

	message, ok := responseBody["message"]
	suite.True(ok, "response body should contain 'message' key")
	suite.Equal("invalid id", message)

	taskID := primitive.NewObjectID()
	suite.usecase.On("GetTaskByID", mock.Anything, taskID).Return(domain.Task{}, fmt.Errorf("task not found"))

	response, err = http.Get(fmt.Sprintf("%s/tasks/%s", suite.testingServer.URL, taskID.Hex()))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusInternalServerError, response.StatusCode)

	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err, "error decoding response body")

	message, ok = responseBody["message"]
	suite.True(ok, "response body should contain 'message' key")
	suite.Equal("unable to fetch task", message)
	suite.usecase.AssertExpectations(suite.T())
}


func (suite *TaskControllerSuite) TestAddTask_Positive() {
	newTask := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "New Task",
		Description: "New Task Description",
		Status:      "Pending",
	}

	suite.usecase.On("AddTask", mock.Anything, &newTask).Return(newTask, nil)

	
	jsonPayload, _ := json.Marshal(&newTask)

	response, err := http.Post(fmt.Sprintf("%s/tasks", suite.testingServer.URL), "application/json", bytes.NewBuffer(jsonPayload))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusCreated, response.StatusCode)

	var responseBody map[string]domain.Task
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err, "error decoding response body")

	addedTaskFromResponse, ok := responseBody["added_task"]
	suite.True(ok, "response body should contain 'added_task' key")

	
	suite.Equal(newTask, addedTaskFromResponse)
	suite.usecase.AssertExpectations(suite.T())
}


func (suite *TaskControllerSuite) TestAddTask_Negative_Req() {

	invalidTask := `{
    "title": "New Task,
    "description": "This is a new task"
	}`

	invalidPayload, err := json.Marshal(&invalidTask)
	response, err := http.Post(fmt.Sprintf("%s/tasks", suite.testingServer.URL), "application/json", bytes.NewBuffer([]byte(invalidPayload)))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusBadRequest, response.StatusCode)

	var responseBody map[string]string
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err, "error decoding response body")

	message, ok := responseBody["message"]
	suite.True(ok, "response body should contain 'message' key")
	suite.Equal("invalid request type", message)
}

func (suite *TaskControllerSuite) TestAddTask_Negative_Res(){
	newTask := domain.Task{
		Title:       "New Task",
		Description: "New Task Description",
		Status:      "Pending",
	}

	suite.usecase.On("AddTask", mock.Anything, &newTask).Return(domain.Task{}, fmt.Errorf("task creation failed"))
	jsonPayload, _ := json.Marshal(&newTask)

	response, err := http.Post(fmt.Sprintf("%s/tasks", suite.testingServer.URL), "application/json", bytes.NewBuffer(jsonPayload))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusInternalServerError, response.StatusCode)

	var responseBody map[string]string
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err, "error decoding response body")

	message, ok := responseBody["message"]
	suite.True(ok, "response body should contain 'message' key")
	suite.Equal("unable to add task", message)
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *TaskControllerSuite) TestUpdateTask_Positive() {
	taskID := primitive.NewObjectID()
	updatedTask := domain.Task{
		ID:          taskID,
		Title:       "Updated Task",
		Description: "Updated Task Description",
		Status:      "Completed",
	}

	suite.usecase.On("UpdateTask", mock.Anything, taskID, &updatedTask).Return(updatedTask, nil)


	jsonPayload, _ := json.Marshal(&updatedTask)

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/tasks/%s", suite.testingServer.URL, taskID.Hex()), bytes.NewBuffer(jsonPayload))
	suite.NoError(err, "no error when creating the request")
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.testingServer.Client().Do(req)
	suite.NoError(err, "no error when performing the request")
	defer resp.Body.Close()

	suite.Equal(http.StatusOK, resp.StatusCode)

	var responseBody map[string]domain.Task
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	suite.NoError(err, "error decoding response body")

	updatedTaskFromResponse, ok := responseBody["updated_task"]
	suite.True(ok, "response body should contain 'updated_task' key")

	suite.Equal(updatedTask, updatedTaskFromResponse)
	suite.usecase.AssertExpectations(suite.T())
}


func (suite *TaskControllerSuite) TestUpdateTask_Negative() {
	// Test with invalid ID
	invalidID := "invalid_id"
	updatedTask := domain.Task{
		Title:       "Updated Task",
		Description: "Updated Task Description",
		Status:      "Completed",
	}


	jsonPayload, _ := json.Marshal(&updatedTask)

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/tasks/%s", suite.testingServer.URL, invalidID), bytes.NewBuffer(jsonPayload))
	suite.NoError(err, "no error when creating the request")
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.testingServer.Client().Do(req)
	suite.NoError(err, "no error when performing the request")
	defer resp.Body.Close()

	suite.Equal(http.StatusBadRequest, resp.StatusCode)

	var responseBody map[string]string
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	suite.NoError(err, "error decoding response body")

	message, ok := responseBody["message"]
	suite.True(ok, "response body should contain 'message' key")
	suite.Equal("invalid id", message)

	// Test with invalid JSON payload
	invalidPayload := `{"title": "Updated Task, "description": "Updated Task Description"}`
	req, err = http.NewRequest(http.MethodPut, fmt.Sprintf("%s/tasks/%s", suite.testingServer.URL, primitive.NewObjectID().Hex()), bytes.NewBuffer([]byte(invalidPayload)))
	suite.NoError(err, "no error when creating the request")
	req.Header.Set("Content-Type", "application/json")

	resp, err = suite.testingServer.Client().Do(req)
	suite.NoError(err, "no error when performing the request")
	defer resp.Body.Close()

	suite.Equal(http.StatusBadRequest, resp.StatusCode)

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	suite.NoError(err, "error decoding response body")

	message, ok = responseBody["message"]
	suite.True(ok, "response body should contain 'message' key")
	suite.Equal("invalid task", message)

	// Test with use case error
	taskID := primitive.NewObjectID()
	suite.usecase.On("UpdateTask", mock.Anything, taskID, &updatedTask).Return(domain.Task{}, fmt.Errorf("task update failed"))

	// Perform the HTTP PUT request with valid task data
	req, err = http.NewRequest(http.MethodPut, fmt.Sprintf("%s/tasks/%s", suite.testingServer.URL, taskID.Hex()), bytes.NewBuffer(jsonPayload))
	suite.NoError(err, "no error when creating the request")
	req.Header.Set("Content-Type", "application/json")

	resp, err = suite.testingServer.Client().Do(req)
	suite.NoError(err, "no error when performing the request")
	defer resp.Body.Close()

	suite.Equal(http.StatusInternalServerError, resp.StatusCode)

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	suite.NoError(err, "error decoding response body")

	
	message, ok = responseBody["message"]
	suite.True(ok, "response body should contain 'message' key")
	suite.Equal("unable to update task", message)
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *TaskControllerSuite) TestDeleteTask_Positive() {
	taskID := primitive.NewObjectID()

	suite.usecase.On("DeleteTask", mock.Anything, taskID).Return(nil)

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/tasks/%s", suite.testingServer.URL, taskID.Hex()), nil)
	suite.NoError(err, "no error when creating the request")

	resp, err := suite.testingServer.Client().Do(req)
	suite.NoError(err, "no error when performing the request")
	defer resp.Body.Close()

	suite.Equal(http.StatusNoContent, resp.StatusCode)

	suite.usecase.AssertExpectations(suite.T())
}

func (suite *TaskControllerSuite) TestDeleteTask_Negative() {
	// Test with invalid ID
	invalidID := "invalid_id"

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/tasks/%s", suite.testingServer.URL, invalidID), nil)
	suite.NoError(err, "no error when creating the request")

	resp, err := suite.testingServer.Client().Do(req)
	suite.NoError(err, "no error when performing the request")
	defer resp.Body.Close()

	suite.Equal(http.StatusBadRequest, resp.StatusCode)

	var responseBody map[string]string
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	suite.NoError(err, "error decoding response body")

	message, ok := responseBody["message"]
	suite.True(ok, "response body should contain 'message' key")
	suite.Equal("invalid id", message)

	// Test with use case error
	taskID := primitive.NewObjectID()
	suite.usecase.On("DeleteTask", mock.Anything, taskID).Return(fmt.Errorf("task deletion failed"))

	req, err = http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/tasks/%s", suite.testingServer.URL, taskID.Hex()), nil)
	suite.NoError(err, "no error when creating the request")

	resp, err = suite.testingServer.Client().Do(req)
	suite.NoError(err, "no error when performing the request")
	defer resp.Body.Close()

	suite.Equal(http.StatusInternalServerError, resp.StatusCode)

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	suite.NoError(err, "error decoding response body")

	message, ok = responseBody["message"]
	suite.True(ok, "response body should contain 'message' key")
	suite.Equal("unable to delete task", message)

	suite.usecase.AssertExpectations(suite.T())
}


func TestTaskControllerSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerSuite))
}
