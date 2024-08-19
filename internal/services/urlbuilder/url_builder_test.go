package metricurlbuilder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMetricURLBuilderBuildURL(t *testing.T) {
	tests := []struct {
		host     string
		endpoint string
		params   map[string]string
		want     string
	}{
		{
			host:     "localhost:8080",
			endpoint: "/api/orders/{number}",
			params:   map[string]string{"number": "123"},
			want:     "http://localhost:8080/api/orders/123",
		},
		{
			host:     "http://localhost:8080",
			endpoint: "/api/orders/{number}",
			params:   map[string]string{"number": "123"},
			want:     "http://localhost:8080/api/orders/123",
		},
		{
			host:     "https://localhost:8080",
			endpoint: "/api/orders/{number}",
			params:   map[string]string{"number": "123"},
			want:     "https://localhost:8080/api/orders/123",
		},
		{
			host:     "host",
			endpoint: "/api/orders/{number}",
			params:   map[string]string{"number": "123"},
			want:     "http://host/api/orders/123",
		},
		{
			host:     "host",
			endpoint: "/api/test",
			params:   map[string]string{},
			want:     "http://host/api/test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			b := New(tt.host)
			assert.Equal(t, tt.want, b.Build(tt.endpoint, tt.params))
		})
	}
}

func TestNewMetricURLBuilder(t *testing.T) {
	tests := []struct {
		host string
		want *urlBuilder
	}{
		{
			host: "host",
			want: &urlBuilder{host: "http://host"},
		},
		{
			host: "http://host",
			want: &urlBuilder{host: "http://host"},
		},
		{
			host: "https://host",
			want: &urlBuilder{host: "https://host"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.host, func(t *testing.T) {
			b := New(tt.host)
			assert.Equal(t, tt.want, b)
		})
	}
}
