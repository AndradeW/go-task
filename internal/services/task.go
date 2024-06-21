package services

import "github.com/andradew/go/tasks/internal/dtos"

type Service struct {
	repository repository
}

func NewService(repository repository) *Service {
	return &Service{repository: repository}
}

type repository interface {
	GetAllTasks() ([]dtos.Task, error)
	GetTaskByID(ID int) (dtos.Task, error)
	CreateTask(task dtos.Task) (dtos.Task, error)
}

func (s *Service) GetAllTask() ([]dtos.Task, error) {
	return s.repository.GetAllTasks()
}

func (s *Service) GetTaskByID(ID int) (dtos.Task, error) {
	return s.repository.GetTaskByID(ID)
}

func (s *Service) CreateTask(task dtos.Task) (dtos.Task, error) {
	return s.repository.CreateTask(task)
}
