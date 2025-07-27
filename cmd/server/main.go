package main

import (
	"github.com/aihmed/user-service/internal/handlers"
	"github.com/aihmed/user-service/internal/repository"
	"log"
	"net/http"
)

func main() {
	repo := repository.NewInMemoryUserRepo()
	userHandler := handlers.NewUserHandler(repo)

	http.HandleFunc("/register", userHandler.Register)
	// Add other routes

	log.Println("User Service started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
