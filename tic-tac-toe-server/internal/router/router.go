package router

import (
	"github.com/go-chi/chi"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common/dependency"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/handler/middleware"
)

var dependencies *dependency.AppDependencies

func NewRouter(deps *dependency.AppDependencies) *chi.Mux {
	dependencies = deps
	api := chi.NewMux()
	api.Use(middleware.Logger)

	api.Route("/auth", authRouterGroup)
	api.Route("/api", func(v1 chi.Router) {
		v1.Use(middleware.AuthMiddleware)
		v1.Route("/v1", func(r chi.Router) {
			r.Route("/rooms", roomsRouterGroup)
			r.Route("/users", usersRouterGroup)
			r.Route("/scores", scoresRouterGroup)
		})
	})

	return api
}
