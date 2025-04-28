package service

import (
	"github.com/google/uuid"
	"stage-1/internal/internal"
	"stage-1/internal/repository"
)

type TaskService interface {
	CreateTask(req internal.TaskResponse) (internal.TaskRequest, error)
	GetAllTasks() ([]internal.TaskRequest, error)
	GetTaskByID(id string) (internal.TaskRequest, error)
	UpdateTask(id string, req internal.TaskResponse) (internal.TaskRequest, error)
	DeleteTask(id string) error
}

type tsService struct {
	repo repository.TaskRepository
}

func NewTaskService(r repository.TaskRepository) TaskService {
	return &tsService{repo: r}
}

func (s *tsService) CreateTask(req internal.TaskResponse) (internal.TaskRequest, error) {
	newTask := repository.Tasks{
		ID:     uuid.NewString(),
		Task:   req.Task,
		IsDone: req.IsDone,
	}

	if err := s.repo.CreateTask(newTask); err != nil {
		return internal.TaskRequest{}, err
	}

	ts := internal.TaskRequest{
		ID:     newTask.ID,
		Task:   newTask.Task,
		IsDone: newTask.IsDone,
	}

	return ts, nil
}

func (s *tsService) GetAllTasks() ([]internal.TaskRequest, error) {
	newTasks, err := s.repo.GetAllTasks()
	ts := make([]internal.TaskRequest, len(newTasks))
	for i, _ := range newTasks {
		ts[i].ID = newTasks[i].ID
		ts[i].Task = newTasks[i].Task
		ts[i].IsDone = newTasks[i].IsDone
	}
	return ts, err
}

func (s *tsService) GetTaskByID(id string) (internal.TaskRequest, error) {
	newTasks, err := s.repo.GetTaskByID(id)
	ts := internal.TaskRequest{
		ID:     newTasks.ID,
		Task:   newTasks.Task,
		IsDone: newTasks.IsDone,
	}
	return ts, err
}

func (s *tsService) UpdateTask(id string, req internal.TaskResponse) (internal.TaskRequest, error) {
	newTasks, err := s.repo.GetTaskByID(id)
	if err != nil {
		return internal.TaskRequest{}, err
	}

	newTasks.Task = req.Task
	newTasks.IsDone = req.IsDone

	if err := s.repo.UpdateTask(newTasks); err != nil {
		return internal.TaskRequest{}, err
	}

	ts := internal.TaskRequest{
		ID:     newTasks.ID,
		Task:   newTasks.Task,
		IsDone: newTasks.IsDone,
	}
	return ts, err
}

func (s *tsService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
