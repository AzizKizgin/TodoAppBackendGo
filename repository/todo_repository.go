package repository

import (
	"database/sql"
	"fmt"
	"todo-app-go/db"
	"todo-app-go/models"
)

type TodoRepository interface {
	CreateTodo(todo models.CreateTodoDTO) (models.TodoResponse, error)
	GetAllTodos() (models.TodoListResponse, error)
	GetTodoByID(id int) (models.TodoResponse, error)
	UpdateTodo(id int,todo models.UpdateTodoDTO) (models.TodoResponse, error)
	DeleteTodo(id int) error
}

type todoRepository struct {
	DB *sql.DB
}

func NewTodoRepository() TodoRepository {
	return &todoRepository{DB: db.GetDB()}
}


func (t *todoRepository) CreateTodo(createTodoDTO models.CreateTodoDTO) (models.TodoResponse, error) {
	query := `INSERT INTO todos (title, desc, created_at, due_date, is_completed)
		VALUES ($1, $2, $3, $4, $5) RETURNING id`

	var todo models.TodoResponse
	err := t.DB.QueryRow(query, createTodoDTO.Title, createTodoDTO.Desc, createTodoDTO.CreatedAt, createTodoDTO.DueDate, createTodoDTO.IsCompleted).Scan(&todo.ID)
	if err != nil {
		return models.TodoResponse{}, fmt.Errorf("unable to insert todo into the database: %v", err)
	}

	return todo, nil
}


func (t *todoRepository) GetAllTodos() (models.TodoListResponse, error) {
	query := `SELECT id, title, desc, created_at, due_date, is_completed FROM todos`
	rows, err := t.DB.Query(query)
	if err != nil {
		return models.TodoListResponse{}, fmt.Errorf("unable to fetch todos: %v", err)
	}
	defer rows.Close()
	var todos []models.TodoResponse
	for rows.Next() {
		var todo models.TodoResponse
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Desc, &todo.CreatedAt, &todo.DueDate, &todo.IsCompleted); err != nil {
			return models.TodoListResponse{}, fmt.Errorf("unable to scan todo: %v", err)
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		return models.TodoListResponse{}, fmt.Errorf("error iterating over todos: %v", err)
	}
	return models.TodoListResponse{Todos: todos}, nil
}

func (t *todoRepository) GetTodoByID(id int) (models.TodoResponse, error) {
	query := `SELECT id, title, desc, created_at, due_date, is_completed FROM todos WHERE id = $1`
	row := t.DB.QueryRow(query, id)

	var todo models.TodoResponse
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Desc, &todo.CreatedAt, &todo.DueDate, &todo.IsCompleted); err != nil {
		if err == sql.ErrNoRows {
			return models.TodoResponse{}, fmt.Errorf("todo with id %d not found", id)
		}
		return models.TodoResponse{}, fmt.Errorf("unable to scan todo: %v", err)
	}

	return todo, nil
}

func (t *todoRepository) UpdateTodo(id int, updateTodoDTO models.UpdateTodoDTO) (models.TodoResponse, error) {
	query := `UPDATE todos 
              SET title = $1, desc = $2, due_date = $3, is_completed = $4 
              WHERE id = $5 
              RETURNING id, title, desc, created_at, due_date, is_completed`

	var updatedTodo models.TodoResponse
	err := t.DB.QueryRow(query, updateTodoDTO.Title, updateTodoDTO.Desc, updateTodoDTO.DueDate, updateTodoDTO.IsCompleted, id).
		Scan(&updatedTodo.ID, &updatedTodo.Title, &updatedTodo.Desc, &updatedTodo.CreatedAt, &updatedTodo.DueDate, &updatedTodo.IsCompleted)

	if err != nil {
		return models.TodoResponse{}, fmt.Errorf("unable to update todo with id %d: %v", id, err)
	}

	return updatedTodo, nil
}



func (t *todoRepository) DeleteTodo(id int) error {
	query := `DELETE FROM todos WHERE id = $1`
	_, err := t.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("unable to delete todo with id %d: %v", id, err)
	}
	return nil
}
