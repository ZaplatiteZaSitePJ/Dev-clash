package middleware

import (
	"context"
	"dev-clash/pkg/logger"
	custom_errors "dev-clash/pkg/server_utils/errors"
	"fmt"
	"net/http"
	"strings"
)
type contextKey string

const userIDKey contextKey = "user_id"

type TokenValidator interface {
	ValidateAccess(token string) (int64, error)
}

func JWTMiddleware(validator TokenValidator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authError := custom_errors.New(fmt.Errorf("unauthorized"), 401)
			authError.AddLogData("unauthorized")
			authError.AddResponseData("unauthorized")
			
			authHeader := r.Header.Get("Authorization")	
			if !strings.HasPrefix(authHeader, "Bearer ") {
				custom_errors.ErrorResponse(w, authError, logger.GetLoger())
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			userID, err := validator.ValidateAccess(token)
			if err!=nil {
				custom_errors.ErrorResponse(w, authError, logger.GetLoger())
				return
			}

			ctx := context.WithValue(r.Context(), userIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}