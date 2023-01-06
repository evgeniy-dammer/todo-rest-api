package service

import "github.com/evgeniy-dammer/todo-rest-api/pkg/repository"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

// NewService constructor for Service
func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
