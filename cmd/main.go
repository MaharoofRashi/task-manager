package main

import (
	"github.com/MaharoofRashi/task-manager/config"
	"github.com/MaharoofRashi/task-manager/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
