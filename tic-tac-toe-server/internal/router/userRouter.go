package router

import (
	"github.com/go-chi/chi"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/handler/controller"
)

func usersRouterGroup(users chi.Router) {
	users.Get("/current", controller.GetCurrentUser)
}
