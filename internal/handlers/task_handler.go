package handlers

import (
	"github.com/MaharoofRashi/task-manager/internal/core"
	"github.com/MaharoofRashi/task-manager/internal/usecase"
	"github.com/MaharoofRashi/task-manager/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskHandler struct {
	usecase *usecase.TaskUsecase
}

func NewTaskHandler(uc *usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{usecase: uc}
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	userID := c.GetString("userID")
	tasks, err := h.usecase.GetAllTasks(userID)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err)
		return
	}
	utils.RespondJSON(c, http.StatusOK, tasks)
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	userID := c.GetString("userID")
	id := c.Param("id")

	task, err := h.usecase.GetTaskByID(userID, id)
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, err)
		return
	}
	utils.RespondJSON(c, http.StatusOK, task)
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task core.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err)
		return
	}
	task.UserID = c.GetString("userID")
	createdTask, err := h.usecase.CreateTask(task)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, err)
		return
	}
	utils.RespondJSON(c, http.StatusCreated, createdTask)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	userID := c.GetString("userID")
	id := c.Param("id")

	var updatedTask core.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		utils.RespondError(c, http.StatusNotFound, err)
		return
	}

	updatedTask.UserID = userID
	updatedTask.ID = id

	task, err := h.usecase.UpdateTask(updatedTask)
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, err)
		return
	}
	utils.RespondJSON(c, http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("userID")

	err := h.usecase.DeleteTask(userID, id)
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, err)
		return
	}
	utils.RespondJSON(c, http.StatusNoContent, nil)
}
