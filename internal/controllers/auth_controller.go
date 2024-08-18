package controllers

import (
	"encoding/json"
	"github.com/AndrXxX/go-loyalty-service/internal/entities"
	"github.com/AndrXxX/go-loyalty-service/internal/interfaces"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"go.uber.org/zap"
	"net/http"
)

type authController struct {
	us interfaces.UserService
}

func NewAuthController(us interfaces.UserService) *authController {
	return &authController{us}
}

func (c *authController) Register(w http.ResponseWriter, r *http.Request) {
	var u *entities.User
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&u)
	if err != nil {
		logger.Log.Error("failed to decode register request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	exist := c.us.Find(u.Login)
	if exist != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	// TODO: hash password
	orm := ormmodels.User{Login: u.Login, Password: u.Password}
	_, err = c.us.Create(&orm)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	// TODO: make auth
}

func (c *authController) Login(_ http.ResponseWriter, _ *http.Request) {
	// TODO
}
