package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// database table names
const (
	userTable     = "users"
	listTable     = "lists"
	itemTable     = "items"
	userListTable = "users_lists"
	listItemTable = "lists_items"
)

// DbConfig is a database config
type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SSLMode  string
}

// NewPostgresDB create connection to database
func NewPostgresDB(cfg DbConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DbName, cfg.Password, cfg.SSLMode,
		),
	)

	if err != nil {
		return nil, err
	}

	// verify connection
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
