package controllers

import (
	"github.com/AndrXxX/go-loyalty-service/internal/entities"
	"github.com/AndrXxX/go-loyalty-service/internal/interfaces"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
)

type orderNumberChecker interface {
	Check(val string) bool
}

type ormConverter[source any, target any] interface {
	ConvertMany(list []source) []target
}

type balanceCounter interface {
	Count(u *ormmodels.User) *entities.Balance
}

type queueRunner interface {
	AddJob(interfaces.QueueJob) error
}

type jobsFactory interface {
	NewUpdateAccrualJob(o *ormmodels.Order) interfaces.QueueJob
}
