package models

import "time"

type Poll struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedUserID uint `gorm:"not null" json:"createdUser"`
	Content string `gorm: "type:text" json: "content"`
	Choices string  `gorm:"varchar(255)" json:"choices"`
	Answers []Answer `gorm:"foreignkey:ID;not null;auto_preload" json:"answers"`
}