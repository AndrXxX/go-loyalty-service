package controllers

import (
	"encoding/json"
	"github.com/AndrXxX/go-loyalty-service/internal/entities"
	"github.com/AndrXxX/go-loyalty-service/internal/enums"
	"github.com/AndrXxX/go-loyalty-service/internal/enums/contenttypes"
	"github.com/AndrXxX/go-loyalty-service/internal/enums/orderstatuses"
	"github.com/AndrXxX/go-loyalty-service/internal/interfaces"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type oConverter ormConverter[*ormmodels.Order, *entities.Order]

type ordersController struct {
	c  orderNumberChecker
	us interfaces.UserService
	os interfaces.OrderService
	oc oConverter
	qr queueRunner
	jf jobsFactory
}

func NewOrdersController(
	c orderNumberChecker,
	us interfaces.UserService,
	os interfaces.OrderService,
	oc oConverter,
	qr queueRunner,
	jf jobsFactory,
) *ordersController {
	return &ordersController{c, us, os, oc, qr, jf}
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
	userID := r.Context().Value(enums.UserID).(uint)
	user := c.us.Find(&ormmodels.User{ID: userID})
	if user == nil {
		logger.Log.Error("failed to find user", zap.Uint("userID", userID))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	existOrder := c.os.Find(&ormmodels.Order{Number: orderNum})
	if existOrder == nil {
		order, err := c.os.Create(&ormmodels.Order{Number: orderNum, AuthorID: user.ID, Status: orderstatuses.Waiting})
		if err != nil {
			logger.Log.Error("failed to create order", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err = c.qr.AddJob(c.jf.NewUpdateAccrualJob(order)); err != nil {
			logger.Log.Error("failed to add UpdateAccrualJob", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		return
	}
	if existOrder.AuthorID != userID {
		w.WriteHeader(http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *ordersController) GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contenttypes.ApplicationJSON)
	userID := r.Context().Value(enums.UserID).(uint)
	user := c.us.Find(&ormmodels.User{ID: userID})
	if user == nil {
		logger.Log.Error("failed to find user", zap.Uint("userID", userID))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	orders := c.os.FindAll(&ormmodels.Order{Author: *user})
	list := c.oc.ConvertMany(orders)
	encoded, err := json.Marshal(list)
	if err != nil {
		logger.Log.Error("failed to encode orders list", zap.Error(err), zap.Uint("userID", userID))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(encoded)
	if err != nil {
		logger.Log.Error("failed to write orders list response", zap.Error(err), zap.Uint("userID", userID))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
