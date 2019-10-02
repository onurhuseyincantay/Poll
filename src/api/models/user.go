package models

import (
	"api/security"
	"time"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:30;not null; unique" json:"nickname"`
	Email     string    `gorm:"not null; unique" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `gorm:"default.current_timestamp()" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default.current_timestamp()" json:"updatedAt"`
}

func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
