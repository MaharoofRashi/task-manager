package usecase

import (
	"github.com/MaharoofRashi/task-manager/internal/core"
	"github.com/MaharoofRashi/task-manager/internal/repository"
	"github.com/google/uuid"
)

type TaskUsecase struct {
	repo repository.TaskRepository
}

func NewTaskUsecase(repo repository.TaskRepository) *TaskUsecase {
	return &TaskUsecase{repo: repo}
}

func (uc *TaskUsecase) GetAllTasks(userID string) ([]core.Task, error) {
	return uc.repo.GetAll(userID)
}

func (uc *TaskUsecase) CreateTask(task core.Task) (core.Task, error) {
	if err := task.Validate(); err != nil {
		return core.Task{}, err
	}
	task.ID = uuid.New().String()
	return uc.repo.Create(task)
}

func (uc *TaskUsecase) GetTaskByID(userID string, id string) (core.Task, error) {
	return uc.repo.GetByID(userID, id)
}

func (uc *TaskUsecase) UpdateTask(task core.Task) (core.Task, error) {
	if err := task.Validate(); err != nil {
		return core.Task{}, err
	}
	return uc.repo.Update(task)
}

func (uc *TaskUsecase) DeleteTask(userID string, id string) error {
	return uc.repo.Delete(userID, id)
}
