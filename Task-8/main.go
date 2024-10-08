package main

import (
	controllers "Task-8/Delivery/Controllers"
	"Task-8/Delivery/routers"
	infrastructure "Task-8/Infrastructure"
	repositories "Task-8/Repositories"
	usecases "Task-8/UseCases"

	"github.com/gin-gonic/gin"
)



func main(){
	r := gin.Default()
	infra := infrastructure.NewInfrastructure()
	dataBase := infrastructure.InitDB("mongodb://localhost:27017")
	taskRepository := repositories.NewMongoTaskRepository(dataBase)
	taskUseCase := usecases.NewTaskUseCase(taskRepository)
	TaskController := controllers.NewTaskController(taskUseCase)
	userRepo:= repositories.NewMongoUserRepository(dataBase)
	userUseCase := usecases.NewUserUseCase(userRepo, infra)
	UserController := controllers.NewUserController(userUseCase)

	routers.InitRouter(TaskController, UserController, r, infra)
	
	r.Run("localhost:8080")
}