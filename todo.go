package todo_rest_api

import "errors"

// TodoList model
type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

// UserLists model
type UserLists struct {
	Id     int
	UserId int
	ListId int
}

// TodoItem model
type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// ListItem model
type ListItem struct {
	Id     int
	ListId int
	ItemId int
}

// UpdateListInput
type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

// Validate checks if update input is nil
func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
