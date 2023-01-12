package repository

import (
	"fmt"
	todo "github.com/evgeniy-dammer/todo-rest-api"
	"github.com/jmoiron/sqlx"
	"strings"
)

// TodoListPostgres
type TodoListPostgres struct {
	db *sqlx.DB
}

// NewTodoListPostgres is a constructor for TodoListPostgres
func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

// Create inserts a new list into database
func (r *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}

	var id int

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", listTable)

	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err = row.Scan(&id); err != nil {
		if err = tx.Rollback(); err != nil {
			return 0, err
		}
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", userListTable)

	if _, err = tx.Exec(createUsersListQuery, userId, id); err != nil {
		if err = tx.Rollback(); err != nil {
			return 0, err
		}
		return 0, err
	}

	return id, tx.Commit()
}

// GetAll selects all lists from database
func (r *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN  %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1",
		listTable, userListTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

// GetById select list by id from database
func (r *TodoListPostgres) GetById(userId int, listId int) (todo.TodoList, error) {
	var list todo.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN  %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2",
		listTable, userListTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

// DeleteById deletes list by id from database
func (r *TodoListPostgres) DeleteById(userId int, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2",
		listTable, userListTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}

// UpdateById updates list by id in database
func (r *TodoListPostgres) UpdateById(userId int, listId int, input todo.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id = $%d AND ul.user_id = $%d",
		listTable, setQuery, userListTable, argId, argId+1)

	args = append(args, listId, userId)

	_, err := r.db.Exec(query, args...)

	return err
}
