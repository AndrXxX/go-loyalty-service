package storages

import (
	"context"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
	"gorm.io/gorm"
)

type userStorage struct {
	db *gorm.DB
}

func (s *userStorage) Find(u *ormmodels.User) *ormmodels.User {
	result := s.db.Model(u).First(u)
	if result.Error != nil {
		return nil
	}
	return u
}

func (s *userStorage) Create(u *ormmodels.User) (*ormmodels.User, error) {
	result := s.db.Create(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

func (s *userStorage) init(ctx context.Context) error {
	return s.db.WithContext(ctx).AutoMigrate(&ormmodels.User{})
}
