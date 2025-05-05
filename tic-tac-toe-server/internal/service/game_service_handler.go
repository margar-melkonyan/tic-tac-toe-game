package service

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
	"golang.org/x/crypto/bcrypt"
)

func (ws *WSServer) handleStep(room *common.RoomSessionResponse, request *GameRequest) {
	rawSymbolPosition, ok := request.Data.(map[string]interface{})
	if ok {
		currentRoom := ws.Rooms[room.ID]
		symbolPosition := &SymbolPosition{
			ID:     rawSymbolPosition["id"].(string),
			Symbol: rawSymbolPosition["symbol"].(string),
		}
		opositeSymbol := opositeSymbol(symbolPosition.Symbol)
		currentRoom.Positions = append(currentRoom.Positions, symbolPosition)
		response := &GameReponse{
			Action: getPositionsAction,
			Data: map[string]interface{}{
				"positions": currentRoom.Positions,
			},
			Symbol: opositeSymbol,
		}
		ws.jsonToAll(room, response)
	}
}

func (ws *WSServer) handleResetGame(room *common.RoomSessionResponse) {
	ws.Rooms[room.ID].Positions = make([]*SymbolPosition, 0)
	response := &GameReponse{
		Action: resetGameAction,
	}
	ws.jsonToAll(room, response)
}

func (ws *WSServer) handleBorderResize(
	currentUserId uuid.UUID,
	room *common.RoomSessionResponse,
	request *GameRequest,
	message []byte,
) {
	if currentUserId == room.CreatorID {
		ws.Rooms[room.ID].BorderSize = request.BorderSize
		ws.broadcastMessageToOther(currentUserId, room, message)
	}
}

func (ws *WSServer) handleSelectSymbol(
	currentUserId uuid.UUID,
	room *common.RoomSessionResponse,
	request *GameRequest,
) {
	currentRoom := ws.Rooms[room.ID]
	for id, user := range currentRoom.Users {
		if user.Symbol != "" {
			continue
		}
		if user.ID == currentUserId {
			currentRoom.Users[id].Symbol = request.Symbol
		} else {
			currentRoom.Users[id].Symbol = opositeSymbol(request.Symbol)
		}
	}
	response := &GameReponse{
		Action: selectedSymbolAction,
		Symbol: opositeSymbol(request.Symbol),
	}
	ws.jsonToOther(currentUserId, room, response)
}

func (ws *WSServer) handleNewConnection(
	currentUserId uuid.UUID,
	room *common.RoomSessionResponse,
	request *GameRequest,
	conn *websocket.Conn,
) {
	currentRoom := ws.Rooms[room.ID]
	if room.IsPrivate != nil && room.Password != "" {
		err := bcrypt.CompareHashAndPassword([]byte(room.Password), []byte(request.Password))
		if err != nil {
			conn.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.ClosePolicyViolation, "password is not valid"),
			)
			ws.CloseConnection(room.ID, conn)
		}
	}
	response := &GameReponse{
		Action: newConnectionToRoomAction,
		UserID: &currentUserId,
	}
	ws.jsonToAll(room, response)
	if currentRoom.BorderSize == 0 {
		currentRoom.BorderSize = 3
	}
	response = &GameReponse{
		Action:      resizeAction,
		BoarderSize: currentRoom.BorderSize,
	}
	ws.jsonToAll(room, response)
	currentPlayerStep := currentRoom.Users[0].Symbol
	if len(currentRoom.Positions) != 0 {
		currentPlayerStep = opositeSymbol(currentRoom.Positions[len(currentRoom.Positions)-1].Symbol)
	}
	response = &GameReponse{
		Action: getPositionsAction,
		Data: map[string]interface{}{
			"positions": currentRoom.Positions,
		},
		Symbol: currentPlayerStep,
	}
	ws.jsonToAll(room, response)
	currentRoom.GameStatus = chooseSymbolStatus
	response = &GameReponse{
		Action: chooseSymbolAction,
		UserID: &currentRoom.Users[0].ID,
	}
	ws.jsonToAll(room, response)
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
	mainSymbol := ""
	for _, user := range currentRoom.Users {
		if user.Symbol != "" {
			mainSymbol = user.Symbol
		}
		if user.Symbol == "" {
			secondarySymbol := opositeSymbol(mainSymbol)
			if secondarySymbol != "" {
				resp := &GameReponse{
					Action: syncSymbolAction,
					Symbol: secondarySymbol,
				}
				raw, err := json.Marshal(resp)
				if err == nil {
					user.Connection.WriteMessage(
						websocket.TextMessage,
						raw,
					)
				}
			}
		}
	}
}

func (ws *WSServer) handleGameEnd(
	currentUserID uuid.UUID,
	room *common.RoomSessionResponse,
	request *GameRequest,
) {
	rawGameEndJson, ok := request.Data.(map[string]interface{})
	if ok {
		score := &common.Score{
			IsWon:    rawGameEndJson["is_won"].(float64),
			UserID:   rawGameEndJson["user_id"].(string),
			Nickname: rawGameEndJson["versus_player_nickname"].(string),
		}
		ws.ScoreService.scoreRepo.Create(context.Background(), score)
	}
	ws.jsonToOther(
		currentUserID,
		room, &GameReponse{
			Action: restartGameAction,
		})
}

func (ws *WSServer) handleCloseRoom(
	room *RoomServer,
) {
	for _, user := range room.Users {
		ws.CloseConnection(room.ID, user.Connection)
	}

	ws.Mu.Lock()
	defer ws.Mu.Unlock()
	delete(ws.Rooms, room.ID)
}

func (ws *WSServer) jsonToAll(room *common.RoomSessionResponse, response interface{}) {
	raw, err := json.Marshal(response)
	if err == nil {
		ws.broadcastMessageToAll(room, raw)
	}
}

func (ws *WSServer) jsonToOther(currentUserID uuid.UUID, room *common.RoomSessionResponse, response interface{}) {
	raw, err := json.Marshal(response)
	if err == nil {
		ws.broadcastMessageToOther(currentUserID, room, raw)
	}
}
