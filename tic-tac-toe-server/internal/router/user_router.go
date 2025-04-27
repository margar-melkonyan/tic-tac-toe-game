package router

import (
	"github.com/go-chi/chi"
)

func usersRouterGroup(users chi.Router) {
	users.Get("/current", dependencies.UserHandler.GetCurrentUser)
}
