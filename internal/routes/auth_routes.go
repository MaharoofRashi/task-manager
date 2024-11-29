package routes

import (
	"github.com/MaharoofRashi/task-manager/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	router.POST("/signup", authHandler.Signup)
	router.POST("/login", authHandler.Login)
}
