package http_handler

import (
	"net/http"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/helper"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	resp := helper.Response{}
	user, err := h.service.GetCurrentUser(r.Context())
	if err != nil {
		resp.Message = err.Error()
		resp.ResponseWrite(w, r, http.StatusConflict)
		return
	}
	resp.Data = user
	resp.ResponseWrite(w, r, http.StatusOK)
}
