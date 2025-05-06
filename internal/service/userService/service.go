package userService

import (
	"github.com/google/uuid"
	"stage-1/internal/model"
	"stage-1/internal/repository/userRepository"
)

type UserService interface {
	CreateUser(res model.UserResponse) (model.UserRequest, error)
	GetAllUsers() ([]model.UserRequest, error)
	GetUserByID(id string) (model.UserRequest, error)
	GetTasksForUser(userId string) ([]model.TaskRequest, error)
	UpdateUser(id string, res model.UserResponse) (model.UserRequest, error)
	DeleteUser(id string) error
}

type usService struct {
	repo userRepository.UserRepository
}

func NewUserService(r userRepository.UserRepository) UserService {
	return &usService{repo: r}
}

func (s *usService) CreateUser(res model.UserResponse) (model.UserRequest, error) {
	newUser := userRepository.Users{
		ID:       uuid.NewString(),
		Email:    res.Email,
		Password: res.Password,
	}

	if err := s.repo.CreateUser(newUser); err != nil {
		return model.UserRequest{}, err
	}

	us := model.UserRequest{
		ID:       newUser.ID,
		Email:    newUser.Email,
		Password: newUser.Password,
	}

	return us, nil
}

func (s *usService) GetAllUsers() ([]model.UserRequest, error) {
	newUsers, err := s.repo.GetAllUsers()
	if err != nil {
		return []model.UserRequest{}, err
	}

	us := make([]model.UserRequest, len(newUsers))
	for i := range newUsers {
		us[i].ID = newUsers[i].ID
		us[i].Email = newUsers[i].Email
		us[i].Password = newUsers[i].Password
	}

	return us, nil
}

func (s *usService) GetUserByID(id string) (model.UserRequest, error) {
	newUser, err := s.repo.GetUserByID(id)
	if err != nil {
		return model.UserRequest{}, err
	}

	us := model.UserRequest{
		ID:       newUser.ID,
		Email:    newUser.Email,
		Password: newUser.Password,
	}

	return us, nil
}

func (s *usService) GetTasksForUser(userId string) ([]model.TaskRequest, error) {
	allTasks, err := s.repo.GetTasksForUser(userId)
	if err != nil {
		return []model.TaskRequest{}, err
	}

	ts := make([]model.TaskRequest, len(allTasks))
	for i := range allTasks {
		ts[i].ID = allTasks[i].ID
		ts[i].Task = allTasks[i].Task
		ts[i].IsDone = allTasks[i].IsDone
		ts[i].UserID = allTasks[i].UserID
	}

	return ts, nil
}

func (s *usService) UpdateUser(id string, res model.UserResponse) (model.UserRequest, error) {
	newUser, err := s.repo.GetUserByID(id)
	if err != nil {
		return model.UserRequest{}, err
	}

	newUser.Email = res.Email
	newUser.Password = res.Password

	if err := s.repo.UpdateUser(newUser); err != nil {
		return model.UserRequest{}, err
	}

	us := model.UserRequest{
		ID:       newUser.ID,
		Email:    newUser.Email,
		Password: newUser.Password,
	}

	return us, nil
}

func (s *usService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}
