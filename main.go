package main

import (
	"log"
	"net/http"
	"todo-app-go/db"
	"todo-app-go/handlers"
	"todo-app-go/repository"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database connection
	db.InitDB()
	defer db.GetDB().Close() // Ensure the DB connection is closed when the application shuts down

	// Create a new repository
	todoRepo := repository.NewTodoRepository()

	// Create a new handler
	todoHandler := handlers.NewTodoHandler(todoRepo)

	// Set up routing using gorilla/mux
	router := mux.NewRouter()

	// Define the routes and associate them with the handler methods
	router.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	router.HandleFunc("/todos", todoHandler.GetAllTodos).Methods("GET")
	router.HandleFunc("/todos/{id:[0-9]+}", todoHandler.GetTodoByID).Methods("GET")
	router.HandleFunc("/todos/{id:[0-9]+}", todoHandler.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id:[0-9]+}", todoHandler.DeleteTodo).Methods("DELETE")

	// Start the HTTP server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
