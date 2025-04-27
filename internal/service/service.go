package service

import "github.com/google/uuid"

type TaskService interface {
	CreateTask(task string, isDone bool) (RequestBody, error)
	GetAllTasks() ([]RequestBody, error)
	GetTaskByID(id string) (RequestBody, error)
	UpdateTask(id, task string, isDone bool) (RequestBody, error)
	DeleteTask(id string) error
}

type tsService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &tsService{repo: r}
}

func (s *tsService) CreateTask(task string, isDone bool) (RequestBody, error) {
	ts := RequestBody{
		ID:     uuid.NewString(),
		Task:   task,
		IsDone: isDone,
	}

	if err := s.repo.CreateTask(ts); err != nil {
		return RequestBody{}, err
	}

	return ts, nil
}

func (s *tsService) GetAllTasks() ([]RequestBody, error) {
	return s.repo.GetAllTasks()
}

func (s *tsService) GetTaskByID(id string) (RequestBody, error) {
	return s.repo.GetTaskByID(id)
}

func (s *tsService) UpdateTask(id, task string, isDone bool) (RequestBody, error) {
	ts, err := s.repo.GetTaskByID(id)
	if err != nil {
		return RequestBody{}, err
	}

	ts.Task = task
	ts.IsDone = isDone

	if err := s.repo.UpdateTask(ts); err != nil {
		return RequestBody{}, err
	}

	return ts, nil
}

func (s *tsService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
