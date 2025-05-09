package service

import (
	"encoding/json"
	"log"
	"log/slog"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
)

func opositeSymbol(mainSymbol string) string {
	opositeSymbol := ""
	if mainSymbol == "" {
		opositeSymbol = ""
	}

	if mainSymbol == "X" {
		opositeSymbol = "O"
	}
	if mainSymbol == "O" {
		opositeSymbol = "X"
	}

	return opositeSymbol
}

func (ws *WSServer) addUser(currentUser *common.User, room *common.RoomSessionResponse, conn *websocket.Conn) {
	ws.Mu.Lock()
	defer ws.Mu.Unlock()

	if ws.Rooms[room.ID] == nil {
		ws.Rooms[room.ID] = &RoomServer{
			ID:         room.ID,
			Users:      make([]*ConnectedUser, 0),
			Positions:  make([]*SymbolPosition, 0),
			BorderSize: DEFAULT_BORDER_SIZE,
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
		err := conn.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseTryAgainLater, "room is full"),
		)
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

func (ws *WSServer) changeGameStatus(roomId uint64) {
	currentRoom, exists := ws.Rooms[roomId]
	if !exists {
		return
	}
	if len(currentRoom.Users) == 2 {
		var prev *ConnectedUser
		for _, user := range currentRoom.Users {
			if user.Symbol != "" {
				if prev != nil {
					prev.Symbol = opositeSymbol(user.Symbol)
				}
				break
			}
			prev = user
		}
	}
	isAllUserSelectedSymbol := 0
	for _, user := range currentRoom.Users {
		if user.Symbol != "" {
			isAllUserSelectedSymbol += 1
		}
	}
	if len(currentRoom.Users) == 2 && isAllUserSelectedSymbol == 2 {
		currentRoom.GameStatus = inProcessStatus
	}
	ws.Mu.Lock()
	defer ws.Mu.Unlock()

	firstPlayerSymbol := ""
	secondarySymbol := ""
	for _, user := range currentRoom.Users {
		if user.Symbol != "" {
			firstPlayerSymbol = user.Symbol
		}
		if user.Symbol == "" {
			secondarySymbol = opositeSymbol(firstPlayerSymbol)
			if secondarySymbol != "" {
				user.Symbol = secondarySymbol
				resp := &GameReponse{
					Action: syncSymbolAction,
					Symbol: secondarySymbol,
				}
				raw, err := json.Marshal(resp)
				if err == nil && user.Connection != nil {
					user.Connection.WriteMessage(
						websocket.TextMessage,
						raw,
					)
				}
			}
		}
	}
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
