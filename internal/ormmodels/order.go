package ormmodels

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Number    string `gorm:"uniqueIndex"`
	Status    string
	Accrual   *float64
	AuthorId  uint
	Author    User `gorm:"foreignKey:AuthorId"`
}
