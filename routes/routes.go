package routes

import (
	"github.com/MaharoofRashi/task-manager/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.GET("", handlers.GetTasks)
		taskRoutes.POST("", handlers.CreateTask)
		taskRoutes.GET("/:id", handlers.GetTaskByID)
		taskRoutes.PUT("/:id", handlers.DeleteTask)
		taskRoutes.DELETE("/:id", handlers.DeleteTask)
	}
}
