package ormmodels

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Number  string `gorm:"uniqueIndex"`
	Status  string
	Accrual *float64
	Author  User `gorm:"embedded"`
}
