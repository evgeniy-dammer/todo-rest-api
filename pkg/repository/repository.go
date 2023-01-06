package repository

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// NewRepository constructor for Repository
func NewRepository() *Repository {
	return &Repository{}
}
