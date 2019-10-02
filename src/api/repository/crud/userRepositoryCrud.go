package crud

import (
	"api/models"
	"api/security"
	"api/utils/channels"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type userRepositoryCrud struct {
	db *gorm.DB
}

// this part needs to be explained :/
func NewUserRepositoryCrud(db *gorm.DB) *userRepositoryCrud {
	return &userRepositoryCrud{db}
}

func (r *userRepositoryCrud) Save(user models.User) (models.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return user, nil
	}
	return models.User{}, err
}

func (r *userRepositoryCrud) FindAll() ([]models.User, error) {
	var err error
	users := []models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Limit(100).Find(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return users, nil
	}
	return nil, err
}

func (r *userRepositoryCrud) FindById(id uint32) (models.User, error) {
	var err error
	user := models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Where("id = ?",id).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return user, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return user, errors.New("User not found")
	}
	return user, err
}

func (r *userRepositoryCrud) Update(id uint32, user models.User) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?",id).Take(&models.User{}).Update(
			 map[string]interface{} {
			 	"nickname": user.Nickname,
			 	"email": user.Email,
			 	"updatedAt": time.Now(),
			 },
		)
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}

func (r *userRepositoryCrud) Delete(id uint32) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?",id).Take(&models.User{}).Update(&models.User{})
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}

func (r *userRepositoryCrud) Login(email string, password string) (models.User, error) {
	var err error
	user := models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		value, err := security.Hash(password)
		if err != nil {
			ch <- false
			return
		}
		err = r.db.Debug().Model(&models.User{}).Where("email = ? AND password = ?", email, string(value)).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return user, errors.New("User not found")
	}
	return user, err
}
