package controllers

import (
	domain "Task-7/Domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	UserUseCase domain.UserUseCase
}

type TaskController struct {
	TaskUseCase domain.TaskUseCase
}

func NewUserController(userUseCase domain.UserUseCase) *UserController{
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func NewTaskController(taskUseCase domain.TaskUseCase) *TaskController{
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
		return 
	}
	task, err := tc.TaskUseCase.GetTaskByID(c.Request.Context(), taskID)
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

	responseTask, err := tc.TaskUseCase.AddTask(c.Request.Context(), &newTask)
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

	responseTask, err := tc.TaskUseCase.UpdateTask(c.Request.Context(), TaskID, &updatedTask)
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

	err = tc.TaskUseCase.DeleteTask(c.Request.Context(),TaskID)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to delete task", "error": err.Error()})
		return 
	}
	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Task Deleted Successfully!"})
}



func (uc *UserController) Register(c *gin.Context){
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil{
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return	
	}

	regUser, err := uc.UserUseCase.Register(c.Request.Context(), &user)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to register user", "error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"registered_user": regUser})
}


func (uc *UserController) Login(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
	  c.JSON(400, gin.H{"error": "Invalid request payload"})
	  return
	}
  
	token, err := uc.UserUseCase.Login(c.Request.Context(),&user)
	if err != nil{
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User logged in successfully", "token": token})
  }


func (uc *UserController) PromoteUser(c *gin.Context){
	
	userID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	
	requestingUserRole, exists := c.Get("role")
	if !exists || requestingUserRole != "admin" {
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": "Only admins can promote users"})
		return
	}

	err = uc.UserUseCase.PromoteUser(c.Request.Context(), objectID)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User promoted to admin"})
}


