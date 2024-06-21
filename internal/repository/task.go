package repository

import (
	"errors"
	"github.com/andradew/go/tasks/internal/dtos"
	"log"
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

func (r *Repository) GetTaskByID(ID int) (dtos.Task, error) {

	if task, ok := r.Storage[ID]; ok {
		return task, nil
	}

	log.Println("Not Found")

	return dtos.Task{}, errors.New("not Found")
}

func (r *Repository) CreateTask(task dtos.Task) (dtos.Task, error) {
	//TODO revisar, cambiar len porque al borrar se sobrescriben los id
	task.ID = len(r.Storage) + 1
	r.Storage[task.ID] = task

	return r.Storage[task.ID], nil
}

func (r *Repository) UpdateTaskByID(ID int, task dtos.Task) (dtos.Task, error) {
	if _, ok := r.Storage[ID]; !ok {
		log.Println("Not Found")
		return dtos.Task{}, errors.New("not Found")
	}

	task.ID = ID
	r.Storage[ID] = task
	return r.Storage[ID], nil
}

func (r *Repository) DeleteTaskByID(ID int) error {

	if _, ok := r.Storage[ID]; !ok {
		log.Println("Not Found")
		return errors.New("not Found")
	}

	delete(r.Storage, ID)

	return nil
}
