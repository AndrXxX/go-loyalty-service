package controllers

import (
	"github.com/AndrXxX/go-loyalty-service/internal/enums"
	"github.com/AndrXxX/go-loyalty-service/internal/enums/orderstatuses"
	"github.com/AndrXxX/go-loyalty-service/internal/interfaces"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type ordersController struct {
	c  orderNumberChecker
	us interfaces.UserService
	os interfaces.OrderService
}

func NewOrdersController(c orderNumberChecker, us interfaces.UserService, os interfaces.OrderService) *ordersController {
	return &ordersController{c, us, os}
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
	userId := r.Context().Value(enums.UserId).(uint)
	user := c.us.Find(&ormmodels.User{ID: userId})
	if user == nil {
		logger.Log.Error("failed to find user", zap.Error(err))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	existOrder := c.os.Find(&ormmodels.Order{Number: orderNum})
	if existOrder == nil {
		_, err := c.os.Create(&ormmodels.Order{Number: orderNum, Author: *user, Status: orderstatuses.Waiting})
		if err != nil {
			logger.Log.Error("failed to create order", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		return
	}
	if existOrder.Author.ID != userId {
		w.WriteHeader(http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *ordersController) GetOrders(_ http.ResponseWriter, _ *http.Request) {
	// TODO
}
