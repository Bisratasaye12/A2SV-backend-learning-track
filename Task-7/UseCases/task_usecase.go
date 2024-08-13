// usecases/task_usecases.go
package usecases

import (
    "context"
    "Task-7/Domain"
    "go.mongodb.org/mongo-driver/bson/primitive"
)



type taskUseCase struct {
    TaskRepo domain.TaskRepository
}


func NewTaskUseCase(taskRepo domain.TaskRepository) domain.TaskUseCase {
    return &taskUseCase{
        TaskRepo: taskRepo,
    }
}


func (tu *taskUseCase) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
    return tu.TaskRepo.GetAllTasks(ctx)
}


func (tu *taskUseCase) GetTaskByID(ctx context.Context, id primitive.ObjectID) (domain.Task, error) {
    return tu.TaskRepo.GetTaskByID(ctx, id)
}


func (tu *taskUseCase) AddTask(ctx context.Context, task *domain.Task) (domain.Task, error) {
    return tu.TaskRepo.AddTask(ctx, task)
}


func (tu *taskUseCase) UpdateTask(ctx context.Context, id primitive.ObjectID, updatedTask *domain.Task) (domain.Task, error) {
    return tu.TaskRepo.UpdateTask(ctx, id, updatedTask)
}

func (tu *taskUseCase) DeleteTask(ctx context.Context, id primitive.ObjectID) error {
    return tu.TaskRepo.DeleteTask(ctx, id)
}
