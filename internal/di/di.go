package di

import (
	"github.com/MaharoofRashi/task-manager/config"
	"github.com/MaharoofRashi/task-manager/internal/handlers"
	"github.com/MaharoofRashi/task-manager/internal/repository"
	"github.com/MaharoofRashi/task-manager/internal/storage"
	usecase2 "github.com/MaharoofRashi/task-manager/internal/usecase"
	"github.com/MaharoofRashi/task-manager/middleware"
	"github.com/MaharoofRashi/task-manager/pkg/utils"
	"github.com/gin-gonic/gin"
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
	authUsecase := usecase2.NewAuthUsecase(userRepo, jwtUtil)

	return handlers.NewAuthHandler(authUsecase, jwtUtil)
}

func BuildAuthMiddleware(config *config.Config) gin.HandlerFunc {
	jwtUtil := utils.NewJWTUtil(config.JWTSecret)
	return middleware.JWTMiddleware(jwtUtil)
}
