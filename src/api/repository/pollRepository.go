package repository

import "api/models"

type PollRepository interface {
	FindPollByID(uint) (models.Poll, error)
	FindPollByUserID(uint) ([]models.Poll, error)
	GetPolls() ([]models.Poll, error)
}