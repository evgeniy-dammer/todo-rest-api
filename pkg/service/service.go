package service

import (
	todo "github.com/evgeniy-dammer/todo-rest-api"
	"github.com/evgeniy-dammer/todo-rest-api/pkg/repository"
)

// Authorization interface
type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
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
	Create(userId int, listId int, item todo.TodoItem) (int, error)
	GetAll(userId int, listId int) ([]todo.TodoItem, error)
	GetById(userId int, itemId int) (todo.TodoItem, error)
	DeleteById(userId int, itemId int) error
	UpdateById(userId int, itemId int, input todo.UpdateItemInput) error
}

// Service
type Service struct {
	Authorization
	TodoList
	TodoItem
}

// NewService constructor for Service
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
