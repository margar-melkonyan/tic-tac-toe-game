package controller

import (
	"net/http"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/helper"
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

func (h *ScoreHandler) GetCurrentUserScores(w http.ResponseWriter, r *http.Request) {
	resp := helper.Response{}
	scores, err := h.service.GetCurrentUserScores(r.Context())
	if err != nil {
		resp.ResponseWriter(w, r, http.StatusInternalServerError)
		return
	}
	resp.Data = scores
	resp.ResponseWriter(w, r, http.StatusOK)
}
