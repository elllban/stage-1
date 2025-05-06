package taskService

import (
	"github.com/google/uuid"
	"stage-1/internal/model"
	"stage-1/internal/repository/taskRepository"
)

type TaskService interface {
	CreateTask(res model.TaskResponse) (model.TaskRequest, error)
	GetAllTasks() ([]model.TaskRequest, error)
	GetTaskByID(id string) (model.TaskRequest, error)
	UpdateTask(id string, req model.TaskResponse) (model.TaskRequest, error)
	DeleteTask(id string) error
}

type tsService struct {
	repo taskRepository.TaskRepository
}

func NewTaskService(r taskRepository.TaskRepository) TaskService {
	return &tsService{repo: r}
}

func (s *tsService) CreateTask(res model.TaskResponse) (model.TaskRequest, error) {
	newTask := taskRepository.Tasks{
		ID:     uuid.NewString(),
		Task:   res.Task,
		IsDone: res.IsDone,
		UserID: res.UserID,
	}

	if err := s.repo.CreateTask(newTask); err != nil {
		return model.TaskRequest{}, err
	}

	ts := model.TaskRequest{
		ID:     newTask.ID,
		Task:   newTask.Task,
		IsDone: newTask.IsDone,
		UserID: newTask.UserID,
	}

	return ts, nil
}

func (s *tsService) GetAllTasks() ([]model.TaskRequest, error) {
	newTasks, err := s.repo.GetAllTasks()
	if err != nil {
		return []model.TaskRequest{}, err
	}

	ts := make([]model.TaskRequest, len(newTasks))
	for i := range newTasks {
		ts[i].ID = newTasks[i].ID
		ts[i].Task = newTasks[i].Task
		ts[i].IsDone = newTasks[i].IsDone
		ts[i].UserID = newTasks[i].UserID
	}
	return ts, nil
}

func (s *tsService) GetTaskByID(id string) (model.TaskRequest, error) {
	newTask, err := s.repo.GetTaskByID(id)
	if err != nil {
		return model.TaskRequest{}, err
	}

	ts := model.TaskRequest{
		ID:     newTask.ID,
		Task:   newTask.Task,
		IsDone: newTask.IsDone,
		UserID: newTask.UserID,
	}
	return ts, nil
}

func (s *tsService) UpdateTask(id string, req model.TaskResponse) (model.TaskRequest, error) {
	newTasks, err := s.repo.GetTaskByID(id)
	if err != nil {
		return model.TaskRequest{}, err
	}

	newTasks.Task = req.Task
	newTasks.IsDone = req.IsDone
	newTasks.UserID = req.UserID

	if err := s.repo.UpdateTask(newTasks); err != nil {
		return model.TaskRequest{}, err
	}

	ts := model.TaskRequest{
		ID:     newTasks.ID,
		Task:   newTasks.Task,
		IsDone: newTasks.IsDone,
		UserID: newTasks.UserID,
	}
	return ts, nil
}

func (s *tsService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
