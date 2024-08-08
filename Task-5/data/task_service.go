package data

import (
	"Task-5/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	dbClient *mongo.Client
	tasks    *mongo.Collection
)

func init() {
	var err error
	dbClient, err = InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	tasks = dbClient.Database("Task_management").Collection("tasks")
}

// GetTasks retrieves all tasks from the database.
// It performs a query to fetch all documents from the "tasks" collection
// and returns them as a slice of models.Task. It returns an error if
// there is an issue with the query or decoding the documents.
func GetTasks() ([]models.Task, error) {
	filter := bson.D{{}}
	var fetchedTasks []models.Task

	cursor, err := tasks.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("unable to access tasks: %v", err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, fmt.Errorf("unable to decode task: %v", err)
		}
		fetchedTasks = append(fetchedTasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return fetchedTasks, nil
}

// GetTask retrieves a single task from the database by its ID.
// It queries the "tasks" collection using the provided ObjectID and
// returns the corresponding task as a models.Task. If the task is not
// found, it returns an error indicating that the task with the specified ID was not found.
func GetTask(id primitive.ObjectID) (models.Task, error) {
	filter := bson.D{{"_id", id}}

	var fetchedTask models.Task
	err := tasks.FindOne(context.TODO(), filter).Decode(&fetchedTask)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, fmt.Errorf("task with ID %s not found", id.Hex())
		}
		return models.Task{}, fmt.Errorf("unable to retrieve task: %v", err)
	}

	return fetchedTask, nil
}

// AddTask inserts a new task into the database.
// It takes a models.Task as input, which should have an empty ID.
// After insertion, it sets the task's ID to the newly generated ObjectID.
// It returns the inserted task with its ID and an error if the insertion fails.
func AddTask(task models.Task) (models.Task, error) {
	if task.ID != (primitive.ObjectID{}) {
		return models.Task{}, fmt.Errorf("task ID should be empty for new tasks")
	}

	insertResult, err := tasks.InsertOne(context.TODO(), task)
	if err != nil {
		return models.Task{}, fmt.Errorf("unable to insert task: %v", err)
	}

	task.ID = insertResult.InsertedID.(primitive.ObjectID)

	return task, nil
}

// UpdateTask updates an existing task in the database.
// It takes the task's ID and an updated models.Task as input.
// It constructs an update document with the fields that are not empty in the updated_task.
// If no fields are specified for update, it returns an error indicating that there are no fields to update.
// After the update, it retrieves and returns the updated task. It returns an error if the update fails or the task cannot be retrieved.
func UpdateTask(id primitive.ObjectID, updated_task models.Task) (models.Task, error) {
	filter := bson.D{{"_id", id}}

	updateFields := bson.D{}

	if updated_task.Title != "" {
		updateFields = append(updateFields, bson.E{"title", updated_task.Title})
	}
	if updated_task.Description != "" {
		updateFields = append(updateFields, bson.E{"description", updated_task.Description})
	}
	if !updated_task.DueDate.IsZero() {
		updateFields = append(updateFields, bson.E{"due_date", updated_task.DueDate})
	}
	if updated_task.Status != "" {
		updateFields = append(updateFields, bson.E{"status", updated_task.Status})
	}

	if len(updateFields) == 0 {
		return models.Task{}, fmt.Errorf("no fields to update")
	}

	update := bson.D{{"$set", updateFields}}

	_, err := tasks.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return models.Task{}, fmt.Errorf("unable to update task: %v", err)
	}

	var updatedTaskResult models.Task
	err = tasks.FindOne(context.TODO(), filter).Decode(&updatedTaskResult)
	if err != nil {
		return models.Task{}, fmt.Errorf("unable to retrieve updated task: %v", err)
	}

	return updatedTaskResult, nil
}

// DeleteTask removes a task from the database by its ID.
// It takes the task's ObjectID as input and deletes the corresponding document
// from the "tasks" collection. It returns an error if the deletion fails.
func DeleteTask(id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}

	_, err := tasks.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("unable to delete task: %v", err)
	}
	return nil
}
