package repository

import (
	"context"
	"database/sql"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
)

type ScoreRepo struct {
	db *sql.DB
}

type ScoreRepository interface {
	FindById(ctx context.Context, id uint64) (*common.Room, error)
	FindAll(ctx context.Context) ([]*common.Room, error)
	DeleteById(ctx context.Context, id uint64) error
	UpdateById(ctx context.Context, entity common.Room) error
}

func NewScoreRepository(db *sql.DB) ScoreRepository {
	return &ScoreRepo{
		db: db,
	}
}

func (repo *ScoreRepo) FindById(ctx context.Context, id uint64) (*common.Room, error) {
	return nil, nil
}

func (repo *ScoreRepo) FindAll(ctx context.Context) ([]*common.Room, error) { return nil, nil }

func (repo *ScoreRepo) DeleteById(ctx context.Context, id uint64) error { return nil }

func (repo *ScoreRepo) UpdateById(ctx context.Context, entity common.Room) error { return nil }
