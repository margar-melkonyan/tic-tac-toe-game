package service

import (
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repoUser repository.UserRepository) *UserService {
	return &UserService{
		repo: repoUser,
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
