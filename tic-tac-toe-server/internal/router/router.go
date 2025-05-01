package router

import (
	"github.com/go-chi/chi"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common/dependency"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/handler/middleware"
)

var dependencies *dependency.AppDependencies

func NewRouter(deps *dependency.AppDependencies) *chi.Mux {
	dependencies = deps
	route := chi.NewMux()
	route.Use(
		middleware.CorsMiddleware,
		middleware.Logger,
	)

	route.Route("/auth", authRouterGroup)
	route.Route("/api", func(api chi.Router) {
		api.Get("/v1/rooms", dependencies.RoomHandler.GetRooms)
		api.Route("/v1", func(v1 chi.Router) {
			v1.Use(middleware.AuthMiddleware(deps))
			v1.Route("/rooms", roomsRouterGroup)
			v1.Route("/users", usersRouterGroup)
			v1.Route("/scores", scoresRouterGroup)
		})
	})

	return route
}
