package dbprovider

import (
	"fmt"
	"github.com/AndrXxX/go-loyalty-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbProvider struct {
	c *config.Config
}

func New(c *config.Config) *dbProvider {
	return &dbProvider{c}
}

func (p *dbProvider) DB() (*gorm.DB, error) {
	if p.c.DatabaseURI == "" {
		return nil, fmt.Errorf("empty DatabaseDSN")
	}
	db, err := gorm.Open(postgres.Open(p.c.DatabaseURI), &gorm.Config{})
	return db, err
}
