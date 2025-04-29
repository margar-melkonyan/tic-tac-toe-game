package service

import (
	"context"
	"log/slog"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/config"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/repository"
	"golang.org/x/crypto/bcrypt"
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

func (service *RoomService) Create(ctx context.Context, form common.RoomRequest) error {
	password, err := bcrypt.GenerateFromPassword([]byte(form.Password), config.ServerConfig.BcryptPower)
	if err != nil {
		return err
	}
	form.Password = string(password)
	room := common.Room{
		CreatorID: form.CreatorID,
		Name:      form.Name,
		Password:  form.Password,
		IsPrivate: form.IsPrivate,
	}
	return service.repo.Create(ctx, room)
}

func (service *RoomService) DeleteById(ctx context.Context, id uint64) error {
	return service.repo.DeleteById(ctx, id)
}
