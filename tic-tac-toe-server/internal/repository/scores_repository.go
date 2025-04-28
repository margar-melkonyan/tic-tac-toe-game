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
	Create(ctx context.Context, score *common.Score) error
	GetAllByUser(ctx context.Context, user *common.User) ([]*common.Score, error)
}

func NewScoreRepository(db *sql.DB) ScoreRepository {
	return &ScoreRepo{
		db: db,
	}
}

func (repo ScoreRepo) Create(ctx context.Context, score *common.Score) error {
	return nil
}

func (repo ScoreRepo) GetAllByUser(ctx context.Context, user *common.User) ([]*common.Score, error) {
	var scores []*common.Score
	query := "SELECT user_id, is_won, created_at FROM scores WHERE user_id = $1"
	rows, err := repo.db.QueryContext(ctx, query)
	defer func() {
		rows.Close()
	}()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var score common.Score
		err := rows.Scan(
			&score.ID,
			&score.UserID,
			&score.IsWon,
			&score.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		scores = append(scores, &score)
	}
	return scores, nil
}
