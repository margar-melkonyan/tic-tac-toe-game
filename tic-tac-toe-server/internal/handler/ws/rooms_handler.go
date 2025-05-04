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
		conn, err := service.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			slog.Error(
				"Error upgrading connection to websockets",
				slog.String("error", err.Error()),
			)
			return
		}
		resp := h.GetRoomInfo(r)
		room, isRoomExist := resp.Data.(*common.RoomSessionResponse)
		currentUser, isUserExist := r.Context().Value(common.USER).(*common.User)
		if !isRoomExist {
			err := conn.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(
					websocket.CloseInternalServerErr,
					"cannot find room",
				),
			)
			if err != nil {
				slog.Error(
					"Error writing closing message:",
					slog.String("error", err.Error()),
				)
			}
			conn.Close()
			return
		}

		if !isUserExist {
			err := conn.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(
					websocket.CloseInternalServerErr,
					"you should be authorized",
				),
			)
			if err != nil {
				slog.Error(
					"You should be authorized:",
					slog.String("error", err.Error()),
				)
			}
			conn.Close()
			return
		}
		defer ws.CloseConnection(room.ID, conn)
		ws.RefreshConnection(currentUser, room, conn)
		for {
			if ws.GameLoop(currentUser, room, conn) {
				break
			}
		}
	}
}
