package userRepository

import (
	"gorm.io/gorm"
	"stage-1/internal/repository/taskRepository"
)

type UserRepository interface {
	CreateUser(us Users) error
	GetAllUsers() ([]Users, error)
	GetUserByID(id string) (Users, error)
	GetTasksForUser(userId string) ([]taskRepository.Tasks, error)
	UpdateUser(us Users) error
	DeleteUser(id string) error
}

type usRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &usRepository{db: db}
}

func (r *usRepository) CreateUser(us Users) error {
	return r.db.Create(&us).Error
}

func (r *usRepository) GetAllUsers() ([]Users, error) {
	var users []Users
	err := r.db.Find(&users).Error
	return users, err
}

func (r *usRepository) GetUserByID(id string) (Users, error) {
	var us Users
	err := r.db.First(&us, "id = ?", id).Error
	return us, err
}

func (r *usRepository) GetTasksForUser(userId string) ([]taskRepository.Tasks, error) {
	var tasks []taskRepository.Tasks
	err := r.db.Find(&tasks, "user_id = ?", userId).Error
	return tasks, err
}

func (r *usRepository) UpdateUser(us Users) error {
	return r.db.Save(&us).Error
}

func (r *usRepository) DeleteUser(id string) error {
	return r.db.Delete(&Users{}, "id = ?", id).Error
}
