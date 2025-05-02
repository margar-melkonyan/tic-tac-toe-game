package ws

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
	http_handler "github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/handler/http"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/service"
)

var ws *service.WSServer

func EnterRoom(h *http_handler.RoomHandler) http.HandlerFunc {
	ws = service.NewWsServer()
	return func(w http.ResponseWriter, r *http.Request) {
		resp := h.GetRoom(w, r)
		room, ok := resp.Data.(*common.Room)
		if !ok {
			resp.Data = nil
			resp.Errors = "room doesn't exists"
			resp.ResponseWrite(w, r, http.StatusNotFound)
			return
		}
		currentUser, ok := r.Context().Value(common.USER).(*common.User)
		if !ok {
			resp.Data = nil
			resp.Errors = "you should be authorized"
			resp.ResponseWrite(w, r, http.StatusUnauthorized)
			return
		}
		conn, err := service.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			slog.Error(err.Error())
			return
		}
		ws.RefreshConnection(currentUser, room, conn)
		defer ws.CloseConnection(room.ID, conn)
		for {
			if resp.Errors != nil {
				err := conn.WriteMessage(
					websocket.CloseMessage,
					websocket.FormatCloseMessage(1011, "cannot find room"),
				)
				if err != nil {
					slog.Error("Error writing closing message:", err)
				}
				resp.ResponseWrite(w, r, http.StatusInternalServerError)
				return
			}
			if ws.GameLoop(
				currentUser,
				room,
				conn,
			) {
				break
			}
		}
	}
}

func GetRooms() map[uint64]*service.RoomServer {
	return ws.Rooms
}
