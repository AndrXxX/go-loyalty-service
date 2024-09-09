package ormmodels

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Login     string `gorm:"uniqueIndex"`
	Password  string
}
