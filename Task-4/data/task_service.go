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


func GetTasks() ([]models.Task, error) {
	if len(tasks) == 0{
		return nil, fmt.Errorf("no available task")
	}

	return tasks, nil
}

func GetTask(id string) (models.Task, error) {
	for _, task := range tasks{
		if id == task.ID{
			return task, nil
		}
	}
	return models.Task{}, fmt.Errorf("task not found")
}

func AddTask(task models.Task) (models.Task, error){
	if task.IsEmpty(){
		return models.Task{}, fmt.Errorf("task is empty")
	}
	tasks = append(tasks, task)
	return task, nil
}

func UpdateTask(updated_task models.Task) models.Task{
	for i, task := range tasks{
		if updated_task.ID == task.ID{
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
			return tasks[i]
		}
	}
	tasks = append(tasks, updated_task)
	return updated_task
}


func DeleteTask(id string) error{

	for i,task := range tasks{
		if task.ID == id{
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task not found")
}