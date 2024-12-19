package storage

import (
	"errors"
	"github.com/MaharoofRashi/task-manager/internal/core"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID       string `gorm:"primaryKey;type:uuid"`
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

type PostgresUserRepo struct {
	db *gorm.DB
}

func NewPostgresUserRepo(db *gorm.DB) (*PostgresTaskRepo, error) {
	err := db.AutoMigrate(&UserModel{})
	if err != nil {
		return nil, err
	}
	return &PostgresTaskRepo{db: db}, nil
}

func toUserModel(user core.User) UserModel {
	return UserModel{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}
}

func toCoreUser(model UserModel) core.User {
	return core.User{
		ID:       model.ID,
		Username: model.Username,
		Password: model.Password,
	}
}

func (repo *PostgresUserRepo) Create(user core.User) (core.User, error) {
	model := toUserModel(user)
	result := repo.db.Create(&model)

	if result.Error != nil {
		if repo.db.Where("username: ?", user.Username).First(&UserModel{}).Error == nil {
			return core.User{}, errors.New("username already exists")
		}
		return core.User{}, result.Error
	}
	return toCoreUser(model), nil
}

func (repo *PostgresUserRepo) FindByUsername(username string) (core.User, error) {
	var model UserModel
	result := repo.db.Where("username = ?", username).First(&model)
	if result.Error != nil {
		if errors.Is(gorm.ErrRecordNotFound, result.Error) {
			return core.User{}, errors.New("user not found")
		}
		return core.User{}, result.Error
	}
	return toCoreUser(model), nil
}
