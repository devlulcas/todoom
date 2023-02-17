package usecase

import (
	"fmt"

	"github.com/devlulcas/todoom/cmd/todos"
	"github.com/devlulcas/todoom/db"
)

type todoUseCase struct {
	todoRepository todos.TodoRepository
}

func NewTodoUseCase(todoRepository todos.TodoRepository) todos.TodoUseCase {
	return &todoUseCase{todoRepository}
}

func (u *todoUseCase) Create(todo todos.TodoInput) (id int64, err error) {
	if todo.Title == "" {
		return 0, fmt.Errorf(todos.ErrEmptyTitle)
	}

	todoState := todos.State(todo.State)

	if todoState != todos.Active && todoState != todos.Completed {
		todoState = todos.Active
	}

	todo.State = todoState

	return u.todoRepository.Create(todo)
}

func (u *todoUseCase) Delete(id int64) (err error) {
	if id <= 0 {
		return fmt.Errorf(todos.ErrInvalidId)
	}

	return u.todoRepository.Delete(id)
}

func (u *todoUseCase) FindAll(pg db.PaginationInput) (todos []todos.Todo, pagination db.Pagination, err error) {
	return u.todoRepository.FindAll(pg)
}

func (u *todoUseCase) FindOne(id int64) (todo todos.Todo, err error) {
	if id <= 0 {
		return todos.Todo{}, fmt.Errorf(todos.ErrInvalidId)
	}

	return u.todoRepository.FindOne(id)
}

func (u *todoUseCase) Update(id int64, todo todos.TodoInput) (err error) {
	if id <= 0 {
		return fmt.Errorf(todos.ErrInvalidId)
	}

	if todo.Title == "" {
		return fmt.Errorf(todos.ErrEmptyTitle)
	}

	if todo.State != todos.Active && todo.State != todos.Completed {
		return fmt.Errorf(todos.ErrInvalidTodoState)
	}

	return u.todoRepository.Update(id, todo)
}
