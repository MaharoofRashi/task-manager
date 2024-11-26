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
	task.ID = int(uuid.New().ID())
	storage.AddTask(task)
	return task
}

func GetTaskByID(id int) (models.Task, error) {
	return storage.FindTaskByID(id)
}

func UpdateTask(id int, updatedTask models.Task) (models.Task, error) {
	return storage.UpdateTask(id, updatedTask)
}

func DeleteTask(id int) error {
	return storage.DeleteTask(id)
}
