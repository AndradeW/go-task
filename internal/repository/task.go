package repository

import (
	"github.com/andradew/go/tasks/internal/dtos"
)

type Repository struct {
	Storage map[int]dtos.Task
}

func NewRepository() *Repository {
	return &Repository{Storage: map[int]dtos.Task{
		1: {
			ID:      1,
			Name:    "Task One",
			Content: "Some content",
		},
		2: {
			ID:      2,
			Name:    "Task Two",
			Content: "Some content",
		},
	}}
}

func (r *Repository) GetAllTasks() ([]dtos.Task, error) {
	tasks := make([]dtos.Task, 0)
	for _, v := range r.Storage {
		tasks = append(tasks, v)
	}
	return tasks, nil
}
