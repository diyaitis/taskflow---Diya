package main

import (
	"log"
	"net/http"
	"os"

	"taskflow/internal/db"
	"taskflow/internal/handlers"
	"taskflow/internal/middleware"
)

func main() {

	// connect DB
	db.Connect(os.Getenv("DB_URL"))

	// public routes
	http.HandleFunc("/auth/register", handlers.Register)
	http.HandleFunc("/auth/login", handlers.Login)

	// protected routes
	http.HandleFunc("/projects", middleware.AuthMiddleware(handlers.Projects))
	http.HandleFunc("/tasks", middleware.AuthMiddleware(handlers.Tasks))

	log.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}