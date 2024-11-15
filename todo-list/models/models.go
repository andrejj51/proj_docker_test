// cmd/main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/andrejj51/proj_docker_test/database"
	"github.com/andrejj51/proj_docker_test/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	database.ConnectDB()

	r := chi.NewRouter()

	r.Get("/tasks", handlers.GetTasks)
	r.Post("/tasks", handlers.PostTask)
	r.Get("/tasks/{id}", handlers.GetTaskByID)
	r.Delete("/tasks/{id}", handlers.DeleteTask)

	if err := http.ListenAndServe(":3000", r); err != nil {
		fmt.Printf("Start server error: %s", err.Error())
		return
	}
}