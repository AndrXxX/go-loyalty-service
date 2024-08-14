package ormmodels

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Login    string `gorm:"uniqueIndex"`
	Password string
}
