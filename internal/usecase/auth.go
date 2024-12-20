package usecase

import (
	"errors"
	"github.com/MaharoofRashi/task-manager/internal/core"
	"github.com/MaharoofRashi/task-manager/internal/repository"
	"github.com/MaharoofRashi/task-manager/pkg/utils"
	"github.com/google/uuid"
)

type AuthUsecase struct {
	userRepo repository.UserRepository
	jwtUtil  *utils.JWTUtil
}

func NewAuthUsecase(userRepo repository.UserRepository, jwtUtil *utils.JWTUtil) *AuthUsecase {
	return &AuthUsecase{
		userRepo: userRepo,
		jwtUtil:  jwtUtil,
	}
}

func (uc *AuthUsecase) Signup(user core.User) (core.User, error) {
	user.ID = uuid.New().String()
	if err := user.HashPassword(); err != nil {
		return core.User{}, err
	}
	return uc.userRepo.Create(user)
}

func (uc *AuthUsecase) Login(username, password string) (string, error) {
	user, err := uc.userRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}
	if err := user.CheckPassword(password); err != nil {
		return "", errors.New("invalid username or password")
	}
	return uc.jwtUtil.GenerateToken(user.ID)
}
