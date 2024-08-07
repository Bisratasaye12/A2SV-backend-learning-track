package models

import "time"

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
   }
   
func (t *Task) IsEmpty() bool {
    return t.ID == "" && t.Title == "" && t.Description == "" && t.DueDate.IsZero() && t.Status == ""
}