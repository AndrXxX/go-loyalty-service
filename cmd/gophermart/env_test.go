package main

import (
	"github.com/AndrXxX/go-loyalty-service/internal/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_parseEnv(t *testing.T) {
	tests := []struct {
		name   string
		config *config.Config
		env    map[string]string
		want   *config.Config
	}{
		{
			name:   "Empty env",
			config: &config.Config{RunAddress: "host"},
			env:    map[string]string{},
			want:   &config.Config{RunAddress: "host"},
		},
		{
			name:   "RUN_ADDRESS=new-host",
			config: &config.Config{RunAddress: "host"},
			env:    map[string]string{"RUN_ADDRESS": "new-host"},
			want:   &config.Config{RunAddress: "new-host"},
		},
		{
			name:   "RUN_ADDRESS=new-host DATABASE_URI=test ACCRUAL_SYSTEM_ADDRESS=test2",
			config: &config.Config{},
			env:    map[string]string{"RUN_ADDRESS": "new-host", "DATABASE_URI": "test", "ACCRUAL_SYSTEM_ADDRESS": "test2"},
			want:   &config.Config{RunAddress: "new-host", DatabaseURI: "test", AccrualSystemAddress: "test2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			for k, v := range tt.env {
				_ = os.Setenv(k, v)
			}
			parseEnv(tt.config)
			assert.Equal(t, tt.want, tt.config)
		})
	}
}