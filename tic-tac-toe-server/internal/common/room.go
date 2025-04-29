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
	CreatorID uuid.UUID `validate:"required,uuid" json:"creator_id"`
	Name      string    `validate:"required,alphanumunicode,min=8,max=255" json:"name"`
	IsPrivate bool      `validate:"boolean" json:"is_private"`
	Password  string    `validate:"required,alpha,min=8,max=255" json:"password"`
}
