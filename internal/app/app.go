package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/AndrXxX/go-loyalty-service/internal/config"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

const shutdownTimeout = 5 * time.Second

type app struct {
	config struct {
		c *config.Config
	}
	storage struct {
		db *gorm.DB
	}
}

func New(c *config.Config, db *gorm.DB) *app {
	return &app{
		config: struct {
			c *config.Config
		}{c},
		storage: struct {
			db *gorm.DB
		}{db},
	}
}

func (a *app) Run(commonCtx context.Context) error {

	r := chi.NewRouter()

	// TODO: realise routes

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
		if a.storage.db != nil {
			db, _ := a.storage.db.DB()
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
