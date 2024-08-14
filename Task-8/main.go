package main

import (
	controllers "Task-8/Delivery/Controllers"
	"Task-8/Delivery/routers"
	infrastructure "Task-8/Infrastructure"
	repositories "Task-8/Repositories"
	usecases "Task-8/UseCases"

	"github.com/gin-gonic/gin"
)

var(
	infra *infrastructure.Infrastruct
)
func init(){
	infra = infrastructure.NewInfrastructure()
	infra.InitDB("mongodb://localhost:27017")
}

func main(){
	r := gin.Default()

	dataBase := infra.Database
	taskRepository := repositories.NewMongoTaskRepository(dataBase)
	taskUseCase := usecases.NewTaskUseCase(taskRepository)
	TaskController := controllers.NewTaskController(taskUseCase)
	userRepo:= repositories.NewMongoUserRepository(dataBase)
	userUseCase := usecases.NewUserUseCase(userRepo, infra)
	UserController := controllers.NewUserController(userUseCase)

	routers.InitRouter(TaskController, UserController, r, infra)
	
	r.Run("localhost:8080")
}