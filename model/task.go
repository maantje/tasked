package model

import (
	"time"
)

type Task struct {
	Model
	Name      string    `json:"name"`
	ExpiresAt time.Time `json:"expires_at"`
	UserId    int       `json:"user_id" gorm:"not null"`
	User      *User     `json:"user,omitempty"`
}
