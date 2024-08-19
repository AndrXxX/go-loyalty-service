package storages

import (
	"context"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
	"gorm.io/gorm"
)

type orderStorage struct {
	db *gorm.DB
}

func (s *orderStorage) Find(m *ormmodels.Order) *ormmodels.Order {
	result := s.db.Model(m).First(m)
	if result.Error != nil {
		return nil
	}
	return m
}

func (s *orderStorage) Create(m *ormmodels.Order) (*ormmodels.Order, error) {
	result := s.db.Create(&m)
	if result.Error != nil {
		return nil, result.Error
	}
	return m, nil
}

func (s *orderStorage) FindAll(m *ormmodels.Order) []*ormmodels.Order {
	var list []*ormmodels.Order
	result := s.db.Where(m).Order("created_at desc").Find(&list)
	if result.Error != nil {
		return make([]*ormmodels.Order, 0)
	}
	return list
}

func (s *orderStorage) init(ctx context.Context) error {
	return s.db.WithContext(ctx).AutoMigrate(&ormmodels.Order{})
}
