package controller

import (
	"net/http"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/service"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) SingIn(w http.ResponseWriter, r *http.Request) {

}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {

}
