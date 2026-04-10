package main

import (
	"log"
	"net/http"
	"os"

	"taskflow/internal/db"
	"taskflow/internal/handlers"
)

func main() {
	db.Connect(os.Getenv("DB_URL"))

	http.HandleFunc("/auth/register", handlers.Register)
	http.HandleFunc("/auth/login", handlers.Login)

	http.HandleFunc("/projects", handlers.Projects)
	http.HandleFunc("/tasks", handlers.Tasks)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
