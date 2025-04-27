package service

import (
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/repository"
)

type RoomService struct {
	repo repository.RoomRepository
}

func NewRoomService(repoRoom repository.RoomRepository) *RoomService {
	return &RoomService{
		repo: repoRoom,
	}
}

func (service *RoomService) GetAll() []*common.Room {
	return nil
}

func (service *RoomService) GetById() *common.Room {
	return nil
}

func (service *RoomService) Create() error {
	return nil
}

func (service *RoomService) UpdateById() error {
	return nil
}

func (service *RoomService) DeleteById() error {
	return nil
}
