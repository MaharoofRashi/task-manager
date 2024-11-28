package repository

import "github.com/MaharoofRashi/task-manager/internal/core"

type UserRepository interface {
	Create(user core.User) (core.User, error)
	FindByUsername(username string) (core.User, error)
}
