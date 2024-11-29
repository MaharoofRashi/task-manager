package handlers

import (
	"github.com/MaharoofRashi/task-manager/internal/core"
	"github.com/MaharoofRashi/task-manager/internal/usecase"
	"github.com/MaharoofRashi/task-manager/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	usecase *usecase.AuthUsecase
	jwtUtil *utils.JWTUtil
}

func NewAuthHandler(uc *usecase.AuthUsecase, jwtUtil *utils.JWTUtil) *AuthHandler {
	return &AuthHandler{
		usecase: uc,
		jwtUtil: jwtUtil,
	}
}

func (h *AuthHandler) Signup(c *gin.Context) {
	var user core.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err)
		return
	}
	createdUser, err := h.usecase.Signup(user)
	if err != nil {
		utils.RespondError(c, http.StatusConflict, err)
		return
	}
	utils.RespondJSON(c, http.StatusCreated, createdUser)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&creds); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err)
		return
	}
	userID, err := h.usecase.Login(creds.Username, creds.Password)
	if err != nil {
		utils.RespondError(c, http.StatusUnauthorized, err)
		return
	}
	token, err := h.jwtUtil.GenerateToken(userID)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err)
		return
	}
	utils.RespondJSON(c, http.StatusOK, token)
}
