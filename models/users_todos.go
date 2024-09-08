package models

import (
	"github.com/google/uuid"
	"time"
)

type UsersTodos struct {
	UserID    uuid.UUID `json:"user_id"`
	TodoID    uuid.UUID `json:"todo_id"`
	CreatedAt time.Time `json:"created_at"`
}
