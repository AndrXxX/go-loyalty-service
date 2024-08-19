package main

import (
	"context"
	"github.com/AndrXxX/go-loyalty-service/internal/app"
	"github.com/AndrXxX/go-loyalty-service/internal/config"
	"github.com/AndrXxX/go-loyalty-service/internal/services/dbprovider"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"github.com/AndrXxX/go-loyalty-service/internal/storages"
	"github.com/asaskevich/govalidator"
	"go.uber.org/zap"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	settings := config.NewConfig()
	if err := logger.Initialize(settings.LogLevel); err != nil {
		log.Fatal(err)
	}
	parseFlags(settings)
	parseEnv(settings)
	if _, err := govalidator.ValidateStruct(settings); err != nil {
		logger.Log.Fatal(err.Error())
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	db, err := dbprovider.New(settings).DB()
	if err != nil {
		logger.Log.Error("failed to connect to database", zap.Error(err))
	}
	sf := storages.Factory(db)
	s := app.Storage{
		DB: db,
		US: sf.UserStorage(ctx),
		OS: sf.OrderStorage(ctx),
		WS: sf.WithdrawStorage(ctx),
	}
	if err := app.New(settings, s).Run(ctx); err != nil {
		logger.Log.Fatal(err.Error())
	}
}
