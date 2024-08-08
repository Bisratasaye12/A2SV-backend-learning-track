package controllers

import (
	"Task-5/data"
	"Task-5/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetTasks retrieves and returns all tasks from the database.
func GetTasks(c *gin.Context) {
	tasks, err := data.GetTasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}


func GetTask( c *gin.Context){
	id := c.Param("id")
	taskID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Task ID"})
		return
	}
	task, err := data.GetTask(taskID)
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"task": task})
}


func AddTask(c *gin.Context){
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	response_task, err := data.AddTask(newTask)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	c.IndentedJSON(http.StatusCreated, response_task)
}



func UpdateTask(c *gin.Context){
	id := c.Param("id")
	taskID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Task ID"})
		return
	}

	var updated_task models.Task
	if err := c.ShouldBindJSON(&updated_task); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
	}

	updated_task, err = data.UpdateTask(taskID, updated_task)
	if err != nil{
		updated_task.ID = taskID
		c.IndentedJSON(http.StatusCreated, updated_task)
	}
	c.IndentedJSON(http.StatusOK, updated_task)
}


func DeleteTask(c *gin.Context){
	id := c.Param("id")
	taskID, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid taskID", "error": err.Error()})
	}
	data.DeleteTask(taskID)
	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Task Removed Successfully!"})
}



