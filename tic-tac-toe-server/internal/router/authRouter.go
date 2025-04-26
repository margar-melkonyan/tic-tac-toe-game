package router

import (
	"github.com/go-chi/chi"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/handler/controller"
)

func authRouterGroup(auth chi.Router) {
	auth.Post("/sign-in", controller.SingIn)
	auth.Post("/sign-up", controller.SignUp)
}
