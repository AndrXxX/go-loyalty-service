package controllers

import (
	"github.com/AndrXxX/go-loyalty-service/internal/entities"
	"github.com/AndrXxX/go-loyalty-service/internal/interfaces"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
	"net/http"
)

type wConverter ormConverter[*ormmodels.Withdraw, *entities.Withdraw]

type balanceController struct {
	c  orderNumberChecker
	us interfaces.UserService
	os interfaces.OrderService
	ws interfaces.WithdrawService
	wc wConverter
}

func NewBalanceController(
	c orderNumberChecker,
	us interfaces.UserService,
	os interfaces.OrderService,
	ws interfaces.WithdrawService,
	wc wConverter,
) *balanceController {
	return &balanceController{c, us, os, ws, wc}
}

func (c *balanceController) Balance(_ http.ResponseWriter, _ *http.Request) {
	// TODO
}

func (c *balanceController) Withdraw(_ http.ResponseWriter, _ *http.Request) {
	// TODO
}

func (c *balanceController) Withdrawals(_ http.ResponseWriter, _ *http.Request) {
	// TODO
}
