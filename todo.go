package todo_rest_api

// TodoList model
type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
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
