package controllers

import "net/http"

type ordersController struct {
}

func NewOrdersController() *ordersController {
	return &ordersController{}
}

func (c *ordersController) PostOrders(_ http.ResponseWriter, _ *http.Request) {
	// TODO
}

func (c *ordersController) GetOrders(_ http.ResponseWriter, _ *http.Request) {
	// TODO
}
