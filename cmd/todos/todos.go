package todos

import (
	"github.com/devlulcas/todoom/db"
)

type State string

const (
	Active    State = "active"
	Completed State = "completed"
)

type Todo struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Title       string `json:"title"`
	State       State  `json:"state"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type TodoInput struct {
	Description string `json:"description"`
	Title       string `json:"title"`
	State       State  `json:"state"`
}

type TodoRepository interface {
	Create(todo TodoInput) (id int64, err error)
	Delete(id int64) (err error)
	FindAll(pg db.PaginationInput) (todos []Todo, pagination db.Pagination, err error)
	FindOne(id int64) (todo Todo, err error)
	Update(id int64, todo TodoInput) (err error)
}

type TodoUseCase interface {
	Create(todo TodoInput) (id int64, err error)
	Delete(id int64) (err error)
	FindAll(pg db.PaginationInput) (todos []Todo, pagination db.Pagination, err error)
	FindOne(id int64) (todo Todo, err error)
	Update(id int64, todo TodoInput) (err error)
}

const (
	ErrEmptyTitle       = "title is empty"
	ErrInvalidTodoState = "invalid todo state"
	ErrInvalidId        = "invalid id"
	ErrTodoNotFound     = "todo not found"
)
