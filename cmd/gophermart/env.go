package main

import (
	"github.com/AndrXxX/go-loyalty-service/internal/config"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
)

type EnvConfig struct {
	RunAddress           string `env:"RUN_ADDRESS"`
	DatabaseURI          string `env:"DATABASE_URI"`
	AccrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
}

func parseEnv(c *config.Config) {
	cfg := EnvConfig{
		RunAddress:           c.RunAddress,
		DatabaseURI:          c.DatabaseURI,
		AccrualSystemAddress: c.AccrualSystemAddress,
	}
	err := env.Parse(&cfg)
	if err != nil {
		logger.Log.Error("Error on parse EnvConfig", zap.Error(err))
		return
	}
	c.RunAddress = cfg.RunAddress
	c.DatabaseURI = cfg.DatabaseURI
	c.AccrualSystemAddress = cfg.AccrualSystemAddress
}
