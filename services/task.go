package services

import (
	"todo_app/models"
	"todo_app/repository"
)

func GetAllTasks() ([]models.Task, error) {
	return repository.GetAllTasks()
}

func CreateTask(task *models.Task) error {
	return repository.CreateTask(task)
}
