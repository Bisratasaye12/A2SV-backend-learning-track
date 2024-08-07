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

// giveId generates a random ID for a new task.
// Returns:
//   - string: A random ID in string format.
func giveId() string {
	randomID := rand.Intn(math.MaxInt)
	return strconv.Itoa(randomID)
}

// main functions


// GetTasks handles the HTTP GET request to retrieve all tasks.
// Responds with the list of tasks and an appropriate HTTP status code.
// Parameters:
//   - c: The Gin context.
func GetTasks(c *gin.Context){
	tasks, err := data.GetTasks()
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "No available Tasks"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}


// GetTask handles the HTTP GET request to retrieve a specific task by its ID.
// Parameters:
//   - c: The Gin context.
// Responds with the task details and an appropriate HTTP status code.
func GetTask( c *gin.Context){
	id := c.Param("id")
	task, err := data.GetTask(id)
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task Not Found"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"task": task})
}


// AddTask handles the HTTP POST request to add a new task.
// Parameters:
//   - c: The Gin context.
// Responds with the newly created task and an HTTP 201 status code.
func AddTask(c *gin.Context){
	var newTask models.Task
	newTask.ID = giveId()
	if err := c.ShouldBindJSON(&newTask); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
	}

	data.AddTask(newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}



// UpdateTask handles the HTTP PUT request to update an existing task by its ID.
// Parameters:
//   - c: The Gin context.
// Responds with the updated task and an appropriate HTTP status code.
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



// DeleteTask handles the HTTP DELETE request to remove a task by its ID.
// Parameters:
//   - c: The Gin context.
// Responds with a success message and an HTTP 204 status code.
func DeleteTask(c *gin.Context){
	id := c.Param("id")
	data.DeleteTask(id)
	c.IndentedJSON(204, gin.H{"message": "Task Removed Successfully!"})
}



