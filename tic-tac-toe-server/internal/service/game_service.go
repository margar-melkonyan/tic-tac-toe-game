package service

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
)

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

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ConnectedUser struct {
	ID          uuid.UUID       `json:"id"`
	Name        string          `json:"name"`
	Symbol      string          `json:"symbol"`
	Connection  *websocket.Conn `json:"connection"`
	IsConnected bool            `json:"is_connected"`
}

type SymbolPosition struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
}

type RoomServer struct {
	ID        uint64            `json:"id"`
	Users     []*ConnectedUser  `json:"users"`
	Positions []*SymbolPosition `json:"symbol_positions"`
}

type WSServer struct {
	Rooms map[uint64]*RoomServer `json:"rooms"`
	Mu    sync.Mutex
}

type GameResponse struct {
	Data string `json:"data"`
}

func NewWsServer() *WSServer {
	return &WSServer{
		Rooms: make(map[uint64]*RoomServer),
	}
}

func (ws *WSServer) GameLoop(currentUser *common.User, room *common.Room, conn *websocket.Conn) bool {
	if ws.isRoomFull(currentUser.ID, room.ID) {
		return true
	}
	ws.RefreshConnection(currentUser, room, conn)
	ws.addUser(currentUser, room, conn)

	_, p, err := conn.ReadMessage()
	if err != nil {
		log.Println("ReadMessage error:", err)
		return true
	}
	resp := GameResponse{
		Data: string(p),
	}
	raw, err := json.Marshal(resp)
	if err != nil {
		slog.Error(err.Error())
		return true
	}
	ws.broadcastToSymbolPosition(currentUser.ID, room, raw)
	return false
}

func (ws *WSServer) RefreshConnection(currentUser *common.User, room *common.Room, conn *websocket.Conn) {
	if ws.isUserInRoom(currentUser.ID, room.ID) {
		for _, user := range ws.Rooms[room.ID].Users {
			if user.ID == currentUser.ID {
				user.Connection = conn // Обновляем соединение
				user.IsConnected = true
				break
			}
		}
	}
}

func (ws *WSServer) CloseConnection(roomID uint64, conn *websocket.Conn) {
	for _, user := range ws.Rooms[roomID].Users {
		if user.Connection == conn {
			user.Connection.Close()
			user.IsConnected = false
			break
		}
	}
}

func (ws *WSServer) addUser(currentUser *common.User, room *common.Room, conn *websocket.Conn) {
	ws.Mu.Lock()
	defer ws.Mu.Unlock()

	if ws.Rooms[room.ID] == nil {
		ws.Rooms[room.ID] = &RoomServer{
			ID:        room.ID,
			Users:     make([]*ConnectedUser, 0),
			Positions: make([]*SymbolPosition, 0),
		}
	}
	if !ws.isUserInRoom(currentUser.ID, room.ID) {
		ws.Rooms[room.ID].Users = append(
			ws.Rooms[room.ID].Users,
			&ConnectedUser{
				ID:          currentUser.ID,
				Name:        currentUser.Name,
				Symbol:      "",
				Connection:  conn, // Сохраняем текущее соединение
				IsConnected: true,
			},
		)
	}
}

func (ws *WSServer) isUserInRoom(userID uuid.UUID, roomID uint64) bool {
	room, exists := ws.Rooms[roomID]
	if !exists {
		return false
	}
	for _, user := range room.Users {
		if user.ID == userID {
			return true
		}
	}
	return false
}

func (ws *WSServer) isRoomFull(userID uuid.UUID, roomID uint64) bool {
	if ws.Rooms[roomID] != nil && len(ws.Rooms[roomID].Users) == 2 && !ws.isUserInRoom(userID, roomID) {
		err := ws.Rooms[roomID].Users[0].Connection.WriteMessage(websocket.CloseMessage, []byte(`{"error":"room is full"}`))
		if err != nil {
			slog.Error("Error sending room full message:", err)
		}
		return true
	}
	return false
}

func (ws *WSServer) broadcastToSymbolPosition(
	currentUserID uuid.UUID,
	room *common.Room,
	raw []byte,
) {
	for _, currentUser := range ws.Rooms[room.ID].Users {
		if currentUser.ID != currentUserID {
			if currentUser.Connection != nil {
				if err := currentUser.Connection.WriteMessage(
					websocket.TextMessage,
					raw,
				); err != nil {
					log.Println("WriteMessage error:", err)
				}
			}
		}
	}
}
