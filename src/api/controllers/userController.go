package controllers

import (
	"api/database"
	"api/models"
	"api/repository"
	"api/repository/crud"
	"api/responses"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetUsers(writer http.ResponseWriter, request *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewUserRepositoryCrud(db)

	func(userRepository repository.UserRepository) {
		users, err := userRepository.FindAll()
		if err != nil {
			responses.Error(writer, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(writer, http.StatusOK, users)
	}(repo)
}

func GetUserById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseUint(params["id"],10,32)

	if err != nil {
		responses.Error(writer, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewUserRepositoryCrud(db)

	func(userRepository repository.UserRepository) {
		users, err := userRepository.FindById(uint32(id))
		if err != nil {
			responses.Error(writer, http.StatusBadRequest, err)
			return
		}
		responses.JSON(writer, http.StatusOK, users)
	}(repo)
}

func RegisterUser(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		responses.Error(writer, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.Error(writer, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewUserRepositoryCrud(db)

	func(userRepository repository.UserRepository) {
		user, err = userRepository.Save(user)
		if err != nil {
			responses.Error(writer, http.StatusUnprocessableEntity, err)
			return
		}
		writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", request.Host, request.RequestURI, user.ID))
		responses.JSON(writer, http.StatusCreated, user)
	}(repo)
}

func LoginUser(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		responses.Error(writer, http.StatusBadRequest, err)
		return
	}

	email := request.PostForm.Get("email")
	password := request.PostForm.Get("password")
	if email == "" || password == "" {
		responses.Error(writer, http.StatusBadRequest, errors.New("Empty body requested"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewUserRepositoryCrud(db)

	func(userRepository repository.UserRepository) {
		user, err := userRepository.Login(email, password)
		if err != nil {
			responses.Error(writer, http.StatusBadRequest, err)
			return
		}
		responses.JSON(writer, http.StatusOK, user)
	}(repo)
}

func UpdateProfile(writer http.ResponseWriter,request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseUint(params["id"],10,32)

	if err != nil {
		responses.Error(writer, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		responses.Error(writer, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.Error(writer, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewUserRepositoryCrud(db)

	func(userRepository repository.UserRepository) {
		rows, err := userRepository.Update(uint32(id), user)
		if err != nil {
			responses.Error(writer, http.StatusBadRequest, err)
			return
		}
		responses.JSON(writer, http.StatusOK, rows)
	}(repo)
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseUint(params["id"],10,32)

	if err != nil {
		responses.Error(writer, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewUserRepositoryCrud(db)

	func(userRepository repository.UserRepository) {
		users, err := userRepository.Delete(uint32(id))
		if err != nil {
			responses.Error(writer, http.StatusBadRequest, err)
			return
		}
		writer.Header().Set("Entity", fmt.Sprint("%d",id))
		responses.JSON(writer, http.StatusNoContent, users)
	}(repo)
}
