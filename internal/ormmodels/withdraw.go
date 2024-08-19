package ormmodels

import (
	"gorm.io/gorm"
	"time"
)

type Withdraw struct {
	gorm.Model
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Order     string `gorm:"uniqueIndex"`
	Sum       *float64
	AuthorID  uint
	Author    User `gorm:"foreignKey:AuthorID"`
}
