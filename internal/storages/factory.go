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

func (f *storagesFactory) OrderStorage(ctx context.Context) *orderStorage {
	s := &orderStorage{f.db}
	err := s.init(ctx)
	if err != nil {
		logger.Log.Error("failed to Init orderStorage", zap.Error(err))
	}
	return s
}

func (f *storagesFactory) WithdrawStorage(ctx context.Context) *withdrawStorage {
	s := &withdrawStorage{f.db}
	err := s.init(ctx)
	if err != nil {
		logger.Log.Error("failed to Init withdrawStorage", zap.Error(err))
	}
	return s
}
