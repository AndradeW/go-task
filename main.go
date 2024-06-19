package main

import (
	"log"
	"net/http"

	"github.com/andradew/go/tasks/config"
	"github.com/andradew/go/tasks/internal/handlers"
	"github.com/andradew/go/tasks/internal/middlewares"
	"github.com/andradew/go/tasks/internal/repository"
	"github.com/andradew/go/tasks/internal/services"
)

func main() {
	router := http.NewServeMux()

	miHandler := handlers.NewHandler(services.NewService(repository.NewRepository()))

	router.HandleFunc("GET /tasks", miHandler.GetAllTask)
	//router.HandleFunc("GET /tasks/{id}", miHandler.GetTaskByID)

	routerWithLogging := middlewares.LoggingMiddleware(router)

	log.Println("Listening on " + config.PORT)
	log.Fatal(http.ListenAndServe(config.PORT, routerWithLogging))
}
