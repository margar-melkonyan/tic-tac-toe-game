package router

import (
	"github.com/go-chi/chi"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/handler/ws"
)

func roomsRouterGroup(rooms chi.Router) {
	rooms.Post("/", dependencies.RoomHandler.CreateRoom)
	rooms.Get("/{id}/info", dependencies.RoomHandler.GetRoom(dependencies.WSServer))
	rooms.Get("/{id}", ws.EnterRoom(dependencies))
	rooms.Get("/my", dependencies.RoomHandler.GetMyRooms(dependencies.WSServer))
	rooms.Delete("/{id}", dependencies.RoomHandler.DestroyRoom)
}
