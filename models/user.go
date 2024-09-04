package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
