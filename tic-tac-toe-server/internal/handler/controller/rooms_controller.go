package controller

import (
	"net/http"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/service"
)

type RoomHandler struct {
	service service.RoomService
}

func NewRoomHandler(service service.RoomService) *RoomHandler {
	return &RoomHandler{
		service: service,
	}
}

func (h *RoomHandler) GetRooms(w http.ResponseWriter, r *http.Request)    {}
func (h *RoomHandler) EnterRoom(w http.ResponseWriter, r *http.Request)   {}
func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request)  {}
func (h *RoomHandler) DestroyRoom(w http.ResponseWriter, r *http.Request) {}
