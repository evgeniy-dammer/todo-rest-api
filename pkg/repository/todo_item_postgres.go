package repository

import (
	"fmt"
	todo "github.com/evgeniy-dammer/todo-rest-api"
	"github.com/jmoiron/sqlx"
	"strings"
)

// TodoItemPostgres
type TodoItemPostgres struct {
	db *sqlx.DB
}

// NewTodoItemPostgres is a constructor for TodoItemPostgres
func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

// Create inserts a new item by list into database
func (r *TodoItemPostgres) Create(listId int, item todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", itemTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)

	err = row.Scan(&itemId)

	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listItemTable)

	_, err = tx.Exec(createListItemsQuery, listId, itemId)

	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	return itemId, tx.Commit()
}

// GetAll selects all items by list from database
func (r *TodoItemPostgres) GetAll(userId int, listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description, tl.done FROM %s tl "+
		"INNER JOIN %s li ON li.item_id = tl.id "+
		"INNER JOIN %s ul ON ul.list_id = li.list_id "+
		"WHERE li.list_id = $1 AND ul.user_id = $2",
		itemTable, listItemTable, userListTable)

	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

// GetById select item by id from database
func (r *TodoItemPostgres) GetById(userId int, itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description, tl.done FROM %s tl "+
		"INNER JOIN %s li ON li.item_id = tl.id "+
		"INNER JOIN %s ul ON ul.list_id = li.list_id "+
		"WHERE ti.id = $1 AND ul.user_id = $2",
		itemTable, listItemTable, userListTable)

	err := r.db.Get(&item, query, itemId, userId)

	return item, err
}

// DeleteById deletes item by id from database
func (r *TodoItemPostgres) DeleteById(userId int, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s li, %s ul WHERE tl.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND tl.id = $2",
		itemTable, listItemTable, userListTable)

	_, err := r.db.Exec(query, userId, itemId)

	return err
}

// UpdateById updates item by id in database
func (r *TodoItemPostgres) UpdateById(userId int, itemId int, input todo.UpdateItemInput) error {
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

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s li, %s ul WHERE tl.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND tl.id = $%d",
		itemTable, setQuery, listItemTable, userListTable, argId, argId+1)

	args = append(args, userId, itemId)

	_, err := r.db.Exec(query, args...)

	return err
}
