package service

import (
	"context"
	"errors"

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

func (service *UserService) GetCurrentUser(ctx context.Context) (*common.User, error) {
	email, ok := ctx.Value("user_email").(string)
	if !ok {
		return nil, errors.New("user email is not valid")
	}
	return service.repo.FindByEmail(ctx, email)
}
