package service

import (
	"context"
	"log/slog"

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

func (service *RoomService) GetAll(ctx context.Context) []*common.Room {
	rooms, err := service.repo.FindAll(context.Background())
	if err != nil {
		slog.Error(err.Error())
		return []*common.Room{}
	}
	return rooms
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
