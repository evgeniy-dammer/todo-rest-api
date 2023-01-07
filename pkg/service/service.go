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
}

// TodoItem interface
type TodoItem interface {
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
	}
}
