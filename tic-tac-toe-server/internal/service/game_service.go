package service

import (
	"bytes"
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

var wsRooms map[uint64]*RoomServer

type GameRequest struct {
	Action string         `json:"action"`
	Data   SymbolPosition `json:"data"`
}

type GameReponse struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

func NewWsServer() *WSServer {
	return &WSServer{
		Rooms: make(map[uint64]*RoomServer),
	}
}

func GetCurrentRoomInfo(roomID uint64) *RoomServer {
	return wsRooms[roomID]
}

func (ws *WSServer) GameLoop(currentUser *common.User, room *common.RoomSessionResponse, conn *websocket.Conn) bool {
	if ws.isRoomFull(currentUser.ID, room.ID, conn) {
		return true
	}
	ws.addUser(currentUser, room, conn)
	_, p, err := conn.ReadMessage()
	if err != nil {
		slog.Error("ReadMessage error:", slog.String("error", err.Error()))
		return true
	}
	if len(p) == 0 {
		slog.Warn("Received empty message")
		return true
	}
	var request GameRequest
	bytesReader := bytes.NewReader(p)
	if err := json.NewDecoder(bytesReader).Decode(&request); err != nil {
		slog.Error("GameRequest")
	}
	if ws.Rooms != nil {
		wsRooms = ws.Rooms
	}
	switch request.Action {
	case "step":
		ws.Rooms[room.ID].Positions = append(ws.Rooms[room.ID].Positions, &request.Data)
		response := &GameReponse{
			Action: "get positions",
			Data: map[string]interface{}{
				"positions": ws.Rooms[room.ID].Positions,
			},
		}
		raw, err := json.Marshal(response)
		if err == nil {
			ws.broadcastMessageToAll(currentUser.ID, room, raw)
		}
	case "reset game":
		ws.Rooms[room.ID].Positions = make([]*SymbolPosition, 0)
		response := &GameReponse{
			Action: "reset game",
		}
		raw, err := json.Marshal(response)
		if err == nil {
			ws.broadcastMessageToAll(currentUser.ID, room, raw)
		}
	}
	return false
}

func (ws *WSServer) RefreshConnection(currentUser *common.User, room *common.RoomSessionResponse, conn *websocket.Conn) {
	if currentUser == nil {
		slog.Error("currentUser  is nil")
		return
	}
	if room == nil {
		slog.Error("room is nil")
		return
	}
	if ws.isUserInRoom(currentUser.ID, room.ID) {
		for _, user := range ws.Rooms[room.ID].Users {
			if user.ID == currentUser.ID && user.Connection != conn {
				user.Connection = conn
				user.IsConnected = true
				break
			}
		}
	}
}

func (ws *WSServer) CloseConnection(roomID uint64, conn *websocket.Conn) {
	ws.Mu.Lock()
	defer ws.Mu.Unlock()

	room, exists := ws.Rooms[roomID]
	if !exists || room == nil {
		slog.Warn("attempted to close connection for non-existent room", slog.Uint64("room_id", roomID))
		return
	}

	for _, user := range room.Users {
		if user.Connection == conn {
			conn.Close()
			user.Connection = nil
			user.IsConnected = false
			break
		}
	}
}

func (ws *WSServer) addUser(currentUser *common.User, room *common.RoomSessionResponse, conn *websocket.Conn) {
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
				Connection:  conn,
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

func (ws *WSServer) isRoomFull(userID uuid.UUID, roomID uint64, conn *websocket.Conn) bool {
	if ws.Rooms[roomID] != nil && len(ws.Rooms[roomID].Users) == 2 && !ws.isUserInRoom(userID, roomID) {
		err := ws.Rooms[roomID].Users[0].Connection.WriteMessage(websocket.CloseMessage, []byte(`{"error":"room is full"}`))
		if err != nil {
			slog.Error(
				"Error sending room full message:",
				slog.String("error", err.Error()),
			)
		}
		return true
	}
	return false
}

func (ws *WSServer) broadcastMessageToOther(
	currentUserID uuid.UUID,
	room *common.RoomSessionResponse,
	raw []byte,
) {
	ws.Mu.Lock()
	defer ws.Mu.Unlock()
	roomData, ok := ws.Rooms[room.ID]
	if !ok || roomData == nil {
		log.Printf("broadcastToSymbolPosition: room %d not found", room.ID)
		return
	}
	for _, currentUser := range roomData.Users {
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

func (ws *WSServer) broadcastMessageToAll(
	currentUserID uuid.UUID,
	room *common.RoomSessionResponse,
	raw []byte,
) {
	ws.Mu.Lock()
	defer ws.Mu.Unlock()
	roomData, ok := ws.Rooms[room.ID]
	if !ok || roomData == nil {
		log.Printf("broadcastToSymbolPosition: room %d not found", room.ID)
		return
	}
	for _, currentUser := range roomData.Users {
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
