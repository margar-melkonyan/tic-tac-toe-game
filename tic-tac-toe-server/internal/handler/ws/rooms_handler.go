package ws

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
	http_handler "github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/handler/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// rooms{
//     room {
//         id,
//         users: [
//             {
//                 id,
//                 name,
//                 symbol
//             },
//             {
//                 id,
//                 name,
//                 symbol
//             },
//         ],
//         positions: [
//             {
//                id,
//                symbol,
//                 user_id
//             },
//             ...
//         ],
//     },
//     ...
// }

type ConnectedUser struct {
	ID         uuid.UUID
	Name       string
	Symbol     string
	Connection *websocket.Conn
}

type SymbolPosition struct {
	ID     string
	Symbol string
}

type RoomServer struct {
	ID        uint64
	Users     []*ConnectedUser
	Positions []*SymbolPosition
}

type WSServer struct {
	Rooms map[uint64]*RoomServer
	mu    sync.Mutex
}

var ws = WSServer{
	Rooms: make(map[uint64]*RoomServer),
}

func EnterRoom(h *http_handler.RoomHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := h.GetRoom(w, r)
		room := resp.Data.(*common.Room)
		currentUser, ok := r.Context().Value(common.USER).(*common.User)
		if !ok {
			resp.Data = nil
			resp.Errors = "you should be authorized"
			resp.ResponseWrite(w, r, http.StatusUnauthorized)
			return
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if ws.Rooms[room.ID] != nil && len(ws.Rooms[room.ID].Users) == 2 {
			conn.WriteMessage(websocket.TextMessage, []byte(`{"error":"room is full"}`))
			conn.Close()
			return
		}
		if err != nil {
			slog.Error(err.Error())
			return
		}
		defer conn.Close()

		if ws.Rooms[room.ID] == nil {
			ws.mu.Lock()
			ws.Rooms[room.ID] = &RoomServer{
				ID: room.ID,
				Users: []*ConnectedUser{
					{
						ID:         currentUser.ID,
						Name:       currentUser.Name,
						Symbol:     "",
						Connection: conn,
					},
				},
				Positions: make([]*SymbolPosition, 0),
			}
			ws.mu.Unlock()
		} else {
			users := ws.Rooms[room.ID].Users
			users = append(users, &ConnectedUser{
				ID:         currentUser.ID,
				Name:       currentUser.Name,
				Symbol:     "",
				Connection: conn,
			})
			ws.Rooms[room.ID].Users = users
		}

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
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Println("ReadMessage error:", err)
				break
			}
			for _, currentUser := range ws.Rooms[room.ID].Users {
				if currentUser.Connection != conn {
					if err := currentUser.Connection.WriteMessage(messageType, []byte(fmt.Sprintf("from server %s", string(p)))); err != nil {
						log.Println("WriteMessage error:", err)
						break
					}
				}
			}
		}
	}
}
