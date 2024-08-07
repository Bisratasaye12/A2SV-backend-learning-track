package controllers

import (
	"Task-4/data"
	"Task-4/models"
	"math"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// utility functions
func giveId() string {
	randomID := rand.Intn(math.MaxInt)
	return strconv.Itoa(randomID)
}

// main functions
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
	newTask.ID = giveId()
	if err := c.ShouldBindJSON(&newTask); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
	}

	data.AddTask(newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func UpdateTask(c *gin.Context){
	id := c.Param("id")
	var updated_task models.Task
	if err := c.ShouldBindJSON(&updated_task); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
	}

	updated_task, err := data.UpdateTask(id, updated_task)
	if err != nil{
		updated_task.ID = id
	}
	c.IndentedJSON(http.StatusOK, updated_task)
}

func DeleteTask(c *gin.Context){
	id := c.Param("id")
	data.DeleteTask(id)
	c.IndentedJSON(204, gin.H{"message": "Task Removed Successfully!"})
}



