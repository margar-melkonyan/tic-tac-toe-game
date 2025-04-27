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

func (repo *RoomRepo) FindAll(ctx context.Context) ([]*common.Room, error) {
	var rooms []*common.Room
	rows, err := repo.db.Query("SELECT * FROM rooms;")
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
	return rooms, nil
}

func (repo *RoomRepo) DeleteById(ctx context.Context, id uint64) error { return nil }

func (repo *RoomRepo) UpdateById(ctx context.Context, entity common.Room) error { return nil }
