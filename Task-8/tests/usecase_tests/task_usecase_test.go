package tests

import (
	domain "Task-8/Domain"
	mocks "Task-8/Mocks"
	usecases "Task-8/UseCases"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskUseCaseSuite struct {
	suite.Suite
	repository    *mocks.TaskRepository
	usecase       *usecases.Taskusecase
}

func (suite *taskUseCaseSuite) SetupTest() {
	suite.repository = new(mocks.TaskRepository)
	suite.usecase = &usecases.Taskusecase{TaskRepo: suite.repository}
}

func (suite *taskUseCaseSuite) TestGetAllTasks_Empty() {
	tasks := []domain.Task{}
	suite.repository.On("GetAllTasks", context.TODO()).Return(tasks, nil)

	ret_tasks, err := suite.usecase.GetAllTasks(context.TODO())
	suite.NoError(err, "no error on valid request")
	suite.Empty(ret_tasks)

	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUseCaseSuite) TestGetAllTasks_Positive() {
	tasks := []domain.Task{
		{
			Title:       "test 1",
			Description: "title for test 1",
		},
	}
	suite.repository.On("GetAllTasks", context.TODO()).Return(tasks, nil)

	ret_tasks, err := suite.usecase.GetAllTasks(context.TODO())
	suite.NoError(err, "no error on valid request")
	suite.NotEmpty(ret_tasks)
	suite.Len(ret_tasks, len(tasks))

	for i := range tasks {
		suite.Equal(tasks[i], ret_tasks[i])
	}

	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUseCaseSuite) TestAddTask_Positive() {
	test_task := &domain.Task{
		Title:       "test title",
		Description: "test description",
		Status:      "in progress",
	}

	suite.repository.On("AddTask", context.TODO(), test_task).Return(*test_task, nil)

	ret_task, err := suite.usecase.AddTask(context.TODO(), test_task)

	suite.NoError(err)
	suite.NotNil(ret_task, "returned task should not be nil")
	suite.Equal(test_task.Title, ret_task.Title, "titles should match")
	suite.Equal(test_task.Description, ret_task.Description, "descriptions should match")
	suite.Equal(test_task.Status, ret_task.Status, "statuses should match")

	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUseCaseSuite) TestAddTask_Negative() {
	test_task := &domain.Task{
		Title:       "",
		Description: "",
		Status:      "in progress",
	}

	ret_task, err := suite.usecase.AddTask(context.TODO(), test_task)

	suite.Error(err, "expected error when adding a task")
	suite.Empty(ret_task, "returned task should be empty")

	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUseCaseSuite) TestGetTaskByID_Positive() {
	testID := primitive.NewObjectID()
	expectedTask := &domain.Task{
		ID:          testID,
		Title:       "test title",
		Description: "test description",
		Status:      "in progress",
	}

	suite.repository.On("GetTaskByID", context.TODO(), testID).Return(*expectedTask, nil)

	retTask, err := suite.usecase.GetTaskByID(context.TODO(), testID)

	suite.NoError(err, "no error on valid request")
	suite.NotNil(retTask, "returned task should not be nil")
	suite.Equal(expectedTask.ID, retTask.ID, "task ID should match")
	suite.Equal(expectedTask.Title, retTask.Title, "task title should match")
	suite.Equal(expectedTask.Description, retTask.Description, "task description should match")
	suite.Equal(expectedTask.Status, retTask.Status, "task status should match")

	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUseCaseSuite) TestGetTaskByID_Negative() {
	testID := primitive.NewObjectID()
	expectedError := errors.New("task not found")

	suite.repository.On("GetTaskByID", context.TODO(), testID).Return(domain.Task{}, expectedError)
	retTask, err := suite.usecase.GetTaskByID(context.TODO(), testID)

	suite.Error(err, "expected an error for non-existent task")
	suite.Empty(retTask, "returned task should be empty")

	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUseCaseSuite) TestUpdateTask_Positive() {
	testID := primitive.NewObjectID()
	updated_Task := &domain.Task{
		ID:          testID,
		Title:       "new title",
		Description: "new description",
		Status:      "completed",
	}

	update_Task := &domain.Task{
		Title:       "new title",
		Description: "new description",
		Status:      "completed",
	}

	suite.repository.On("UpdateTask", context.TODO(), testID, update_Task).Return(*updated_Task, nil)

	result, err := suite.usecase.UpdateTask(context.TODO(), testID, update_Task)

	suite.NoError(err, "no error on valid update")
	suite.Equal(update_Task.Title, result.Title, "task title should match")
	suite.Equal(update_Task.Description, result.Description, "task description should match")
	suite.Equal(update_Task.Status, result.Status, "task status should match")

	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUseCaseSuite) TestDeleteTask_Positive() {
	testID := primitive.NewObjectID()
	suite.repository.On("DeleteTask", context.TODO(), testID).Return(nil)

	err := suite.usecase.DeleteTask(context.TODO(), testID)

	suite.NoError(err, "no error on valid delete")

	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUseCaseSuite) TestDeleteTask_Negative() {
	testID := primitive.NewObjectID()
	expectedError := errors.New("task not found")
	suite.repository.On("DeleteTask", context.TODO(), testID).Return(expectedError)

	err := suite.usecase.DeleteTask(context.TODO(), testID)

	suite.Error(err, "expected an error for non-existent task")

	suite.repository.AssertExpectations(suite.T())
}


func TestTaskUsecase(t *testing.T) {
	suite.Run(t, new(taskUseCaseSuite))
}