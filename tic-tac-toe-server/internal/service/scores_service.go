package service

import (
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/repository"
)

type ScoreService struct {
	repo repository.RoomRepository
}

func NewScoreService(repoRoom repository.ScoreRepository) *ScoreService {
	return &ScoreService{
		repo: repoRoom,
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
