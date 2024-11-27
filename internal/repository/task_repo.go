package repository

import "github.com/MaharoofRashi/task-manager/internal/core"

type TaskRepository interface {
	GetAll(userID string) ([]core.Task, error)
	GetByID(userID string, id string) (core.Task, error)
	Create(task core.Task) (core.Task, error)
	Update(task core.Task) (core.Task, error)
	Delete(userID, id string) error
}
