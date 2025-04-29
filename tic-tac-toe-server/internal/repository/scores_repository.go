package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
)

const TABLE_NAME = "scores"

type ScoreRepo struct {
	db *sql.DB
}

type ScoreRepository interface {
	Create(ctx context.Context, score *common.Score) error
	FindAllByUser(ctx context.Context, user *common.User) ([]*common.Score, error)
	GetWonScore(ctx context.Context, user *common.User) (uint, error)
}

func NewScoreRepository(db *sql.DB) ScoreRepository {
	return &ScoreRepo{
		db: db,
	}
}

func (repo ScoreRepo) Create(ctx context.Context, score *common.Score) error {
	return nil
}

func (repo ScoreRepo) FindAllByUser(ctx context.Context, user *common.User) ([]*common.Score, error) {
	var scores []*common.Score
	query := fmt.Sprintf(
		"SELECT id, user_id, is_won, created_at FROM %v WHERE user_id = $1 AND deleted_at IS NULL",
		TABLE_NAME,
	)
	rows, err := repo.db.QueryContext(ctx, query, user.ID)
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
	if scores == nil {
		scores = []*common.Score{}
	}
	return scores, nil
}

func (repo ScoreRepo) GetWonScore(ctx context.Context, user *common.User) (uint, error) {
	var currentScore *uint
	query := fmt.Sprintf(
		"SELECT COUNT(*) as current_score FROM %v WHERE user_id = $1 AND is_won = true",
		TABLE_NAME,
	)
	row := repo.db.QueryRowContext(ctx, query, user.ID)
	err := row.Scan(
		&currentScore,
	)
	if err != nil {
		return 0, err
	}
	return *currentScore, nil
}
