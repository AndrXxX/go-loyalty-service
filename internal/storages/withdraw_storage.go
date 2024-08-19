package storages

import (
	"context"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type withdrawStorage struct {
	db *gorm.DB
}

func (s *withdrawStorage) Find(m *ormmodels.Withdraw) *ormmodels.Withdraw {
	var found *ormmodels.Withdraw
	result := s.db.Where(m).First(found)
	if result.Error != nil {
		logger.Log.Info("failed to find Withdraw", zap.Error(result.Error), zap.Any("withdraw", m))
		return nil
	}
	return found
}

func (s *withdrawStorage) Create(m *ormmodels.Withdraw) (*ormmodels.Withdraw, error) {
	result := s.db.Create(&m)
	if result.Error != nil {
		logger.Log.Info("failed to create Withdraw", zap.Error(result.Error), zap.Any("withdraw", m))
		return nil, result.Error
	}
	return m, nil
}

func (s *withdrawStorage) FindAll(m *ormmodels.Withdraw) []*ormmodels.Withdraw {
	var list []*ormmodels.Withdraw
	result := s.db.Where(m).Order("created_at desc").Find(&list)
	if result.Error != nil {
		logger.Log.Info("failed to find all Withdraws", zap.Error(result.Error), zap.Any("withdraw", m))
		return make([]*ormmodels.Withdraw, 0)
	}
	return list
}

func (s *withdrawStorage) init(ctx context.Context) error {
	return s.db.WithContext(ctx).AutoMigrate(&ormmodels.Withdraw{})
}
