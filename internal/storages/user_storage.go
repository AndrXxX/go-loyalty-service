package storages

import (
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
	"gorm.io/gorm"
)

type userStorage struct {
	db *gorm.DB
}

func User(db *gorm.DB) *userStorage {
	return &userStorage{db}
}

func (s *userStorage) Find(login string) *ormmodels.User {
	var result *ormmodels.User
	s.db.Model(ormmodels.User{Login: login}).First(result)
	return result
}

func (s *userStorage) Create(u *ormmodels.User) (*ormmodels.User, error) {
	result := s.db.Create(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}
