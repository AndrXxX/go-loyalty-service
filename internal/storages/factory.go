package storages

import (
	"context"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type storagesFactory struct {
	db *gorm.DB
}

func Factory(db *gorm.DB) *storagesFactory {
	return &storagesFactory{db}
}

func (f *storagesFactory) UserStorage(ctx context.Context) *userStorage {
	us := &userStorage{f.db}
	err := us.init(ctx)
	if err != nil {
		logger.Log.Error("failed to Init userStorage", zap.Error(err))
	}
	return us
}
