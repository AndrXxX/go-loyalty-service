package main

import (
	"context"
	"fmt"
	"github.com/AndrXxX/go-loyalty-service/internal/app"
	"github.com/AndrXxX/go-loyalty-service/internal/config"
	"github.com/AndrXxX/go-loyalty-service/internal/services/dbprovider"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"github.com/AndrXxX/go-loyalty-service/internal/services/queue"
	"github.com/AndrXxX/go-loyalty-service/internal/storages"
	"github.com/asaskevich/govalidator"
	"go.uber.org/zap"
	"log"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	settings, err := initSetting()
	if err != nil {
		log.Fatal(err)
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	s, err := initStorage(ctx, settings)
	if err != nil {
		logger.Log.Error("failed to connect to database", zap.Error(err))
		return
	}

	qr := queue.NewRunner(time.Second).SetWorkersCount(5)
	if err := app.New(settings, *s, qr).Run(ctx); err != nil {
		logger.Log.Fatal(err.Error())
	}
}

func initSetting() (*config.Config, error) {
	settings := config.NewConfig()
	if err := logger.Initialize(settings.LogLevel); err != nil {
		return nil, err
	}
	parseFlags(settings)
	parseEnv(settings)
	if _, err := govalidator.ValidateStruct(settings); err != nil {
		return nil, err
	}
	return settings, nil
}

func initStorage(ctx context.Context, c *config.Config) (*app.Storage, error) {
	db, err := dbprovider.New(c).DB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	sf := storages.Factory(db)
	return &app.Storage{
		DB: db,
		US: sf.UserStorage(ctx),
		OS: sf.OrderStorage(ctx),
		WS: sf.WithdrawStorage(ctx),
	}, nil
}
