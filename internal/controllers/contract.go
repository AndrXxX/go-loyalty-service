package controllers

import (
	"github.com/AndrXxX/go-loyalty-service/internal/entities"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
)

type orderNumberChecker interface {
	Check(val string) bool
}

type orderConverter interface {
	ConvertMany(list []*ormmodels.Order) []*entities.Order
}

type withdrawConverter interface {
	ConvertMany(list []*ormmodels.Withdraw) []*entities.Withdraw
}
