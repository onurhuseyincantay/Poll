package crud

import (
	"api/models"
	"api/utils/channels"
	"github.com/jinzhu/gorm"
)

type PollRepositoryCrud struct {
	db *gorm.DB
}

func NewPollRepositoryCrud(db *gorm.DB) *PollRepositoryCrud {
	return &PollRepositoryCrud{db}
}

func (r *PollRepositoryCrud) GetPolls() ([]models.Poll, error) {
	var err error
	var polls []models.Poll
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Find(&polls).Error
		if err != nil {
			ch <- false
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return polls, nil
	}
	return polls, err
}

func (r *PollRepositoryCrud) FindPollByID(pollID uint) (models.Poll, error) {
	var err error
	poll := models.Poll{ID:pollID}
	var answers []models.Answer
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		r.db.Preload("answers")
		err = r.db.Debug().Model(&poll).Related(&answers).Where("id = ?", pollID).Find(&poll).Error
		if err != nil {
			ch <- false
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		poll.Answers = answers
		return poll, nil
	}
	return models.Poll{}, err
}

func (r *PollRepositoryCrud) FindPollByUserID(userID uint) ([]models.Poll, error) {
	var err error
	var polls []models.Poll
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&polls).Where("created_user_id = ?",userID).Find(&polls).Error
		if err != nil {
			ch <- false
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return polls, nil
	}
	return []models.Poll{}, err
}