package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/helper"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/service"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "auth") || strings.Contains(r.URL.Path, "current-user") {
			token := r.Header.Get("Authorization")
			if token == "" {
				resp := helper.Response{}
				resp.Message = "You should be authorized!"
				resp.ResponseWriter(w, r, http.StatusUnauthorized)
				return
			}

			claims, err := service.CheckTokenIsNotExpired(token)
			if err != nil {
				resp := helper.Response{}
				resp.Message = err.Error()
				resp.ResponseWriter(w, r, http.StatusUnauthorized)
				return
			}

			req := context.WithValue(r.Context(), "user_email", claims.Sub.Email)
			r = r.WithContext(req)
		}

		next.ServeHTTP(w, r)
	})
}
