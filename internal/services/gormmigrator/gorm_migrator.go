package gormmigrator

import (
	"context"
	models "github.com/AndrXxX/go-loyalty-service/internal/entities"
	"gorm.io/gorm"
)

type migrator struct {
	db *gorm.DB
}

func New() *migrator {
	return &migrator{}
}

func (p *migrator) Migrate(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&models.User{})
}
