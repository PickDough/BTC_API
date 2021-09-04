package controllers

import (
	"User_Service/domain"
	"User_Service/utils"
	"encoding/json"
	"net/http"
)

type UserService interface {
	AddUser(user domain.User) error
	LoginUser(user domain.User) error
}

var UserServ UserService

func Login(w http.ResponseWriter, r *http.Request) {
	user := &domain.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, utils.Message("Invalid request parameters"))
		return
	}

	err := UserServ.LoginUser(*user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, utils.Message(err.Error()))
		return
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	user := &domain.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, utils.Message("Invalid request parameters"))
		return
	}

	err := UserServ.AddUser(*user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.Respond(w, utils.Message(err.Error()))
		return
	}
}
