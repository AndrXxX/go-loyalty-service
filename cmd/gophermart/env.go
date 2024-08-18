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
	AuthKey              string `env:"AUTH_SECRET_KEY"`
	AuthKeyExpired       int    `env:"AUTH_SECRET_KEY_EXPIRED"`
	PasswordKey          string `env:"PASSWORD_SECRET_KEY"`
}

func parseEnv(c *config.Config) {
	cfg := EnvConfig{
		RunAddress:           c.RunAddress,
		DatabaseURI:          c.DatabaseURI,
		AccrualSystemAddress: c.AccrualSystemAddress,
		AuthKey:              c.AuthKey,
		AuthKeyExpired:       c.AuthKeyExpired,
		PasswordKey:          c.PasswordKey,
	}
	err := env.Parse(&cfg)
	if err != nil {
		logger.Log.Error("Error on parse EnvConfig", zap.Error(err))
		return
	}
	c.RunAddress = cfg.RunAddress
	c.DatabaseURI = cfg.DatabaseURI
	c.AccrualSystemAddress = cfg.AccrualSystemAddress
	c.AuthKey = cfg.AuthKey
	c.AuthKeyExpired = cfg.AuthKeyExpired
	c.PasswordKey = cfg.PasswordKey
}
