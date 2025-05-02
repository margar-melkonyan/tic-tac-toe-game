package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/handler/ws"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/helper"
)

func roomsRouterGroup(rooms chi.Router) {
	rooms.Post("/", dependencies.RoomHandler.CreateRoom)
	rooms.Get("/{id}", ws.EnterRoom(&dependencies.RoomHandler))
	rooms.Delete("/{id}", dependencies.RoomHandler.DestroyRoom)
	rooms.Get("/all", func(w http.ResponseWriter, r *http.Request) {
		resp := helper.Response{}
		resp.Data = ws.GetRooms()
		resp.ResponseWrite(w, r, http.StatusOK)
	})
}
