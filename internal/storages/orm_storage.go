package storages

import (
	"context"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ormStorage[T any] struct {
	db *gorm.DB
}

func (s *ormStorage[T]) Find(m *T) *T {
	result := s.db.Where(m).First(m)
	if result.Error != nil {
		logger.Log.Info("failed to find Order", zap.Error(result.Error), zap.Any("order", m))
		return nil
	}
	return m
}

func (s *ormStorage[T]) Create(m *T) (*T, error) {
	result := s.db.Create(&m)
	if result.Error != nil {
		logger.Log.Info("failed to create model", zap.Error(result.Error), zap.Any("model", m))
		return nil, result.Error
	}
	return m, nil
}

func (s *ormStorage[T]) Update(m *T) error {
	result := s.db.Save(&m)
	if result.Error != nil {
		logger.Log.Info("failed to update model", zap.Error(result.Error), zap.Any("model", m))
		return result.Error
	}
	return nil
}

func (s *ormStorage[T]) FindAll(m *T) []*T {
	var list []*T
	result := s.db.Where(m).Order("created_at desc").Find(&list)
	if result.Error != nil {
		logger.Log.Info("failed to find all models", zap.Error(result.Error), zap.Any("model", m))
		return make([]*T, 0)
	}
	return list
}

func (s *ormStorage[T]) init(ctx context.Context, m *T) error {
	return s.db.WithContext(ctx).AutoMigrate(m)
}
