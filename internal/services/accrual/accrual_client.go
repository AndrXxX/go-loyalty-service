package accrual

import (
	"encoding/json"
	"github.com/AndrXxX/go-loyalty-service/internal/entities"
	"github.com/AndrXxX/go-loyalty-service/internal/services/gzipcompressor"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"go.uber.org/zap"
	"io"
	"strings"
)

const fetchRoute = "/api/orders/{number}"

type accrualClient struct {
	c  httpClient
	ub urlBuilder
}

func NewClient(c httpClient, ub urlBuilder) *accrualClient {
	return &accrualClient{c, ub}
}

func (c *accrualClient) Fetch(order string) (statusCode int, m *entities.Accrual) {
	url := c.ub.Build(fetchRoute, map[string]string{"number": order})

	resp, err := c.c.Get(url)
	if err != nil {
		logger.Log.Error("failed to get accrual", zap.Error(err))
		return resp.StatusCode, m
	}
	defer resp.Body.Close()

	contentEncoding := resp.Header.Get("Content-Encoding")
	sendsGzip := strings.Contains(contentEncoding, "gzip")
	if sendsGzip {
		cr, err := gzipcompressor.NewCompressReader(resp.Body)
		if err != nil {
			logger.Log.Error("Error creating gzip compressor", zap.Error(err))
			return resp.StatusCode, m
		}
		resp.Body = cr
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Error("failed to read body on get accrual", zap.Error(err))
		return resp.StatusCode, m
	}

	err = json.Unmarshal(body, &m)
	if err != nil {
		logger.Log.Error("failed to unmarshal body on get accrual", zap.Error(err))
		return resp.StatusCode, m
	}

	return resp.StatusCode, m
}
