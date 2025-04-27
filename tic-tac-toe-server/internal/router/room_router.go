package router

import (
	"github.com/go-chi/chi"
)

func roomsRouterGroup(rooms chi.Router) {
	rooms.Get("/", dependencies.RoomHandler.GetRooms)
	rooms.Post("/{id}", dependencies.RoomHandler.EnterRoom)
	rooms.Post("/", dependencies.RoomHandler.CreateRoom)
	rooms.Delete("/{id}", dependencies.RoomHandler.DestroyRoom)
}
