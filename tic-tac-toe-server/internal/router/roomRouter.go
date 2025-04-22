package router

import (
	"github.com/go-chi/chi"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/handler/controller"
)

func roomsRouterGroup(rooms chi.Router) {
	rooms.Get("/", controller.GetRooms)
	rooms.Post("/{id}", controller.EnterRoom)
	rooms.Post("/", controller.CreateRoom)
	rooms.Delete("/{id}", controller.DestroyRoom)
}
