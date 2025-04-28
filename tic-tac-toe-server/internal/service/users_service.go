package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/repository"
)

type UserService struct {
	userRepo  repository.UserRepository
	scoreRepo repository.ScoreRepository
}

func NewUserService(userRepo repository.UserRepository, scoreRepo repository.ScoreRepository) *UserService {
	return &UserService{
		userRepo:  userRepo,
		scoreRepo: scoreRepo,
	}
}

func (service *UserService) GetCurrentUser(ctx context.Context) (*common.UserResponse, error) {
	email, ok := ctx.Value("user_email").(string)
	if !ok {
		return nil, errors.New("user email is not valid")
	}
	user, err := service.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	scores, err := service.scoreRepo.GetAllByUser(ctx, user)
	if err != nil {
		return nil, err
	}
	fmt.Println(scores)
	userResponse := &common.UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Scores: scores,
	}
	return userResponse, err
}
