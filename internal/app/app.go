package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/AndrXxX/go-loyalty-service/internal/config"
	"github.com/AndrXxX/go-loyalty-service/internal/controllers"
	"github.com/AndrXxX/go-loyalty-service/internal/middlewares"
	"github.com/AndrXxX/go-loyalty-service/internal/services/hashgenerator"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
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
	r.Post("/api/user/register", controllers.NewAuthController(a.storage.US, hg).Register)
	r.Post("/api/user/login", controllers.NewAuthController(a.storage.US, hg).Login)

	r.Route("/api/user", func(r chi.Router) {
		ts := tokenservice.New(a.config.c.AuthKey, time.Duration(a.config.c.AuthKeyExpired)*time.Second)
		r.Use(middlewares.IsAuthorized(ts).Handle)

		r.Route("/orders", func(r chi.Router) {
			r.Post("/", controllers.NewOrdersController().PostOrders)
			r.Get("/", controllers.NewOrdersController().GetOrders)
		})

		r.Route("/balance", func(r chi.Router) {
			r.Get("/", controllers.NewBalanceController().Balance)
			r.Post("/withdraw", controllers.NewBalanceController().Withdraw)
		})

		r.Route("/withdrawals", func(r chi.Router) {
			r.Get("/", controllers.NewBalanceController().Withdrawals)
		})
	})

}
