package app

import (
	"context"
	"gorm.io/gorm"
)

type migrator interface {
	Migrate(ctx context.Context, db *gorm.DB) error
}
