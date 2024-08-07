package router

import (
	"Task-4/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine){
	r.GET("/tasks", controllers.GetTasks)
	r.GET("/tasks/:id", controllers.GetTask)
	r.POST("/tasks", controllers.AddTask)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)

}