package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfig(t *testing.T) {
	config := NewConfig()
	assert.NotEmpty(t, config.RunAddress)
	assert.NotEmpty(t, config.LogLevel)
}
