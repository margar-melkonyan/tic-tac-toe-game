package service

import (
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/repository"
)

type UserService struct {
	repo repository.RoomRepository
}

func NewUserService(repoRoom repository.RoomRepository) *UserService {
	return &UserService{
		repo: repoRoom,
	}
}

func (service *UserService) GetAll() []*common.Room {
	return nil
}

func (service *UserService) GetById() *common.Room {
	return nil
}

func (service *UserService) Create() error {
	return nil
}

func (service *UserService) UpdateById() error {
	return nil
}

func (service *UserService) DeleteById() error {
	return nil
}
