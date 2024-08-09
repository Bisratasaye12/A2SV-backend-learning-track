package router

import (
	"Task-6/controllers"
	"Task-6/middleware"

	"github.com/gin-gonic/gin"
)

// InitRouter initializes the HTTP routes for task management using the provided Gin engine.
// Parameters:
//   - r: The Gin engine to configure with routes.

func InitRouter(r *gin.Engine) {
	// Task routes group
	taskRoutes := r.Group("/tasks", middleware.AuthMiddleware("user","admin"))
	{
		taskRoutes.GET("/", controllers.GetTasks)
		taskRoutes.GET("/:id", controllers.GetTask)
		taskRoutes.POST("/", middleware.AuthMiddleware("admin"), controllers.AddTask)
		taskRoutes.PUT("/:id", middleware.AuthMiddleware("admin"), controllers.UpdateTask)
		taskRoutes.DELETE("/:id", middleware.AuthMiddleware("admin"), controllers.DeleteTask)
	}

	// Auth routes group
	authRoutes := r.Group("/")
	{
		authRoutes.POST("/register", controllers.Register)
		authRoutes.POST("/login", controllers.Login)
	}

	// User routes group
	userRoutes := r.Group("/users", middleware.AuthMiddleware("user", "admin"))
	{
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.POST("/promote/:id", middleware.AuthMiddleware("admin"), controllers.PromoteUser)
	}
}
