package handlers

import (
	"encoding/json"
	"github.com/andradew/go/tasks/internal/dtos"
	"net/http"
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
	GetAllTask() (dtos.Tasks, error)
}

func (h *Handler) GetAllTask(w http.ResponseWriter, req *http.Request) {

	tasks, err := h.service.GetAllTask()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(&tasks)
}

//func (h handler) GetTaskByID(w http.ResponseWriter, req *http.Request) {
//	idString := req.URL.Query().Get("id")
//	id, err := strconv.Atoi(idString)
//	if err != nil {
//		return
//	}
//	var task dtos.Task
//	for _, eachTask := range dtos.Tasks {
//		if eachTask.ID == id {
//			task = eachTask
//			break
//		}
//	}
//	w.Header().Add("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(&task)
//}
