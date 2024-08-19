package controllers

import (
	"encoding/json"
	"github.com/AndrXxX/go-loyalty-service/internal/entities"
	"github.com/AndrXxX/go-loyalty-service/internal/enums"
	"github.com/AndrXxX/go-loyalty-service/internal/enums/contenttypes"
	"github.com/AndrXxX/go-loyalty-service/internal/interfaces"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"go.uber.org/zap"
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

func (c *balanceController) Withdrawals(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(enums.UserId).(uint)
	user := c.us.Find(&ormmodels.User{ID: userId})
	if user == nil {
		logger.Log.Error("failed to find user", zap.Uint("userId", userId))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	withdrawals := c.ws.FindAll(&ormmodels.Withdraw{Author: *user})
	list := c.wc.ConvertMany(withdrawals)
	encoded, err := json.Marshal(list)
	if err != nil {
		logger.Log.Error("failed to encode withdraws list", zap.Error(err), zap.Uint("userId", userId))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(encoded)
	if err != nil {
		logger.Log.Error("failed to write withdraws list response", zap.Error(err), zap.Uint("userId", userId))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", contenttypes.ApplicationJSON)
	w.WriteHeader(http.StatusOK)
}
