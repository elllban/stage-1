package service

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(ts RequestBody) error
	GetAllTasks() ([]RequestBody, error)
	GetTaskByID(id string) (RequestBody, error)
	UpdateTask(ts RequestBody) error
	DeleteTask(id string) error
}

type tsRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &tsRepository{db: db}
}

func (r *tsRepository) CreateTask(ts RequestBody) error {
	return r.db.Create(&ts).Error
}

func (r *tsRepository) GetAllTasks() ([]RequestBody, error) {
	var tasks []RequestBody
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *tsRepository) GetTaskByID(id string) (RequestBody, error) {
	var ts RequestBody
	err := r.db.First(&ts, "id = ?", id).Error
	return ts, err
}

func (r *tsRepository) UpdateTask(ts RequestBody) error {
	return r.db.Save(&ts).Error
}

func (r *tsRepository) DeleteTask(id string) error {
	return r.db.Delete(&RequestBody{}, "id = ?", id).Error
}
