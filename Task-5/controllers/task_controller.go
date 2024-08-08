package controllers

import (
	"Task-5/data"
	"Task-5/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetTasks handles the HTTP GET request to retrieve and return all tasks from the database.
// It calls the data layer to fetch tasks and responds with the list of tasks in JSON format.
// If there is an error fetching the tasks, it responds with an internal server error status.
func GetTasks(c *gin.Context) {
	tasks, err := data.GetTasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}

// GetTask handles the HTTP GET request to retrieve a single task by its ID.
// It extracts the task ID from the request parameters, converts it to an ObjectID,
// and fetches the task from the database. If the task is not found or there is an error,
// it responds with a not found or bad request status, respectively.
// Otherwise, it returns the task in JSON format.
func GetTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Task ID"})
		return
	}
	task, err := data.GetTask(taskID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"task": task})
}

// AddTask handles the HTTP POST request to add a new task to the database.
// It binds the incoming JSON payload to a models.Task struct, inserts the task into the database,
// and responds with the newly created task and a status of created. 
// If there is an error during the binding or insertion, it responds with a bad request or internal server error status.
func AddTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	responseTask, err := data.AddTask(newTask)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, responseTask)
}

// UpdateTask handles the HTTP PUT request to update an existing task in the database.
// It extracts the task ID from the request parameters, binds the incoming JSON payload to a models.Task struct,
// and updates the task in the database. It responds with the updated task and a status of OK.
// If there is an error during binding or updating, it responds with a bad request or internal server error status.
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Task ID"})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	updatedTask, err = data.UpdateTask(taskID, updatedTask)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, updatedTask)
}

// DeleteTask handles the HTTP DELETE request to remove a task from the database by its ID.
// It extracts the task ID from the request parameters and deletes the task from the database.
// If there is an error during deletion or if the ID is invalid, it responds with a bad request or internal server error status.
// If successful, it responds with a no content status and a confirmation message.
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Task ID", "error": err.Error()})
		return
	}
	err = data.DeleteTask(taskID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Task Removed Successfully!"})
}
