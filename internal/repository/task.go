package repository

import "github.com/andradew/go/tasks/internal/dtos"

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetAllTasks() ([]dtos.Task, error) {
	return dtos.Tasks{
		{
			ID:      1,
			Name:    "Task One",
			Content: "Some content",
		},
	}, nil
}
