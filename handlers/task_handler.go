package handlers

import (
	"github.com/MaharoofRashi/task-manager/models"
	"github.com/MaharoofRashi/task-manager/services"
	"github.com/MaharoofRashi/task-manager/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTasks(c *gin.Context) {
	tasks := services.GetAllTasks()
	utils.RespondJSON(c, http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdTask := services.CreateTask(task)
	utils.RespondJSON(c, http.StatusCreated, createdTask)
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := services.GetTaskByID(id)
	if err != nil {
		utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	utils.RespondJSON(c, http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := services.UpdateTask(id, updatedTask)
	if err != nil {
		utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	utils.RespondJSON(c, http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteTask(id)
	if err != nil {
		utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	utils.RespondJSON(c, http.StatusNoContent, nil)
}
