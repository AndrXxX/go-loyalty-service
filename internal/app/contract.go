package app

import (
	"context"
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
	OS interfaces.OrderService
	WS interfaces.WithdrawService
}

type queueRunner interface {
	Run() error
	Stop(context.Context) error
	AddJob(interfaces.QueueJob) error
}
