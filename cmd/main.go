package main

import (
	"github.com/MaharoofRashi/task-manager/internal/di"
	"github.com/MaharoofRashi/task-manager/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	taskHandler := di.BuildTaskHandler()
	api := r.Group("/api/v1")
	routes.RegisterTaskRoutes(api.Group("/tasks"), taskHandler)

	r.Run(":8080")
}
