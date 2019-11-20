package auto

import (
	"api/database"
	"api/models"
	"api/utils/console"
	"log"
)

func Load() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Debug().DropTableIfExists(&models.User{}, &models.Poll{}, &models.Answer{}).Error
	if err != nil {
		log.Fatal(err)
	}



	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().AutoMigrate(&models.Poll{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().AutoMigrate(&models.Answer{}).Error
	if err != nil {
		log.Fatal(err)
	}

	db.Model(&models.Poll{}).AddForeignKey("created_user_id","users(id)","CASCADE","CASCADE")
	db.Model(&models.Answer{}).AddForeignKey("answered_user_id","users(id)","CASCADE","CASCADE").AddForeignKey("poll_id","polls(id)","RESTRICT","RESTRICT")



	for _, user := range users {
		err = db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			log.Fatal(err)
		}
		console.Pretty(user)
	}

	for _, poll := range polls {
		err = db.Debug().Model(&models.Poll{}).Create(&poll).Error
		if err != nil {
			log.Fatal(err)
		}
		console.Pretty(poll)
	}

	/*for _, answer := range answers {
		err = db.Debug().Model(&models.Answer{}).Create(&answer).Error
		if err != nil {
			log.Fatal(err)
		}
		console.Pretty(answer)
	}*/





}
