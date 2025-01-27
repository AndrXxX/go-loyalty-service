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
	bc balanceCounter
}

func NewBalanceController(
	c orderNumberChecker,
	us interfaces.UserService,
	os interfaces.OrderService,
	ws interfaces.WithdrawService,
	wc wConverter,
	bc balanceCounter,
) *balanceController {
	return &balanceController{c, us, os, ws, wc, bc}
}

func (c *balanceController) Balance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contenttypes.ApplicationJSON)
	userID := r.Context().Value(enums.UserID).(uint)
	user := c.us.Find(&ormmodels.User{ID: userID})
	if user == nil {
		logger.Log.Error("failed to find user", zap.Uint("userID", userID))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	encoded, err := json.Marshal(c.bc.Count(user))
	if err != nil {
		logger.Log.Error("failed to encode balance", zap.Error(err), zap.Uint("userID", userID))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(encoded)
	if err != nil {
		logger.Log.Error("failed to write balance response", zap.Error(err), zap.Uint("userID", userID))
	}
}

func (c *balanceController) Withdraw(w http.ResponseWriter, r *http.Request) {
	var m *entities.Withdraw
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&m)
	if err != nil {
		logger.Log.Error("failed to decode withdraw body", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !c.c.Check(m.Order) {
		logger.Log.Info("failed to check order number on withdraw", zap.Error(err), zap.String("orderNum", m.Order))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	userID := r.Context().Value(enums.UserID).(uint)
	user := c.us.Find(&ormmodels.User{ID: userID})
	if user == nil {
		logger.Log.Info("failed to find user", zap.Uint("userID", userID))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	balance := c.bc.Count(user)
	if *balance.Current < *m.Sum {
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = c.ws.Create(&ormmodels.Withdraw{AuthorID: userID, Order: m.Order, Sum: m.Sum})
	if err != nil {
		logger.Log.Error("failed to create withdraw model", zap.Uint("userID", userID), zap.Any("withdraw", m), zap.Error(err))
	}
}

func (c *balanceController) Withdrawals(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contenttypes.ApplicationJSON)
	userID := r.Context().Value(enums.UserID).(uint)
	withdrawals := c.ws.FindAll(&ormmodels.Withdraw{AuthorID: userID})
	list := c.wc.ConvertMany(withdrawals)
	if len(list) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	encoded, err := json.Marshal(list)
	if err != nil {
		logger.Log.Error("failed to encode withdraws list", zap.Error(err), zap.Uint("userID", userID))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(encoded)
	if err != nil {
		logger.Log.Error("failed to write withdraws list response", zap.Error(err), zap.Uint("userID", userID))
	}
}
