package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// DbConfig
type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SSLMode  string
}

// NewPostgresDB
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

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
