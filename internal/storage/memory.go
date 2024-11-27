package storage

import (
	"errors"
	"github.com/MaharoofRashi/task-manager/internal/core"
	"sync"
)

type InMemoryTaskRepo struct {
	mu    sync.Mutex
	tasks map[string][]core.Task
}

func NewInMemoryTaskRepo() *InMemoryTaskRepo {
	return &InMemoryTaskRepo{
		tasks: make(map[string][]core.Task),
	}
}

func (repo *InMemoryTaskRepo) GetAll(userID string) ([]core.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	return repo.tasks[userID], nil
}

func (repo *InMemoryTaskRepo) GetByID(userID string, id string) (core.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	for _, task := range repo.tasks[userID] {
		if id == task.ID {
			return task, nil
		}
	}
	return core.Task{}, errors.New("task not found")
}

func (repo *InMemoryTaskRepo) Create(task core.Task) (core.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.tasks[task.UserID] = append(repo.tasks[task.UserID], task)
	return task, nil
}

func (repo *InMemoryTaskRepo) Update(task core.Task) (core.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	for i, t := range repo.tasks[task.UserID] {
		if t.ID == task.ID {
			repo.tasks[task.UserID][i] = task
			return task, nil
		}
	}
	return core.Task{}, errors.New("task not found")
}

func (repo *InMemoryTaskRepo) Delete(userID, id string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	for i, task := range repo.tasks[userID] {
		if task.ID == id {
			repo.tasks[userID] = append(repo.tasks[userID][:i], repo.tasks[userID][i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
