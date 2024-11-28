package storage

import (
	"errors"
	"github.com/MaharoofRashi/task-manager/internal/core"
	"sync"
)

type InMemoryUserRepo struct {
	mu    sync.Mutex
	users map[string]core.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{users: make(map[string]core.User)}
}

func (repo *InMemoryUserRepo) Create(user core.User) (core.User, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.users[user.Username]; exists {
		return core.User{}, errors.New("username already exists")
	}
	repo.users[user.Username] = user
	return user, nil
}

func (repo *InMemoryUserRepo) FindByUsername(username string) (core.User, error) {
	repo.mu.Lock()
	defer repo.mu.Lock()

	user, exists := repo.users[username]
	if !exists {
		return core.User{}, errors.New("user not found")
	}
	return user, nil
}
