// usecases/task_usecases.go
package usecases

import (
    "context"
    "Task-8/Domain"
    "Task-8/Repositories"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskUseCase defines the interface for task-related use cases.
type TaskUseCase interface {
    GetAllTasks(ctx context.Context) ([]domain.Task, error)
    GetTaskByID(ctx context.Context, id primitive.ObjectID) (domain.Task, error)
    AddTask(ctx context.Context, task *domain.Task) (domain.Task, error)
    UpdateTask(ctx context.Context, id primitive.ObjectID, updatedTask *domain.Task) (domain.Task, error)
    DeleteTask(ctx context.Context, id primitive.ObjectID) error
}


type taskUseCase struct {
    TaskRepo repositories.TaskRepository
}


func NewTaskUseCase(taskRepo repositories.TaskRepository) TaskUseCase {
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
