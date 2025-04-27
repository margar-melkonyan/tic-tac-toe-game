package postgres_storage

import (
	"context"
	"database/sql"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
)

type RoomsRepo struct {
	db *sql.DB
}

type RoomsRepository interface {
	FindById(ctx context.Context, id uint64) (*common.Room, error)
	FindAll(ctx context.Context) ([]*common.Room, error)
	DeleteById(ctx context.Context, id uint64) error
	UpdateById(ctx context.Context, entity common.Room) error
}

func NewRoomRepository(db *sql.DB) *RoomsRepo {
	return &RoomsRepo{
		db: db,
	}
}

func (repo *RoomsRepo) FindById(ctx context.Context, id uint64) (*common.Room, error) {
	return nil, nil
}
func (repo *RoomsRepo) FindAll(ctx context.Context) ([]*common.Room, error)      { return nil, nil }
func (repo *RoomsRepo) DeleteById(ctx context.Context, id uint64) error          { return nil }
func (repo *RoomsRepo) UpdateById(ctx context.Context, entity common.Room) error { return nil }
