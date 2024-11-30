package main

import (
	"github.com/MaharoofRashi/task-manager/config"
	"github.com/MaharoofRashi/task-manager/internal/di"
	"github.com/MaharoofRashi/task-manager/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	r := gin.Default()

	taskHandler := di.BuildTaskHandler()
	authHandler := di.BuildAuthHandler()
	authMiddleware := di.BuildAuthMiddleware(cfg)

	api := r.Group("/api/v1")
	routes.RegisterAuthRoutes(api.Group("/auth"), authHandler)

	protected := api.Group("/tasks")
	protected.Use(authMiddleware)
	routes.RegisterTaskRoutes(protected, taskHandler)

	r.Run(":8080")
}
