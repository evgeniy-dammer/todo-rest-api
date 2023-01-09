package service

import (
	todo "github.com/evgeniy-dammer/todo-rest-api"
	"github.com/evgeniy-dammer/todo-rest-api/pkg/repository"
)

// TodoListService
type TodoListService struct {
	repo repository.TodoList
}

// NewTodoListService is a constructor for TodoListService
func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

// Create creates new list in the system
func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

// GetAll returns all lists from the system
func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

// GetById returns list by id from the system
func (s *TodoListService) GetById(userId int, listId int) (todo.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

// DeleteById deletes list by id from the system
func (s *TodoListService) DeleteById(userId int, listId int) error {
	return s.repo.DeleteById(userId, listId)
}

// UpdateById updates list by id in the system
func (s *TodoListService) UpdateById(userId int, listId int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.UpdateById(userId, listId, input)
}
