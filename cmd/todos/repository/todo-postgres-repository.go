package repository

import (
	"database/sql"
	"fmt"

	"github.com/devlulcas/todoom/cmd/todos"
	"github.com/devlulcas/todoom/db"
)

const todoTableName = "todo"

type postgresTodoRepository struct {
	Conn *sql.DB
}

/**
 * Create a new Postgres repository (returns generic TodoRepository)
 */
func NewPostgresTodoRepository(db *sql.DB) todos.TodoRepository {
	return &postgresTodoRepository{db}
}

/**
 * Create a new todo on the database
 */
func (r *postgresTodoRepository) Create(todo todos.TodoInput) (id int64, err error) {
	sqlString := fmt.Sprintf("INSERT INTO %s (title, description, state) VALUES ($1, $2, $3) RETURNING id", todoTableName)

	// Execute the query and scan the returned id into the id
	err = r.Conn.QueryRow(sqlString, todo.Title, todo.Description, todo.State).Scan(&id)

	return id, err
}

/**
 * Find all todos on the database
 */
func (r *postgresTodoRepository) FindAll(pg db.PaginationInput) (list []todos.Todo, pagination db.Pagination, err error) {
	sqlString := fmt.Sprintf("SELECT * FROM %s LIMIT $1 OFFSET $2", todoTableName)

	rows, err := r.Conn.Query(sqlString, pagination.PerPage, pagination.Current)

	if err != nil {
		return nil, db.Pagination{}, err
	}

	// Close the statement after we finish
	defer rows.Close()

	for rows.Next() {
		var todo todos.Todo

		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.State, &todo.CreatedAt, &todo.UpdatedAt)

		if err != nil {
			return nil, db.Pagination{}, err
		}

		list = append(list, todo)
	}

	todosPagination := r.getPagination(pg.Page, pg.Limit)

	return list, todosPagination, nil
}

/**
 * Find one todo on the database
 */
func (r *postgresTodoRepository) FindOne(id int64) (todo todos.Todo, err error) {
	sqlString := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", todoTableName)

	// Execute the query
	err = r.Conn.QueryRow(sqlString, id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.State, &todo.CreatedAt, &todo.UpdatedAt)

	return todo, err
}

/**
 * Update a todo on the database
 */
func (r *postgresTodoRepository) Update(id int64, todo todos.TodoInput) (err error) {
	sqlString := fmt.Sprintf("UPDATE %s SET title = $1, description = $2, state = $3 WHERE id = $4", todoTableName)

	res, err := r.Conn.Exec(sqlString, todo.Title, todo.Description, todo.State, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No rows affected")
	}

	if rowsAffected != 1 {
		return fmt.Errorf("Something weird happened. Rows affected: %d", rowsAffected)
	}

	return nil
}

/**
 * Delete a todo on the database
 */
func (r *postgresTodoRepository) Delete(id int64) (err error) {
	sqlString := fmt.Sprintf("DELETE FROM %s WHERE id = $1", todoTableName)

	res, err := r.Conn.Exec(sqlString, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No rows affected")
	}

	if rowsAffected != 1 {
		return fmt.Errorf("Something weird happened. Rows affected: %d", rowsAffected)
	}

	return nil
}

/**
 * Get the pagination for the todos
 */
func (r *postgresTodoRepository) getPagination(current int64, perPage int64) (pagination db.Pagination) {
	total := 0

	sqlString := fmt.Sprintf("SELECT COUNT(*) FROM %s", todoTableName)

	err := r.Conn.QueryRow(sqlString).Scan(&total)

	if err != nil {
		return db.Pagination{}
	}

	nextPage := pagination.Current + 1

	if nextPage > int64(total) {
		nextPage = int64(total)
	}

	previousPage := pagination.Current - 1

	if previousPage < 0 {
		previousPage = 0
	}

	todosPagination := db.Pagination{
		PerPage:      perPage,
		Current:      current,
		PreviousPage: previousPage,
		NextPage:     nextPage,
		Total:        int64(total),
	}

	return todosPagination
}
