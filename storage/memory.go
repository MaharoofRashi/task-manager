package storage

import (
	"errors"
	"github.com/MaharoofRashi/task-manager/models"
)

var tasks = []models.Task{}

func GetTasks() []models.Task {
	return tasks
}

func AddTask(task models.Task) {
	tasks = append(tasks, task)
}

func FindTaskByID(id string) (models.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func UpdateTask(id string, updatedTask models.Task) (models.Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			updatedTask.ID = id
			tasks[i] = updatedTask
			return updatedTask, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func DeleteTask(id string) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
