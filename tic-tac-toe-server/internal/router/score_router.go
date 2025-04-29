package router

import (
	"github.com/go-chi/chi"
)

func scoresRouterGroup(scores chi.Router) {
	scores.Get("/", dependencies.ScoreHandler.GetCurrentUserScores)
}
