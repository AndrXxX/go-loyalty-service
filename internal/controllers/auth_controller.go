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
	u := c.fetchUser(r)
	if u == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	exist := c.us.Find(&ormmodels.User{Login: u.Login})
	if exist != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	created, err := c.us.Create(&ormmodels.User{Login: u.Login, Password: c.hg.Generate([]byte(u.Password))})
	if err != nil {
		logger.Log.Error("failed to create user on register request", zap.Error(err))
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err := c.setAuthToken(w, created); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *authController) Login(w http.ResponseWriter, r *http.Request) {
	u := c.fetchUser(r)
	if u == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	exist := c.us.Find(&ormmodels.User{Login: u.Login})
	if exist == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if exist.Password != c.hg.Generate([]byte(u.Password)) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err := c.setAuthToken(w, exist); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *authController) fetchUser(r *http.Request) *entities.User {
	var u *entities.User
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&u)
	if err != nil {
		logger.Log.Error("failed to decode request", zap.Error(err))
		return nil
	}
	if _, err := govalidator.ValidateStruct(u); err != nil {
		logger.Log.Error("failed to validate request", zap.Error(err))
		return nil
	}
	return u
}

func (c *authController) setAuthToken(w http.ResponseWriter, user *ormmodels.User) error {
	token, err := c.ts.Encrypt(user.ID)
	if err != nil {
		logger.Log.Error("failed to encrypt token on request", zap.Error(err))
		return err
	}
	http.SetCookie(w, &http.Cookie{Name: enums.AuthToken, Value: token})
	return err
}
