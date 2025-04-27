package common

import "time"

type Room struct {
	ID        uint64     `json:"id"`
	Name      string     `json:"name"`
	IsPrivate bool       `json:"is_private"`
	Password  string     `json:"-"`
	Capacity  uint8      `json:"capacity"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type RoomRequest struct {
	Name      string `validate:"required,alphanumunicode,min=8,max=255" json:"name"`
	IsPrivate bool   `validate:"required,boolean" json:"is_private"`
	Password  string `validate:"required,alpha,min=8,max=255" json:"password"`
}
