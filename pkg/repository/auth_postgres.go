package repository

import (
	"fmt"
	todo "github.com/evgeniy-dammer/todo-rest-api"
	"github.com/jmoiron/sqlx"
)

// AuthPostgres
type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres constructor for AuthPostgres
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// CreateUser insert User into database
func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// GetUser returns users id by username and password hash
func (r *AuthPostgres) GetUser(username string, password string) (todo.User, error) {
	var user todo.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password_hash = $2", userTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
