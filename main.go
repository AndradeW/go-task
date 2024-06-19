package main

import (
	"github.com/andradew/go/tasks/api"
	"log"
	"net/http"

	"github.com/andradew/go/tasks/config"
	"github.com/andradew/go/tasks/internal/handlers"
	"github.com/andradew/go/tasks/internal/repository"
	"github.com/andradew/go/tasks/internal/services"
)

func main() {
	handler := handlers.NewHandler(
		services.NewService(
			repository.NewRepository(),
		))

	router := api.RegisterRoutes(handler)

	log.Println("Listening on " + config.GetPort())
	log.Fatal(http.ListenAndServe(config.GetPort(), router))
}
