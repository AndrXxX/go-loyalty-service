package app

import (
	"context"
	"database/sql"
	"github.com/AndrXxX/go-loyalty-service/internal/config"
)

type app struct {
	config struct {
		c *config.Config
	}
	storage struct {
		db *sql.DB
	}
}

func New(c *config.Config, db *sql.DB) *app {
	return &app{
		config: struct {
			c *config.Config
		}{c},
		storage: struct {
			db *sql.DB
		}{db},
	}
}

func (a *app) Run(_ context.Context) error {
	// TODO: realise app
	return nil
}
