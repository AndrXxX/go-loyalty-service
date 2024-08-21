package storages

import (
	"context"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
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

func (f *storagesFactory) UserStorage(ctx context.Context) *ormStorage[ormmodels.User] {
	return getStorage[ormmodels.User](ctx, f.db, new(ormmodels.User))
}

func (f *storagesFactory) OrderStorage(ctx context.Context) *ormStorage[ormmodels.Order] {
	return getStorage[ormmodels.Order](ctx, f.db, new(ormmodels.Order))
}

func (f *storagesFactory) WithdrawStorage(ctx context.Context) *ormStorage[ormmodels.Withdraw] {
	return getStorage[ormmodels.Withdraw](ctx, f.db, new(ormmodels.Withdraw))
}

func getStorage[T interface{}](ctx context.Context, db *gorm.DB, m *T) *ormStorage[T] {
	s := &ormStorage[T]{db}
	err := s.init(ctx, m)
	if err != nil {
		logger.Log.Error("failed to init storage", zap.Error(err))
	}
	return s
}
