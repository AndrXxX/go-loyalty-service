package controllers

import "net/http"

type balanceController struct {
}

func NewBalanceController() *balanceController {
	return &balanceController{}
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
