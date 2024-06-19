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
}

func (s *Service) GetAllTask() (dtos.Tasks, error) {
	return s.repository.GetAllTasks()
}
