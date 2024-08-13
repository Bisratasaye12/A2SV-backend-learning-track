package main

import (
	controllers "Task-8/Delivery/Controllers"
	"Task-8/Delivery/routers"
	infrastructure "Task-8/Infrastructure"
	repositories "Task-8/Repositories"
	usecases "Task-8/UseCases"

	"github.com/gin-gonic/gin"
)

func init(){
	infrastructure.InitDB("mongodb://localhost:27017")
}

func main(){
	r := gin.Default()

	dataBase := infrastructure.Database
	taskRepository := repositories.NewMongoTaskRepository(dataBase)
	taskUseCase := usecases.NewTaskUseCase(taskRepository)
	TaskController := controllers.NewTaskController(taskUseCase)
	userRepo:= repositories.NewMongoUserRepository(dataBase)
	userUseCase := usecases.NewUserUseCase(userRepo)
	UserController := controllers.NewUserController(userUseCase)

	routers.InitRouter(TaskController, UserController, r)
	
	r.Run("localhost:8080")
}