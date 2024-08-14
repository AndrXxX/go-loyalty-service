package main

import (
	"context"
	"database/sql"
	"github.com/AndrXxX/go-loyalty-service/internal/app"
	"github.com/AndrXxX/go-loyalty-service/internal/config"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"github.com/asaskevich/govalidator"
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

	var db *sql.DB // TODO
	a := app.New(settings, db)
	if err := a.Run(ctx); err != nil {
		logger.Log.Fatal(err.Error())
	}
}
