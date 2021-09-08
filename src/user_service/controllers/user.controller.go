package controllers

import (
	"User_Service/domain"
	"User_Service/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserService interface {
	AddUser(user domain.User) error
	LoginUser(user domain.User) error
}

type MessageBroker interface {
	SendInfo(msg string)
	SendWarning(msg string)
	SendError(msg string)
	Close()
}

var UserServ UserService
var MsgBroker MessageBroker

func Login(w http.ResponseWriter, r *http.Request) {
	MsgBroker.SendInfo(fmt.Sprintf("Request at \\user\\login"))
	user := &domain.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		MsgBroker.SendError(fmt.Sprintf("user_service: %s", "Invalid request parameters"))
		utils.Respond(w, utils.Message("Invalid request parameters"))
		return
	}

	err := UserServ.LoginUser(*user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		MsgBroker.SendError(fmt.Sprintf("user_service: %s", err.Error()))
		utils.Respond(w, utils.Message(err.Error()))
		return
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	user := &domain.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		MsgBroker.SendError(fmt.Sprintf("user_service: %s", "Invalid request parameters"))
		utils.Respond(w, utils.Message("Invalid request parameters"))
		return
	}

	err := UserServ.AddUser(*user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		MsgBroker.SendError(fmt.Sprintf("user_service: %s", err.Error()))
		utils.Respond(w, utils.Message(err.Error()))
		return
	}
}
