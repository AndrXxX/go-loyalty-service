package middlewares

import (
	"github.com/AndrXxX/go-loyalty-service/internal/services/gzipcompressor"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type gzipMiddleware struct {
}

func (m *gzipMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ow := w

		acceptEncoding := r.Header.Get("Accept-Encoding")
		supportsGzip := strings.Contains(acceptEncoding, "gzip")
		if supportsGzip {
			cw := gzipcompressor.NewCompressWriter(w)
			ow = cw
			defer func() {
				_ = cw.Close()
			}()
		}

		contentEncoding := r.Header.Get("Content-Encoding")
		sendsGzip := strings.Contains(contentEncoding, "gzip")
		if sendsGzip {
			cr, err := gzipcompressor.NewCompressReader(r.Body)
			if err != nil {
				logger.Log.Error("Error creating gzip compressor", zap.Error(err))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			r.Body = cr
			defer func() {
				_ = cr.Close()
			}()
		}
		next.ServeHTTP(ow, r)
	})

}

func CompressGzip() *gzipMiddleware {
	return &gzipMiddleware{}
}
