package common

import (
	"time"
)

type Score struct {
	ID        uint64    `json:"-"`
	UserID    string    `json:"-" validate:"required,uuid"`
	IsWon     float64   `json:"is_won" validate:"required,boolean"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
