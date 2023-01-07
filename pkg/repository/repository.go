package repository

import (
	todo "github.com/evgeniy-dammer/todo-rest-api"
	"github.com/jmoiron/sqlx"
)

// Authorization interface
type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username string, password string) (todo.User, error)
}

// TodoList interface
type TodoList interface {
}

// TodoItem interface
type TodoItem interface {
}

// Repository
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// NewRepository constructor for Repository
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
