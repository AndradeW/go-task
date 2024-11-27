package api

import (
	"net/http"

	"github.com/andradew/go/tasks/internal/handlers"
	"github.com/andradew/go/tasks/internal/middlewares"
)

func RegisterRoutes(handler *handlers.Handler) http.Handler {
	router := http.NewServeMux()

	routerWithGlobalMiddlewares := middlewares.ApplyGlobalMiddlewares(
		router,
		middlewares.LoggingMiddleware,
		middlewares.JsonMiddleware)

	router.HandleFunc("GET /tasks", handler.GetAllTask)
	router.HandleFunc("GET /tasks/{id}", handler.GetTaskByID)
	router.HandleFunc("POST /tasks", handler.CreateTask)
	router.HandleFunc("PUT /tasks/{id}", handler.UpdateTaskByID)
	router.HandleFunc("DELETE /tasks/{id}", handler.DeleteTask)

	return routerWithGlobalMiddlewares
}
