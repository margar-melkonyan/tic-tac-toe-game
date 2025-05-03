package common

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID        uint64     `json:"id"`
	Name      string     `json:"name"`
	IsPrivate bool       `json:"is_private"`
	CreatorID uuid.UUID  `json:"creator_id"`
	Password  string     `json:"-"`
	Capacity  uint8      `json:"capacity"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type RoomRequest struct {
	CreatorID uuid.UUID `json:"creator_id"`
	Name      string    `validate:"required,min=8,max=255" json:"name"`
	IsPrivate *bool     `validate:"required,boolean" json:"is_private"`
	Password  *string   `validate:"required_if=IsPrivate true,max=255" json:"password"`
}

type RoomResponse struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	IsPrivate *bool  `json:"is_private"`
	Capacity  uint8  `json:"capacity"`
	PlayerIn  int    `json:"player_in"`
}

type RoomSessionResponse struct {
	ID        uint64          `json:"id"`
	Name      string          `json:"name"`
	CreatorID uuid.UUID       `json:"creator_id"`
	IsPrivate *bool           `json:"is_private,omitempty"`
	Capacity  uint8           `json:"capacity"`
	Users     []*UserResponse `json:"users"`
}
