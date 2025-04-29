package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
)

type RoomRepo struct {
	db *sql.DB
}

type RoomRepository interface {
	FindAll(ctx context.Context) ([]*common.Room, error)
	Create(ctx context.Context, room common.Room) error
	DeleteById(ctx context.Context, id uint64) error
}

func NewRoomRepository(db *sql.DB) RoomRepository {
	return &RoomRepo{
		db: db,
	}
}

func (repo *RoomRepo) FindAll(ctx context.Context) ([]*common.Room, error) {
	var rooms []*common.Room
	rows, err := repo.db.QueryContext(ctx, "SELECT * FROM rooms WHERE deleted_at IS NULL")
	defer func() {
		rows.Close()
	}()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var room common.Room
		err := rows.Scan(
			&room.ID,
			&room.Name,
			&room.IsPrivate,
			&room.Password,
			&room.CreatorID,
			&room.Capacity,
			&room.CreatedAt,
			&room.UpdatedAt,
			&room.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}
	if rooms == nil {
		rooms = []*common.Room{}
	}
	return rooms, nil
}

func (repo *RoomRepo) Create(ctx context.Context, room common.Room) error {
	query := "INSERT INTO rooms (name, is_private, creator_id, password) VALUES ($1, $2, $3, $4)"
	result, err := repo.db.ExecContext(
		ctx,
		query,
		room.Name,
		room.IsPrivate,
		room.CreatorID,
		room.Password,
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

func (repo *RoomRepo) DeleteById(ctx context.Context, id uint64) error {
	result, err := repo.db.ExecContext(
		ctx,
		"UPDATE rooms SET deleted_at = now() WHERE id = $1 AND deleted_at IS NULL",
		id,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("room was not deleted")
	}
	return nil
}
