package controllers

import (
	domain "Task-7/Domain"
	usecases "Task-7/UseCases"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	UserUseCase usecases.UserUseCase
}

type TaskController struct {
	TaskUseCase usecases.TaskUseCase
}

func NewUserController(userUseCase usecases.UserUseCase) *UserController{
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func NewTaskController(taskUseCase usecases.TaskUseCase) *TaskController{
	return &TaskController{
		TaskUseCase: taskUseCase,
	}
}




func (tc *TaskController) GetAllTasks(c *gin.Context){
	tasks, err := tc.TaskUseCase.GetAllTasks(c.Request.Context())
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve tasks", "error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}


func (tc *TaskController) GetTaskByID(c *gin.Context){
	id := c.Param("id")
	
	taskID, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id", "error": err.Error()})
	}
	task, err := tc.TaskUseCase.GetTaskByID(c, taskID)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to fetch task", "error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"task": task})
}



func (tc *TaskController) AddTask(c *gin.Context){
	var newTask domain.Task

	if err := c.ShouldBindJSON(&newTask); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request type", "error": err.Error()})
		return
	}

	responseTask, err := tc.TaskUseCase.AddTask(c, &newTask)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to add task", "error": err.Error()})
		return 
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"added_task": responseTask})
}


func (tc *TaskController) UpdateTask(c *gin.Context){
	id := c.Param("id")
	TaskID, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	var updatedTask domain.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid task", "error": err.Error()})
		return
	}

	responseTask, err := tc.TaskUseCase.UpdateTask(c, TaskID, &updatedTask)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to update task", "error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"updated_task": responseTask})
}


func (tc *TaskController) DeleteTask(c *gin.Context){
	id := c.Param("id")
	TaskID, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	err = tc.TaskUseCase.DeleteTask(c,TaskID)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to delete task", "error": err.Error()})
		return 
	}
	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Task Deleted Successfully!"})
}