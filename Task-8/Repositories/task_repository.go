package repositories

import (
	"Task-8/Domain"
	"context"
	"fmt"

	// "log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)



type MongoTaskRepository struct {
    Collection *mongo.Collection
}

func NewMongoTaskRepository(db *mongo.Database) *MongoTaskRepository {
    return &MongoTaskRepository{
        Collection: db.Collection("tasks"),
    }
}

func (r *MongoTaskRepository) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
    var tasks []domain.Task
    cursor, err := r.Collection.Find(ctx, bson.D{{}})
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
	
    return tasks, cursor.Err()
}

func (r *MongoTaskRepository) GetTaskByID(ctx context.Context, id primitive.ObjectID) (domain.Task, error) {
    var task domain.Task
    err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
    return task, err
}


func (r *MongoTaskRepository) AddTask(ctx context.Context, task *domain.Task) (domain.Task, error) {
    if task.ID != (primitive.ObjectID{}) {
		return domain.Task{}, fmt.Errorf("task ID should be empty for new tasks")
	}
	insertResult, err := r.Collection.InsertOne(ctx, task)
    if err == nil{
        task.ID = insertResult.InsertedID.(primitive.ObjectID)
    }
	return *task, err
}

func (r *MongoTaskRepository) UpdateTask(ctx context.Context, id primitive.ObjectID, updatedTask *domain.Task) (domain.Task, error) {
    filter :=  bson.M{"_id": id}

    updateFields := bson.D{}

	if updatedTask.Title != "" {
		updateFields = append(updateFields, bson.E{Key: "title", Value: updatedTask.Title})
	}
	if updatedTask.Description != "" {
		updateFields = append(updateFields, bson.E{Key: "description", Value: updatedTask.Description})
	}
	if !updatedTask.DueDate.IsZero() {
		updateFields = append(updateFields, bson.E{Key: "due_date", Value: updatedTask.DueDate})
	}
	if updatedTask.Status != "" {
		updateFields = append(updateFields, bson.E{Key: "status", Value: updatedTask.Status})
	}

	_, err := r.Collection.UpdateOne(ctx,filter, bson.M{"$set": updateFields})
	if err != nil{
        return domain.Task{}, fmt.Errorf("unable to update task")
    }

    var updatedTaskResult domain.Task
	err = r.Collection.FindOne(context.TODO(), filter).Decode(&updatedTaskResult)
	if err != nil {
		return domain.Task{}, fmt.Errorf("unable to retrieve updated task: %v", err)
	}

	return updatedTaskResult, nil
}

func (r *MongoTaskRepository) DeleteTask(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}