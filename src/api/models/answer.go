package models

import (
	"time"
)

type Answer struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	AnsweredUserID uint `gorm:" not null" json:"answeredUserId"`
	PollID uint `gorm:" not null" json:"pollId"`
	Content string 	`gorm:"not null" json:"content"`
}