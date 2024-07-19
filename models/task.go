package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type TaskResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
}

func (t *Task) ToResponse() TaskResponse {
	return TaskResponse{
		ID:          t.ID,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
		Title:       t.Title,
		Description: t.Description,
		Done:        t.Done,
	}
}
