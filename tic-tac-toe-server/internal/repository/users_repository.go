package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
)

type UserRepo struct {
	db *sql.DB
}

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*common.User, error)
	Create(ctx context.Context, form common.AuthSignUpRequest) error
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepo{
		db: db,
	}
}

func (repo *UserRepo) FindByEmail(ctx context.Context, email string) (*common.User, error) {
	var user common.User
	query := "SELECT id, name, email, password, created_at FROM users WHERE email = $1 AND deleted_at IS NULL"
	row := repo.db.QueryRowContext(ctx, query, email)
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepo) Create(ctx context.Context, form common.AuthSignUpRequest) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"
	result, err := repo.db.ExecContext(
		ctx,
		query,
		&form.Name,
		&form.Email,
		&form.Password,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("room was not created")
	}
	return nil
}
