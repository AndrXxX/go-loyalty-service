package middlewares

import (
	"context"
	"github.com/AndrXxX/go-loyalty-service/internal/enums"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"go.uber.org/zap"
	"net/http"
)

type isAuthorized struct {
	ts tokenService
}

func (m *isAuthorized) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie(enums.AuthToken)
		if err != nil {
			logger.Log.Error("failed to get auth token cookie", zap.Error(err))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		userId, err := m.ts.Decrypt(token.Value)
		if err != nil {
			logger.Log.Error("failed to decrypt token from cookie", zap.Error(err))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, enums.UserID, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func IsAuthorized(ts tokenService) *isAuthorized {
	return &isAuthorized{ts}
}
