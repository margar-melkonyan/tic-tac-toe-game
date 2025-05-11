// Package router предоставляет функциональность для настройки маршрутизации HTTP запросов.
package router

import (
	"github.com/go-chi/chi"
)

func authRouterGroup(auth chi.Router) {
	auth.Post("/sign-in", dependencies.AuthHandler.SingIn)
	auth.Post("/sign-up", dependencies.AuthHandler.SignUp)
}
