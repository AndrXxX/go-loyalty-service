package controllers

import (
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type ordersController struct {
	c orderNumberChecker
}

func NewOrdersController(c orderNumberChecker) *ordersController {
	return &ordersController{c}
}

func (c *ordersController) PostOrders(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log.Error("failed to read body on post orders request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	orderNum := string(bytes)
	if !c.c.Check(orderNum) {
		logger.Log.Error("failed to check order number", zap.Error(err), zap.String("orderNum", orderNum))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	// TODO: put order to DB
	w.WriteHeader(http.StatusOK)
}

func (c *ordersController) GetOrders(_ http.ResponseWriter, _ *http.Request) {
	// TODO
}
