package repository

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(ts Tasks) error
	GetAllTasks() ([]Tasks, error)
	GetTaskByID(id string) (Tasks, error)
	UpdateTask(ts Tasks) error
	DeleteTask(id string) error
}

type tsRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &tsRepository{db: db}
}

func (r *tsRepository) CreateTask(ts Tasks) error {
	return r.db.Create(&ts).Error
}

func (r *tsRepository) GetAllTasks() ([]Tasks, error) {
	var tasks []Tasks
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *tsRepository) GetTaskByID(id string) (Tasks, error) {
	var ts Tasks
	err := r.db.First(&ts, "id = ?", id).Error
	return ts, err
}

func (r *tsRepository) UpdateTask(ts Tasks) error {
	return r.db.Save(&ts).Error
}

func (r *tsRepository) DeleteTask(id string) error {
	return r.db.Delete(&Tasks{}, "id = ?", id).Error
}
