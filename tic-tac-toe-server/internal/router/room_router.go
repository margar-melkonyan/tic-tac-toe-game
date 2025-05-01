package router

import (
	"github.com/go-chi/chi"
)

func roomsRouterGroup(rooms chi.Router) {
	rooms.Post("/", dependencies.RoomHandler.CreateRoom)
	rooms.Post("/{id}", dependencies.RoomHandler.EnterRoom)
	rooms.Delete("/{id}", dependencies.RoomHandler.DestroyRoom)
}
