package repositories

import (
	domain "Task-8/Domain"
	infrastructure_pack "Task-8/Infrastructure"
	mocks "Task-8/Mocks"
	"context"
	"log"
	"os"

	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type TaskRepositorySuite struct{
	suite.Suite
	repository 		*mongoTaskRepository
	db_cleaner 		*DB_Cleanup
	infrastructure  *mocks.Infrastructure
}

func (suite *TaskRepositorySuite) SetupSuite(){
    err := godotenv.Load("../.env")
	suite.NoError(err) 
    
    uri := os.Getenv("MONGODB_URI")
	
	db := infrastructure_pack.InitDB(uri)
	log.Println("DB: ", db)
	testTaskRepo := &mongoTaskRepository{collection: db.Collection("test-tasks")}
	cleaner := InitCleanupDB(db, "test-tasks")

	suite.repository = testTaskRepo
	suite.db_cleaner = cleaner
}


// func (suite *TaskRepositorySuite) TearDownTest(){
// 	defer suite.db_cleaner.CleanUp("test-tasks")
// }


func (suite *TaskRepositorySuite) TearDownSuite(){
	defer suite.db_cleaner.CleanUp("test-tasks")
}

func (suite *TaskRepositorySuite) TestAddTask_Positive() {
	newTask := &domain.Task{
		Title:       "New Task",
		Description: "New Description",
		DueDate:     time.Now(),
		Status:      "pending",
	}
	task, err := suite.repository.AddTask(context.TODO(), newTask)

	suite.NoError(err, "no error on valid input")
	suite.NotEmpty(task.ID, "new id should be given by the database")
	suite.Equal(newTask.Title, task.Title)
}

func (suite *TaskRepositorySuite) TestAddTask_Negative() {
	invalidTask := &domain.Task{
		ID: primitive.NewObjectID(), 
		Title: "Invalid Task",
	}

	task, err := suite.repository.AddTask(context.TODO(), invalidTask)

	suite.Error(err)
	suite.Empty(task.ID)
}


func (suite *TaskRepositorySuite) TestGetAllTasks_Empty() {
    tasks, err := suite.repository.GetAllTasks(context.TODO())

    suite.NoError(err, "there is no error")
    suite.Empty(tasks)
}

func (suite *TaskRepositorySuite) TestGetAllTasks_Positive() {
    testTask1 := &domain.Task{
        Title:       "Test Task 1",
        Description: "Test Description 1",
        DueDate:     time.Now(),
        Status:      "pending",
    }
	testTask2 := &domain.Task{
        Title:       "Test Task 2",
        Description: "Test Description 2",
        DueDate:     time.Now(),
        Status:      "pending",
    }
	testTask3 := &domain.Task{
        Title:       "Test Task",
        Description: "Test Description",
        DueDate:     time.Now(),
        Status:      "pending",
    }
    _, err := suite.repository.AddTask(context.TODO(), testTask1)
    suite.NoError(err, "no error on valid input")
	_, err = suite.repository.AddTask(context.TODO(), testTask2)
    suite.NoError(err, "no error on valid input")
	_, err = suite.repository.AddTask(context.TODO(), testTask3)
    suite.NoError(err, "no error on valid input")


    tasks, err := suite.repository.GetAllTasks(context.TODO())

    suite.NoError(err, "no error on getting all tasks when the collection is not empty")
    suite.Len(tasks, 3)
}


func (suite *TaskRepositorySuite) TestGetTaskByID_Positive(){
	testTask := &domain.Task{
        Title:       "Test Task",
        Description: "Test Description",
        DueDate:     time.Now(),
        Status:      "pending",
    }
    newTask, err := suite.repository.AddTask(context.TODO(), testTask)
    suite.NoError(err)

	id := newTask.ID

	task, err := suite.repository.GetTaskByID(context.TODO(), id)
	suite.NoError(err)
    suite.Equal(testTask.Title, task.Title)
}

func (suite *TaskRepositorySuite) TestGetTaskByID_Negative() {

	task, err := suite.repository.GetTaskByID(context.TODO(), primitive.NewObjectID())

	suite.Error(err)
	suite.Empty(task)
}

func (suite *TaskRepositorySuite) TestUpdateTask_Positive(){
	test_task1 := &domain.Task{
		Title: "updated task 1",
		Description: "updating task 1",
		Status: "In progress",
	}


	insertedTask1, err := suite.repository.AddTask(context.TODO(), test_task1)
	suite.NoError(err)

	update_test_task := &domain.Task{
		Description: "updated",
		Status: "Done",
	}
	id := insertedTask1.ID

	updated, err := suite.repository.UpdateTask(context.TODO(), id, update_test_task)
	suite.NoError(err, "no error on valid input")
	suite.NotNil(updated)
	suite.Equal(test_task1.Title, updated.Title)
	suite.Equal(update_test_task.Description, updated.Description)
	suite.Equal(update_test_task.Status, updated.Status)

}

func (suite *TaskRepositorySuite) TestDeleteTask_Postitive(){
	testTask := &domain.Task{
        Title:       "Test Task",
        Description: "Test Description",
        DueDate:     time.Now(),
        Status:      "pending",
    }
    newTask, err := suite.repository.AddTask(context.TODO(), testTask)
    suite.NoError(err)

	id := newTask.ID

	err = suite.repository.DeleteTask(context.TODO(), id)
	suite.NoError(err) 


}


func (suite *TaskRepositorySuite) TestDeleteTask_Negative(){
	testTask := &domain.Task{
        Title:       "Test Task",
        Description: "Test Description",
        DueDate:     time.Now(),
        Status:      "pending",
    }
    newTask, err := suite.repository.AddTask(context.TODO(), testTask)
    suite.NoError(err)

	id := newTask.ID

	err = suite.repository.DeleteTask(context.TODO(), id)
	suite.NoError(err) 

	val, err := suite.repository.GetTaskByID(context.TODO(),id)
	suite.Error(err)
	suite.Empty(val)	

}


func TestTaskRepositorySuite(t *testing.T) {
	suite.Run(t, new(TaskRepositorySuite))
}

