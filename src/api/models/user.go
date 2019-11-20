package models

import (
	"api/security"
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Nickname  string    `gorm:"size:30;not null; unique" json:"nickname"`
	Email     string    `gorm:"not null; unique" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
}

func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
