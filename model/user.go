package model

import (
	"errors"
	"tasked/authentication"
	"tasked/database"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Model
	Name            string     `gorm:"not null;size:200" json:"name"`
	Email           string     `gorm:"uniqueIndex;not null;size:200" json:"email"`
	Password        string     `json:"-"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	Tasks           *[]Task    `json:"tasks,omitempty"`
}

func (u *User) IsEmailUnique(email string) bool {
	err := database.ORM().Where("id != ?", u.ID).First(&User{}, "email = ?", email).Error

	return errors.Is(err, gorm.ErrRecordNotFound)
}

func (u *User) Authenticate(email string, password string) bool {
	err := database.ORM().First(&u, "email = ?", email).Error

	if err != nil {
		return false
	}

	return authentication.CheckPasswordHash(password, u.Password)
}
