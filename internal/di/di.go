package di

import (
	"github.com/MaharoofRashi/task-manager/internal/handlers"
	"github.com/MaharoofRashi/task-manager/internal/repository"
	"github.com/MaharoofRashi/task-manager/internal/storage"
	usecase2 "github.com/MaharoofRashi/task-manager/internal/usecase"
)

func BuildTaskHandler() *handlers.TaskHandler {
	var repo repository.TaskRepository = storage.NewInMemoryTaskRepo()
	usecase := usecase2.NewTaskUsecase(repo)
	return handlers.NewTaskHandler(usecase)
}
