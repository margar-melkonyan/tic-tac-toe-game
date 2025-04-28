package common

import (
	"time"

	"github.com/google/uuid"
)

type Score struct {
	ID        uint64    `json:"-"`
	UserID    uuid.UUID `json:"user_id" validate:"required,uuid"`
	IsWon     bool      `json:"is_won" validate:"required,boolean"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
