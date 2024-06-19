package handlers

import (
	"encoding/json"
	"github.com/andradew/go/tasks/internal/dtos"
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
