package service

import (
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/repository"
)

type ScoreService struct {
	repo repository.ScoreRepository
}

func NewScoreService(repoScore repository.ScoreRepository) *ScoreService {
	return &ScoreService{
		repo: repoScore,
	}
}

func (service *ScoreService) GetAll() []*common.Room {
	return nil
}

func (service *ScoreService) GetById() *common.Room {
	return nil
}

func (service *ScoreService) Create() error {
	return nil
}

func (service *ScoreService) UpdateById() error {
	return nil
}

func (service *ScoreService) DeleteById() error {
	return nil
}
