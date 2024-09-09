package luhn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlgorithmCheckerCheck(t *testing.T) {
	tests := []struct {
		name string
		val  string
		want bool
	}{
		{
			name: "Test FALSE with 12345",
			val:  "12345",
			want: false,
		},
		{
			name: "Test FALSE with 111111111",
			val:  "111111111",
			want: false,
		},
		{
			name: "Test TRUE with 9278923470",
			val:  "9278923470",
			want: true,
		},
		{
			name: "Test TRUE with 12345678903",
			val:  "12345678903",
			want: true,
		},
		{
			name: "Test TRUE with 346436439",
			val:  "346436439",
			want: true,
		},
		{
			name: "Test TRUE with 4561261212345467",
			val:  "4561261212345467",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := luhnAlgorithmChecker{}
			assert.Equal(t, tt.want, c.Check(tt.val))
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *luhnAlgorithmChecker
	}{
		{
			name: "Test OK",
			want: &luhnAlgorithmChecker{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Checker())
		})
	}
}
