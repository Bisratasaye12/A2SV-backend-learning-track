package main

import (
	controllers "Task-7/Delivery/Controllers"
	"Task-7/Delivery/routers"
	infrastructure "Task-7/Infrastructure"
	repositories "Task-7/Repositories"
	usecases "Task-7/UseCases"

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