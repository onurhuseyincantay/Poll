package controllers

import (
	"api/database"
	"api/repository"
	"api/repository/crud"
	"api/responses"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetPolls(writer http.ResponseWriter, request *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewPollRepositoryCrud(db)

	func(pollRepository repository.PollRepository) {
		polls, err := pollRepository.GetPolls()
		if err != nil {
			responses.Error(writer, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(writer, http.StatusOK, polls)
	}(repo)
}

func GetPollsByUserId(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.ParseUint(request.URL.Query()["userId"][0],10,32)

	if err != nil {
		responses.Error(writer, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewPollRepositoryCrud(db)

	func(pollRepository repository.PollRepository) {
		polls, err := pollRepository.FindPollByUserID(uint(id))
		if err != nil {
			responses.Error(writer, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(writer, http.StatusOK, polls)
	}(repo)
}

func GetPollByID(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseUint(params["id"],10,32)

	if err != nil {
		print("gg")
		responses.Error(writer, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewPollRepositoryCrud(db)

	func(pollRepository repository.PollRepository) {
		polls, err := pollRepository.FindPollByID(uint(id))
		if err != nil {
			responses.Error(writer, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(writer, http.StatusOK, polls)
	}(repo)
}

