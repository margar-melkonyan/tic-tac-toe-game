package service

import (
	"context"
	"errors"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/repository"
)

type ScoreService struct {
	scoreRepo repository.ScoreRepository
	userRepo  repository.UserRepository
}

func NewScoreService(
	scoreRepo repository.ScoreRepository,
	userRepo repository.UserRepository,
) *ScoreService {
	return &ScoreService{
		scoreRepo: scoreRepo,
		userRepo:  userRepo,
	}
}

func (service *ScoreService) GetCurrentUserScores(ctx context.Context) ([]*common.Score, error) {
	email, ok := ctx.Value("user_email").(string)
	if !ok {
		return nil, errors.New("user email is not valid")
	}
	user, err := service.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return service.scoreRepo.GetAllByUser(ctx, user)
}
