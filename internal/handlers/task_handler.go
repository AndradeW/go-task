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
	//UpdateTaskByID(ID int) (dtos.Task, error)
}

func (h *Handler) GetAllTask(w http.ResponseWriter, req *http.Request) {

	tasks, err := h.service.GetAllTask()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(&task)
}

func (h *Handler) CreateTask(w http.ResponseWriter, req *http.Request) {

	var task dtos.Task
	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	validate := validator.New()
	err = validate.Struct(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tasks, err := h.service.CreateTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(&tasks)
}

//func (h *Handler) UpdateTask(w http.ResponseWriter, req *http.Request) {
//
//	tasks, err := h.service.GetAllTask()
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	json.NewEncoder(w).Encode(&tasks)
//}
