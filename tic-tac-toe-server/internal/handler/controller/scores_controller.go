package controller

import (
	"net/http"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/service"
)

type ScoreHandler struct {
	service service.ScoreService
}

func NewScoreHandler(service service.ScoreService) *ScoreHandler {
	return &ScoreHandler{
		service: service,
	}
}

func (h *ScoreHandler) GetCurrentUserScores(w http.ResponseWriter, r *http.Request) {}
