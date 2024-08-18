package app

import (
	"github.com/AndrXxX/go-loyalty-service/internal/config"
	"github.com/AndrXxX/go-loyalty-service/internal/interfaces"
	"gorm.io/gorm"
)

type appConfig struct {
	c *config.Config
}

type Storage struct {
	DB *gorm.DB
	US interfaces.UserService
}
