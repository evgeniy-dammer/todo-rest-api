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
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId int, listId int) (todo.TodoList, error)
	DeleteById(userId int, listId int) error
	UpdateById(userId int, listId int, input todo.UpdateListInput) error
}

// TodoItem interface
type TodoItem interface {
	Create(listId int, item todo.TodoItem) (int, error)
	GetAll(userId int, listId int) ([]todo.TodoItem, error)
	GetById(userId int, itemId int) (todo.TodoItem, error)
	DeleteById(userId int, itemId int) error
	UpdateById(userId int, itemId int, input todo.UpdateItemInput) error
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
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
