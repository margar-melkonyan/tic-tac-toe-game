package repository

import (
	"context"
	"database/sql"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
)

type RoomRepo struct {
	db *sql.DB
}

type RoomRepository interface {
	FindById(ctx context.Context, id uint64) (*common.Room, error)
	FindAll(ctx context.Context) ([]*common.Room, error)
	DeleteById(ctx context.Context, id uint64) error
	UpdateById(ctx context.Context, entity common.Room) error
}

func NewRoomRepository(db *sql.DB) RoomRepository {
	return &RoomRepo{
		db: db,
	}
}

func (repo *RoomRepo) FindById(ctx context.Context, id uint64) (*common.Room, error) {
	return nil, nil
}
func (repo *RoomRepo) FindAll(ctx context.Context) ([]*common.Room, error)      { return nil, nil }
func (repo *RoomRepo) DeleteById(ctx context.Context, id uint64) error          { return nil }
func (repo *RoomRepo) UpdateById(ctx context.Context, entity common.Room) error { return nil }
