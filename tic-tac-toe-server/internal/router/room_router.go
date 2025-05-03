package router

import (
	"github.com/go-chi/chi"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/handler/ws"
)

func roomsRouterGroup(rooms chi.Router) {
	rooms.Post("/", dependencies.RoomHandler.CreateRoom)
	rooms.Get("/{id}/info", dependencies.RoomHandler.GetRoom)
	rooms.Get("/{id}", ws.EnterRoom(&dependencies.RoomHandler))
	rooms.Delete("/{id}", dependencies.RoomHandler.DestroyRoom)
}
