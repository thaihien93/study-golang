package repository

import (
	"todo_app/database"
	"todo_app/models"
)

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	result := database.DB.Find(&tasks)
	return tasks, result.Error
}

func CreateTask(task *models.Task) error {
	return database.DB.Create(task).Error
}
