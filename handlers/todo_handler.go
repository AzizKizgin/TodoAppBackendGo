package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo-app-go/models"
	"todo-app-go/repository"
)

type TodoHandler struct {
	TodoRepository repository.TodoRepository
}

func NewTodoHandler(todoRepository repository.TodoRepository) *TodoHandler {
	return &TodoHandler{
		TodoRepository: todoRepository,
	}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var createTodoDTO models.CreateTodoDTO
	if err := json.NewDecoder(r.Body).Decode(&createTodoDTO); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	todo, err := h.TodoRepository.CreateTodo(createTodoDTO)
	if err != nil {
		http.Error(w, "Unable to create todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.TodoRepository.GetAllTodos()
	if err != nil {
		http.Error(w, "Unable to fetch todos", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Convert string ID to integer
	todoID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	todo, err := h.TodoRepository.GetTodoByID(todoID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to fetch todo: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Convert string ID to integer
	todoID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updateTodoDTO models.UpdateTodoDTO
	if err := json.NewDecoder(r.Body).Decode(&updateTodoDTO); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	todo, err := h.TodoRepository.UpdateTodo(todoID, updateTodoDTO)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to update todo: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Convert string ID to integer
	todoID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = h.TodoRepository.DeleteTodo(todoID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to delete todo: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
