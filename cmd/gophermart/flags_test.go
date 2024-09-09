package main

import (
	"flag"
	"github.com/AndrXxX/go-loyalty-service/internal/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type testCase struct {
	name   string
	config *config.Config
	flags  []string
	want   *config.Config
}

func Test_parseFlags(t *testing.T) {
	tests := []testCase{
		{
			name:   "Empty flags",
			config: &config.Config{RunAddress: "host"},
			flags:  []string{},
			want:   &config.Config{RunAddress: "host"},
		},
		{
			name:   "-a=new-host",
			config: &config.Config{RunAddress: "host"},
			flags:  []string{"-a", "new-host"},
			want:   &config.Config{RunAddress: "new-host"},
		},
		{
			name:   "-a=new-host -d=test -r=test2",
			config: &config.Config{},
			flags:  []string{"-a", "new-host", "-d", "test", "-r", "test2"},
			want:   &config.Config{RunAddress: "new-host", DatabaseURI: "test", AccrualSystemAddress: "test2"},
		},
	}
	for _, tt := range tests {
		run(t, tt)
	}
}

func run(t *testing.T, tt testCase) {
	t.Run(tt.name, func(t *testing.T) {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		os.Args = os.Args[:1]
		os.Args = append(os.Args[:1], tt.flags...)
		parseFlags(tt.config)
		assert.Equal(t, tt.want, tt.config)
	})
}
