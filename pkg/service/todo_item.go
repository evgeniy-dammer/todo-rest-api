package service

import (
	todo "github.com/evgeniy-dammer/todo-rest-api"
	"github.com/evgeniy-dammer/todo-rest-api/pkg/repository"
)

// TodoItemService
type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

// NewTodoItemService is a constructor for TodoItemService
func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

// Create creates new item in list in the system
func (s *TodoItemService) Create(userId int, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)

	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, item)
}

// GetAll returns all items of list from the system
func (s *TodoItemService) GetAll(userId int, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

// GetById returns item by id from the system
func (s *TodoItemService) GetById(userId int, itemId int) (todo.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

// DeleteById deletes item by id from the system
func (s *TodoItemService) DeleteById(userId int, itemId int) error {
	return s.repo.DeleteById(userId, itemId)
}

// UpdateById updates item by id in the system
func (s *TodoItemService) UpdateById(userId int, itemId int, input todo.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.UpdateById(userId, itemId, input)
}
