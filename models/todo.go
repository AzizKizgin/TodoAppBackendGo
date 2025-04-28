package models

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID          int           `json:"id"`
	Title       string        `json:"title"`
	Desc        sql.NullString `json:"desc"`
	CreatedAt   time.Time     `json:"created_at"`
	DueDate     sql.NullTime  `json:"due_date"`
	IsCompleted bool          `json:"is_completed"`
}

type CreateTodoDTO struct {
	Title      string       `json:"title"`
	Desc       sql.NullString `json:"desc"`
	CreatedAt  time.Time   `json:"created_at"` 
	DueDate    sql.NullTime `json:"due_date"`
	IsCompleted bool       `json:"is_completed"`
}

type UpdateTodoDTO struct {
	Title      string        `json:"title"`
	Desc       sql.NullString `json:"desc"`
	DueDate    sql.NullTime  `json:"due_date"`
	IsCompleted bool         `json:"is_completed"`
}

type TodoResponse struct {
	ID          int           `json:"id"`
	Title       string        `json:"title"`
	Desc        sql.NullString `json:"desc"`
	CreatedAt   time.Time     `json:"created_at"`
	DueDate     sql.NullTime  `json:"due_date"`
	IsCompleted bool          `json:"is_completed"`
}

type TodoListResponse struct {
	Todos []TodoResponse `json:"todos"`
}
