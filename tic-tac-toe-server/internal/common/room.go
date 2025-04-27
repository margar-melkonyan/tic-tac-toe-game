package common

import "time"

type Room struct {
	Id        uint64
	Name      string
	IsPrivate bool
	Password  string
	Capacity  uint8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type RoomRequest struct {
	Name      string `json:"name"`
	IsPrivate bool   `json:"is_private,omitempty"`
	Password  string `json:"password"`
}
