package controllers

import (
	"encoding/json"
	"github.com/AndrXxX/go-loyalty-service/internal/entities"
	"github.com/AndrXxX/go-loyalty-service/internal/enums"
	"github.com/AndrXxX/go-loyalty-service/internal/interfaces"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"github.com/asaskevich/govalidator"
	"go.uber.org/zap"
	"net/http"
)

type authController struct {
	us interfaces.UserService
	hg interfaces.HashGenerator
	ts interfaces.TokenService
}

func NewAuthController(us interfaces.UserService, hg interfaces.HashGenerator, ts interfaces.TokenService) *authController {
	return &authController{us, hg, ts}
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
	if _, err := govalidator.ValidateStruct(u); err != nil {
		logger.Log.Error("failed to validate on register request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	exist := c.us.Find(u.Login)
	if exist != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	orm := ormmodels.User{Login: u.Login, Password: c.hg.Generate([]byte(u.Password))}
	created, err := c.us.Create(&orm)
	if err != nil {
		logger.Log.Error("failed to create user on register request", zap.Error(err))
		w.WriteHeader(http.StatusConflict)
		return
	}
	token, err := c.ts.Encrypt(created.ID)
	if err != nil {
		logger.Log.Error("failed to encrypt token on register request", zap.Error(err))
		return
	}
	r.AddCookie(&http.Cookie{Name: enums.AuthToken, Value: token})
	w.WriteHeader(http.StatusOK)
}

func (c *authController) Login(_ http.ResponseWriter, _ *http.Request) {
	// TODO
}
