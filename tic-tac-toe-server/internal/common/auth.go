package common

import (
	"time"

	"github.com/google/uuid"
)

const USER_MAIL = "user_mail"
const USER = "user"

type AuthSignInRequest struct {
	Email    string `json:"email" validate:"required,min=4,max=255"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}

type AuthSignUpRequest struct {
	Name                 string `json:"name" validate:"required,min=4,max=255"`
	Email                string `json:"email" validate:"required,email,min=8,max=255"`
	Password             string `json:"password" validate:"required,min=8,max=255,eqfield=PasswordConfirmation"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,min=8,max=255"`
}

type User struct {
	ID        uuid.UUID  `json:"-"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type UserResponse struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email,omitempty"`
	WonScore  *uint      `json:"current_won_score,omitempty"`
	Symbol    string     `json:"symbol,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}
