package service

import (
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/repository"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repoRoom repository.UserRepository) *AuthService {
	return &AuthService{
		repo: repoRoom,
	}
}

func (service *AuthService) GetAll() []*common.Room {
	return nil
}

func (service *AuthService) GetById() *common.Room {
	return nil
}

func (service *AuthService) Create() error {
	return nil
}

func (service *AuthService) UpdateById() error {
	return nil
}

func (service *AuthService) DeleteById() error {
	return nil
}
