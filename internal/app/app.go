package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/AndrXxX/go-loyalty-service/internal/config"
	"github.com/AndrXxX/go-loyalty-service/internal/controllers"
	"github.com/AndrXxX/go-loyalty-service/internal/middlewares"
	"github.com/AndrXxX/go-loyalty-service/internal/services/converters"
	"github.com/AndrXxX/go-loyalty-service/internal/services/hashgenerator"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"github.com/AndrXxX/go-loyalty-service/internal/services/luhn"
	"github.com/AndrXxX/go-loyalty-service/internal/services/tokenservice"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

const shutdownTimeout = 5 * time.Second

type app struct {
	config  appConfig
	storage Storage
}

func New(c *config.Config, s Storage) *app {
	return &app{
		config:  appConfig{c},
		storage: s,
	}
}

func (a *app) Run(commonCtx context.Context) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	a.registerAPI(r)

	srv := &http.Server{Addr: a.config.c.RunAddress, Handler: r}

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logger.Log.Info("HTTP server ListenAndServe", zap.Error(err))
		}
	}()

	logger.Log.Info("listening", zap.String("host", a.config.c.RunAddress))

	<-commonCtx.Done()
	logger.Log.Info("shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}

	shutdown := make(chan struct{}, 1)
	go func() {
		if a.storage.DB != nil {
			db, _ := a.storage.DB.DB()
			_ = db.Close()
		}
		shutdown <- struct{}{}
	}()

	select {
	case <-shutdownCtx.Done():
		return fmt.Errorf("server shutdown: %w", shutdownCtx.Err())
	case <-shutdown:
		log.Println("finished")
	}

	return nil
}

func (a *app) registerAPI(r *chi.Mux) {
	hg := hashgenerator.Factory().SHA256(a.config.c.PasswordKey)
	ts := tokenservice.New(a.config.c.AuthKey, time.Duration(a.config.c.AuthKeyExpired)*time.Second)

	ac := controllers.NewAuthController(a.storage.US, hg, ts)
	r.Post("/api/user/register", ac.Register)
	r.Post("/api/user/login", ac.Login)

	r.Route("/api/user", func(r chi.Router) {

		r.Use(middlewares.IsAuthorized(ts).Handle)

		r.Route("/orders", func(r chi.Router) {
			oConverter := converters.NewOrderConverter()
			oc := controllers.NewOrdersController(luhn.Checker(), a.storage.US, a.storage.OS, oConverter)
			r.Post("/", oc.PostOrders)
			r.Get("/", oc.GetOrders)
		})

		bc := controllers.NewBalanceController()
		r.Route("/balance", func(r chi.Router) {
			r.Get("/", bc.Balance)
			r.Post("/withdraw", bc.Withdraw)
		})

		r.Route("/withdrawals", func(r chi.Router) {
			r.Get("/", bc.Withdrawals)
		})
	})

}
