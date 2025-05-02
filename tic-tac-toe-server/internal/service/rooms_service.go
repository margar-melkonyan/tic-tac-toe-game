package service

import (
	"context"
	"errors"
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

func (service *RoomService) GetAll(ctx context.Context) []*common.RoomResponse {
	rooms, err := service.repo.FindAll(context.Background())
	var roomsResponse []*common.RoomResponse
	roomsResponse = make([]*common.RoomResponse, 0)

	for _, room := range rooms {
		roomsResponse = append(roomsResponse, &common.RoomResponse{
			ID:        room.ID,
			Name:      room.Name,
			Capacity:  room.Capacity,
			IsPrivate: &room.IsPrivate,
		})
	}
	if err != nil {
		slog.Error(err.Error())
		return []*common.RoomResponse{}
	}
	return roomsResponse
}

func (service *RoomService) GetById(ctx context.Context, id uint64) (*common.Room, error) {
	return service.repo.FindById(ctx, id)
}

func (service *RoomService) Create(ctx context.Context, form common.RoomRequest) error {
	password, err := bcrypt.GenerateFromPassword([]byte(*form.Password), config.ServerConfig.BcryptPower)
	if err != nil {
		return err
	}
	hashedPassword := string(password)
	form.Password = &hashedPassword
	user, ok := ctx.Value(common.USER).(*common.User)
	if !ok {
		return errors.New("userId is not correct")
	}
	room := common.Room{
		CreatorID: user.ID,
		Name:      form.Name,
		Password:  *form.Password,
		IsPrivate: *form.IsPrivate,
	}
	return service.repo.Create(ctx, room)
}

func (service *RoomService) DeleteById(ctx context.Context, id uint64) error {
	return service.repo.DeleteById(ctx, id)
}
