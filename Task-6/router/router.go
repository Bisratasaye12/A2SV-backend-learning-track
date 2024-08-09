package router

import (
	"Task-6/controllers"
	"github.com/gin-gonic/gin"
)


// InitRouter initializes the HTTP routes for task management using the provided Gin engine.
// Parameters:
//   - r: The Gin engine to configure with routes.
func InitRouter(r *gin.Engine){
	r.GET("/tasks", controllers.GetTasks)
	r.GET("/tasks/:id", controllers.GetTask)
	r.POST("/tasks", controllers.AddTask)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)

}