package controllers

import (
	usecases "Task-7/UseCases"
	"net/http"

	"github.com/gin-gonic/gin"
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