package storage

import (
	"errors"
	"github.com/MaharoofRashi/task-manager/internal/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TaskModel struct {
	gorm.Model
	ID          string `gorm:"primaryKey;type:uuid"`
	Title       string `gorm:"not null"`
	Description string
	Status      string
	UserID      string `gorm:"index"`
}

type PostgresTaskRepo struct {
	db *gorm.DB
}

func NewPostgresTaskRepo(dsn string) (*PostgresTaskRepo, error) {
	// create postgres connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// auto migrate to connect TaskModel to gorm to enable update in real db
	err = db.AutoMigrate(&TaskModel{})
	if err != nil {
		return nil, err
	}

	return &PostgresTaskRepo{db: db}, nil
}

// just to convert to model of taskModel
func toModel(task core.Task) TaskModel {
	return TaskModel{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		UserID:      task.UserID,
	}
}

// just to convert back to core structure
func toCore(model TaskModel) core.Task {
	return core.Task{
		ID:          model.ID,
		Title:       model.Title,
		Description: model.Description,
		Status:      model.Status,
		UserID:      model.UserID,
	}
}

// returns all tasks of given user
func (repo *PostgresTaskRepo) GetAll(userID string) ([]core.Task, error) {

	// find the data and inserted to models slice
	var models []TaskModel

	result := repo.db.Where("user_id = ?", userID).Find(&models)
	if result.Error != nil {
		return nil, result.Error
	}

	// convert to core.Task and return it
	tasks := make([]core.Task, len(models))
	for i, model := range models {
		tasks[i] = toCore(model)
	}
	return tasks, nil
}

func (repo *PostgresTaskRepo) GetByID(userID, id string) (core.Task, error) {
	var model TaskModel

	result := repo.db.Where("user_id = ? AND id = ?", userID, id).Find(&model)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return core.Task{}, errors.New("task not found")
		}
		return core.Task{}, result.Error
	}
	return toCore(model), nil
}

func (repo *PostgresTaskRepo) Create(task core.Task) (core.Task, error) {
	model := toModel(task)
	result := repo.db.Create(&model)
	if result.Error != nil {
		return core.Task{}, result.Error
	}
	return toCore(model), nil
}

func (repo *PostgresTaskRepo) Update(task core.Task) (core.Task, error) {
	model := toModel(task)
	result := repo.db.Where("user_id = ? AND id = ?", task.UserID, task.ID).Updates(&model)

	if result.Error != nil {
		return core.Task{}, result.Error
	}

	if result.RowsAffected == 0 {
		return core.Task{}, errors.New("task not found")
	}
	return task, nil
}

func (repo *PostgresTaskRepo) Delete(userID, id string) error {
	result := repo.db.Where("user_id = ? AND id = ?", userID, id).Delete(&TaskModel{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("task not found")
	}
	return nil
}
