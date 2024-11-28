package routes

import (
	"github.com/MaharoofRashi/task-manager/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.RouterGroup, taskHandler *handlers.TaskHandler) {
	router.GET("", taskHandler.GetTasks)
	router.POST("", taskHandler.CreateTask)
	router.GET("/:id", taskHandler.GetTaskByID)
	router.PUT("/:id", taskHandler.UpdateTask)
	router.DELETE("/:id", taskHandler.DeleteTask)
}
