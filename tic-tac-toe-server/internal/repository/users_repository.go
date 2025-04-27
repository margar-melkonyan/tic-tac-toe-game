package repository

import (
	"context"
	"database/sql"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
)

type UserRepo struct {
	db *sql.DB
}

type UserRepository interface {
	FindById(ctx context.Context, id uint64) (*common.Room, error)
	FindAll(ctx context.Context) ([]*common.Room, error)
	DeleteById(ctx context.Context, id uint64) error
	UpdateById(ctx context.Context, entity common.Room) error
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepo{
		db: db,
	}
}

func (repo *UserRepo) FindById(ctx context.Context, id uint64) (*common.Room, error) {
	return nil, nil
}

func (repo *UserRepo) FindAll(ctx context.Context) ([]*common.Room, error) { return nil, nil }

func (repo *UserRepo) DeleteById(ctx context.Context, id uint64) error { return nil }

func (repo *UserRepo) UpdateById(ctx context.Context, entity common.Room) error { return nil }
