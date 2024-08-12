package routers

import (
	controllers "Task-7/Delivery/Controllers"

	"github.com/gin-gonic/gin"
)

// InitRouter initializes the HTTP routes for task management using the provided Gin engine.
// Parameters:
//   - r: The Gin engine to configure with routes.

func InitRouter(tc *controllers.TaskController, uc *controllers.UserController, r *gin.Engine) {
	// Task routes group
	taskRoutes := r.Group("/tasks",)
	{
		taskRoutes.GET("/", tc.GetAllTasks)
		taskRoutes.GET("/:id", tc.GetTaskByID)
		taskRoutes.POST("/", tc.AddTask)
		taskRoutes.PUT("/:id", tc.UpdateTask)
		taskRoutes.DELETE("/:id", tc.DeleteTask)
	}

	// Auth routes group
	// authRoutes := r.Group("/")
	// {
	// 	authRoutes.POST("/register", uc.Register)
	// 	authRoutes.POST("/login", uc.Login)
	// }

	// User routes group
	// userRoutes := r.Group("/users")
	// {
	// 	userRoutes.GET("/", uc.GetUsers)
	// 	userRoutes.PUT("/promote/:id", uc.PromoteUser)
	// }
}