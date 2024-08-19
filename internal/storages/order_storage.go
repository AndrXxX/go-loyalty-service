package storages

import (
	"context"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type orderStorage struct {
	db *gorm.DB
}

func (s *orderStorage) Find(m *ormmodels.Order) *ormmodels.Order {
	var found *ormmodels.Order
	result := s.db.Where(m).First(found)
	if result.Error != nil {
		logger.Log.Info("failed to find Order", zap.Error(result.Error), zap.Any("order", m))
		return nil
	}
	return found
}

func (s *orderStorage) Create(m *ormmodels.Order) (*ormmodels.Order, error) {
	result := s.db.Create(&m)
	if result.Error != nil {
		logger.Log.Info("failed to create Order", zap.Error(result.Error), zap.Any("order", m))
		return nil, result.Error
	}
	return m, nil
}

func (s *orderStorage) FindAll(m *ormmodels.Order) []*ormmodels.Order {
	var list []*ormmodels.Order
	result := s.db.Where(m).Order("created_at desc").Find(&list)
	if result.Error != nil {
		logger.Log.Info("failed to find all Orders", zap.Error(result.Error), zap.Any("order", m))
		return make([]*ormmodels.Order, 0)
	}
	return list
}

func (s *orderStorage) init(ctx context.Context) error {
	return s.db.WithContext(ctx).AutoMigrate(&ormmodels.Order{})
}
