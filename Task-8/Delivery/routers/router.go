package routers

import (
	controllers "Task-8/Delivery/Controllers"
	infrastructure "Task-8/Infrastructure"

	"github.com/gin-gonic/gin"
)

// InitRouter initializes the HTTP routes for task management using the provided Gin engine.
// Parameters:
//   - r: The Gin engine to configure with routes.

func InitRouter(tc *controllers.TaskController, uc *controllers.UserController, r *gin.Engine, infrastructure *infrastructure.Infrastruct) {
	// Task routes group
	taskRoutes := r.Group("/tasks", infrastructure.AuthMiddleware("user","admin"))
	{
		taskRoutes.GET("/", tc.GetAllTasks)
		taskRoutes.GET("/:id", tc.GetTaskByID)
		taskRoutes.POST("/", infrastructure.AuthMiddleware("admin"), tc.AddTask)
		taskRoutes.PUT("/:id", infrastructure.AuthMiddleware("admin"), tc.UpdateTask)
		taskRoutes.DELETE("/:id", infrastructure.AuthMiddleware("admin"), tc.DeleteTask)
	}

	// Auth routes group
	authRoutes := r.Group("/")
	{
		authRoutes.POST("/register", uc.Register)
		authRoutes.POST("/login", uc.Login)
	}

	// User routes group
	userRoutes := r.Group("/users", infrastructure.AuthMiddleware("user", "admin"))
	{
		userRoutes.PUT("/promote/:id", infrastructure.AuthMiddleware("admin"), uc.PromoteUser)
	}
}