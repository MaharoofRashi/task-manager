package di

import (
	"github.com/MaharoofRashi/task-manager/config"
	"github.com/MaharoofRashi/task-manager/internal/handlers"
	"github.com/MaharoofRashi/task-manager/internal/repository"
	"github.com/MaharoofRashi/task-manager/internal/storage"
	usecase2 "github.com/MaharoofRashi/task-manager/internal/usecase"
	"github.com/MaharoofRashi/task-manager/pkg/utils"
)

func BuildTaskHandler() *handlers.TaskHandler {
	var repo repository.TaskRepository = storage.NewInMemoryTaskRepo()
	usecase := usecase2.NewTaskUsecase(repo)
	return handlers.NewTaskHandler(usecase)
}

func BuildAuthHandler() *handlers.AuthHandler {
	cfg := config.LoadConfig()

	jwtUtil := utils.NewJWTUtil(cfg.JWTSecret)

	userRepo := storage.NewInMemoryUserRepo()
	authUsercase := usecase2.NewAuthUsecase(userRepo, jwtUtil)

	return handlers.NewAuthHandler(authUsercase, jwtUtil)
}
