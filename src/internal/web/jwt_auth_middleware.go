package web

import (
	"TicTacToe/internal/service"
	"context"
	"net/http"
	"strings"
)

type contextKey string

const UserIDKey contextKey = "user_id"

type JWTAuthenticator struct {
	authService service.AuthService
}

func NewJWTAuthenticator(authService service.AuthService) *JWTAuthenticator {
	return &JWTAuthenticator{
		authService: authService,
	}
}

func (a *JWTAuthenticator) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			sendError(w, "authorization required", http.StatusUnauthorized)
			return
		}

		const prefix = "Bearer "
		if !strings.HasPrefix(authHeader, prefix) {
			sendError(w, "invalid auth scheme", http.StatusBadRequest)
			return
		}

		tokenString := authHeader[len(prefix):]
		user, err := a.authService.GetUserByAccessToken(r.Context(), tokenString)
		if err != nil {
			sendError(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, user.ID)
		next(w, r.WithContext(ctx))
	}
}
