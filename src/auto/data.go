package auto

import "api/models"

var users = []models.User{
	models.User{Nickname: "OnurCnty", Email: "ohc3807@gmail.com", Password: "123123."},

	models.User{Nickname: "Ege", Email: "ege@gmail.com", Password: "123123."},
}

var polls = []models.Poll{
	models.Poll{CreatedUserID:1, Content:"Should I take pills?", Choices:"yes", Answers: answers},
}

var answers = []models.Answer {
	models.Answer{Content:"Should I take pills?", PollID:1, AnsweredUserID: 2},
}
