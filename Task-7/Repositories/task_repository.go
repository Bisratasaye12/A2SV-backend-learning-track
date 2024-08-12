package repositories

import (
	"Task-7/Domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type TaskRepository interface {
    GetAllTasks(ctx context.Context) ([]domain.Task, error)
    GetTaskByID(ctx context.Context, id primitive.ObjectID) (domain.Task, error)
    AddTask(ctx context.Context, task *domain.Task) (domain.Task, error)
    UpdateTask(ctx context.Context, id primitive.ObjectID, updated_task *domain.Task) (domain.Task, error)
    DeleteTask(ctx context.Context, id primitive.ObjectID) error
}


type mongoTaskRepository struct {
    collection *mongo.Collection
}

func NewMongoTaskRepository(db *mongo.Database) *mongoTaskRepository {
    return &mongoTaskRepository{
        collection: db.Collection("tasks"),
    }
}

func (r *mongoTaskRepository) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
    var tasks []domain.Task
    cursor, err := r.collection.Find(ctx, bson.D{{}})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var task domain.Task
        if err := cursor.Decode(&task); err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }
	fmt.Println(tasks)
    return tasks, cursor.Err()
}

func (r *mongoTaskRepository) GetTaskByID(ctx context.Context, id primitive.ObjectID) (domain.Task, error) {
    var task domain.Task
    err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
    return task, err
}


func (r *mongoTaskRepository) AddTask(ctx context.Context, task *domain.Task) (domain.Task, error) {
	_, err := r.collection.InsertOne(ctx, task)
	return *task, err
}

func (r *mongoTaskRepository) UpdateTask(ctx context.Context, id primitive.ObjectID, updatedTask *domain.Task) (domain.Task, error) {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": updatedTask})
	return *updatedTask, err
}

func (r *mongoTaskRepository) DeleteTask(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}