package data

import (
	"Task-4/models"
	"fmt"
	"time"
)

// dummy in-memory data
var tasks = []models.Task{
	{ID: "1", Title: "Task Manager Project", Description: "Add/View/Delete Tasks", DueDate: time.Now(), Status: "In Progress"},
	{ID: "2", Title: "Books Management Project", Description: "Add/View/Delete Books", DueDate: time.Now().AddDate(0, 0, -1), Status: "Completed"},
	{ID: "3", Title: "Library Management Project", Description: "Add/View/Delete Books", DueDate: time.Now().AddDate(0, 0, -1), Status: "Completed"},
	{ID: "4", Title: "Task Management Project (with db)", Description: "Add/View/Delete Books", DueDate: time.Now().AddDate(0, 0, -1), Status: "Completed"},

}


// GetTasks retrieves all tasks from the in-memory data store.
// Returns a slice of Task models and an error if no tasks are available.
func GetTasks() ([]models.Task, error) {
	
	if len(tasks) == 0{
		return nil, fmt.Errorf("no available task")
	}

	return tasks, nil
}


// GetTask retrieves a specific task by its ID from the in-memory data store.
// Parameters:
//   - id: The ID of the task to retrieve.
// Returns:
//   - Task: The task with the specified ID.
//   - error: An error if the task is not found.

func GetTask(id string) (models.Task, error) {
	for _, task := range tasks{
		if id == task.ID{
			return task, nil
		}
	}
	return models.Task{}, fmt.Errorf("task not found")
}



// AddTask adds a new task to the in-memory data store.
// Parameters:
//   - task: The Task model to be added.
// Returns:
//   - Task: The added task with all its details.
//   - error: An error if the task is empty.
func AddTask(task models.Task) (models.Task, error){
	if task.IsEmpty(){
		return models.Task{}, fmt.Errorf("task is empty")
	}
	tasks = append(tasks, task)
	return task, nil
}



// UpdateTask updates an existing task in the in-memory data store based on the provided ID.
// Parameters:
//   - id: The ID of the task to update.
//   - updated_task: The Task model containing the updated details.
// Returns:
//   - Task: The updated task.
//   - error: An error if the task to update does not exist.
func UpdateTask(id string, updated_task models.Task) (models.Task, error){
	for i, task := range tasks{
		if id == task.ID{
			if updated_task.Title != ""{
				tasks[i].Title = updated_task.Title
			}
			if updated_task.Description != ""{
				tasks[i].Description = updated_task.Description
			}
			if updated_task.Status != ""{
				tasks[i].Status = updated_task.Status
			}
			tasks[i].DueDate = updated_task.DueDate
			return tasks[i], nil
		}
	}
	tasks = append(tasks, updated_task)
	return updated_task, fmt.Errorf("no such task")
}



// DeleteTask removes a task from the in-memory data store based on the provided ID.
// Parameters:
//   - id: The ID of the task to delete.
func DeleteTask(id string){

	for i,task := range tasks{
		if task.ID == id{
			tasks = append(tasks[:i], tasks[i+1:]...)
			return 
		}
	}
}