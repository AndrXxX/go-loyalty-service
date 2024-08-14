package main

import (
	fl "flag"
	"github.com/AndrXxX/go-loyalty-service/internal/config"
)

func parseFlags(c *config.Config) {
	fl.StringVar(&c.RunAddress, "a", c.RunAddress, "Run address host:port")
	fl.StringVar(&c.DatabaseURI, "d", c.DatabaseURI, "Database URI")
	fl.StringVar(&c.AccrualSystemAddress, "r", c.AccrualSystemAddress, "Accrual System Address")
	fl.Parse()
}
