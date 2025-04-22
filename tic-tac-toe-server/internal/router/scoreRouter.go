package router

import (
	"github.com/go-chi/chi"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/handler/controller"
)

func scoresRouterGroup(scores chi.Router) {
	scores.Get("/scores", controller.GetCurrentUserScore)
}
