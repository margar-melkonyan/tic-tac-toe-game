package middleware

import (
	"bufio"
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common/dependency"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/helper"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/service"
)

type responseWriterWrapper struct {
	http.ResponseWriter
}

func (w *responseWriterWrapper) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if h, ok := w.ResponseWriter.(http.Hijacker); ok {
		return h.Hijack()
	}
	return nil, nil, errors.New("the ResponseWriter doesn't support hijacking")
}

func AuthMiddleware(dependency *dependency.AppDependencies) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" && r.URL.Query().Get("token") != "" {
				token = r.URL.Query().Get("token")
			}
			if token == "" {
				resp := helper.Response{}
				resp.Message = "You should be authorized!"
				resp.ResponseWrite(w, r, http.StatusUnauthorized)
				return
			}
			claims, err := service.CheckTokenIsNotExpired(token)
			if err != nil {
				resp := helper.Response{}
				resp.Message = err.Error()
				resp.ResponseWrite(w, r, http.StatusUnauthorized)
				return
			}
			user, err := dependency.GlobalRepositories.UserRepository.FindByEmail(r.Context(), claims.Sub.Email)
			if err != nil {
				resp := helper.Response{}
				resp.Message = err.Error()
				resp.ResponseWrite(w, r, http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), common.USER_MAIL, claims.Sub.Email)
			ctx = context.WithValue(ctx, common.USER, user)
			r = r.WithContext(ctx)

			wrappedWriter := &responseWriterWrapper{w}
			next.ServeHTTP(wrappedWriter, r)
		})
	}
}
