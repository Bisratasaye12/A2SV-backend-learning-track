package controllers

import (
	"Task-4/data"
	"net/http"
	"Task-4/models"
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context){
	tasks, err := data.GetTasks()
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "No available Tasks"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}

func GetTask( c *gin.Context){
	id := c.Param("id")
	task, err := data.GetTask(id)
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task Not Found"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"task": task})
}

func AddTask(c *gin.Context){
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
	}

	data.AddTask(newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func UpdateTask(c *gin.Context){
	var updated_task models.Task
	if err := c.ShouldBindJSON(&updated_task); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
	}

	data.UpdateTask(updated_task)
	c.IndentedJSON(http.StatusOK, updated_task)
}

func DeleteTask(c *gin.Context){
	var toBeRemoved models.Task
	if err := c.ShouldBindJSON(&toBeRemoved); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	data.DeleteTask(toBeRemoved.ID)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task Removed Successfully!"})
}



