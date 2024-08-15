package usecases

import (
	"Task-8/Domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)



type Taskusecase struct {
    TaskRepo domain.TaskRepository
}


func NewTaskUseCase(taskRepo domain.TaskRepository) domain.TaskUseCase {
    return &Taskusecase{
        TaskRepo: taskRepo,
    }
}


func (tu *Taskusecase) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
    ret, err := tu.TaskRepo.GetAllTasks(ctx)
    return ret, err
}


func (tu *Taskusecase) GetTaskByID(ctx context.Context, id primitive.ObjectID) (domain.Task, error) {
    ret, err := tu.TaskRepo.GetTaskByID(ctx, id)
    return ret, err
}


func (tu *Taskusecase) AddTask(ctx context.Context, task *domain.Task) (domain.Task, error) {
    if task.Title == "" || task.Description == ""{
        return domain.Task{}, fmt.Errorf("task title and description cannot be empty")
    }
    ret, err := tu.TaskRepo.AddTask(ctx, task)
    return ret, err
}


func (tu *Taskusecase) UpdateTask(ctx context.Context, id primitive.ObjectID, updatedTask *domain.Task) (domain.Task, error) {
    ret, err := tu.TaskRepo.UpdateTask(ctx, id, updatedTask)
    return ret, err
}

func (tu *Taskusecase) DeleteTask(ctx context.Context, id primitive.ObjectID) error {
    err := tu.TaskRepo.DeleteTask(ctx, id)
    return err
}
