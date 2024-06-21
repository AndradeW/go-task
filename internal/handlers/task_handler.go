package handlers

import (
	"encoding/json"
	"github.com/andradew/go/tasks/internal/dtos"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type Handler struct {
	service
}

func NewHandler(service service) *Handler {
	return &Handler{
		service: service,
	}
}

type service interface {
	GetAllTask() ([]dtos.Task, error)
	GetTaskByID(ID int) (dtos.Task, error)
	CreateTask(task dtos.Task) (dtos.Task, error)
	UpdateTaskByID(ID int, task dtos.Task) (dtos.Task, error)
	DeleteTaskByID(ID int) error
}

func (h *Handler) GetAllTask(w http.ResponseWriter, req *http.Request) {

	tasks, err := h.service.GetAllTask()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&tasks)
}

func (h *Handler) GetTaskByID(w http.ResponseWriter, req *http.Request) {
	idString := req.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func (h *Handler) CreateTask(w http.ResponseWriter, req *http.Request) {

	var task dtos.Task
	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tasks, err := h.service.CreateTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&tasks)
}

func (h *Handler) UpdateTaskByID(w http.ResponseWriter, req *http.Request) {
	idString := req.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var task dtos.Task
	err = json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskUpdated, err := h.service.UpdateTaskByID(id, task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&taskUpdated)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, req *http.Request) {
	idString := req.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteTaskByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
