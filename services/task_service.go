package services

import (
	"github.com/MaharoofRashi/task-manager/models"
	"github.com/MaharoofRashi/task-manager/storage"
	"github.com/google/uuid"
)

func GetAllTasks() []models.Task {
	return storage.GetTasks()
}

func CreateTask(task models.Task) models.Task {
	task.ID = uuid.New().String()
	storage.AddTask(task)
	return task
}

func GetTaskByID(id string) (models.Task, error) {
	return storage.FindTaskByID(id)
}

func UpdateTask(id string, updatedTask models.Task) (models.Task, error) {
	return storage.UpdateTask(id, updatedTask)
}

func DeleteTask(id string) error {
	return storage.DeleteTask(id)
}
