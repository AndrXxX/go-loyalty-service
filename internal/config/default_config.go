package config

import (
	"github.com/AndrXxX/go-loyalty-service/internal/enums/defaults"
)

func NewConfig() *Config {
	return &Config{
		RunAddress:           defaults.RunAddress,
		LogLevel:             defaults.LogLevel,
		DatabaseURI:          "",
		AccrualSystemAddress: "",
		AuthKey:              "auth-secret-key",
		AuthKeyExpired:       defaults.AuthKeyExpired,
	}
}
