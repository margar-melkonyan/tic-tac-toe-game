package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/config"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type RoomService struct {
	repo repository.RoomRepository
}

const op = "room_serivce"

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
		playerIn := 0
		roomInfo := GetCurrentRoomInfo(room.ID)
		if roomInfo != nil {
			playerIn = len(roomInfo.Users)
		}
		roomsResponse = append(roomsResponse, &common.RoomResponse{
			ID:        room.ID,
			Name:      room.Name,
			Capacity:  room.Capacity,
			IsPrivate: &room.IsPrivate,
			PlayerIn:  playerIn,
		})
	}
	if err != nil {
		slog.Error(
			fmt.Sprintf("[%v:GetAll]", op),
			slog.String("error", err.Error()),
		)
		return []*common.RoomResponse{}
	}
	return roomsResponse
}

func (service *RoomService) GetById(ctx context.Context, id uint64) (*common.RoomSessionResponse, error) {
	room, err := service.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	users := make([]*common.UserResponse, 0)
	roomSession := GetCurrentRoomInfo(room.ID)
	if roomSession != nil {
		for _, user := range roomSession.Users {
			users = append(users, &common.UserResponse{
				ID:   user.ID,
				Name: user.Name,
			})
		}
	}
	resp := &common.RoomSessionResponse{
		ID:        room.ID,
		Name:      room.Name,
		CreatorID: room.CreatorID,
		IsPrivate: &room.IsPrivate,
		Capacity:  room.Capacity,
		Users:     users,
	}
	return resp, err
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
