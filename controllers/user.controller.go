package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/shaikdev/groshop-server/helpers"
	"github.com/shaikdev/groshop-server/models"
	"github.com/shaikdev/groshop-server/services"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	helpers.Header(w, "POST")
	// decode
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	isBodyCheck, err := user.UserBodyCheck()
	if isBodyCheck {
		helpers.ResponseErrorSender(w, err, helpers.FAILED, http.StatusBadRequest)
		return
	}
	createUser, createUserErr := services.CreateUser(user)
	if createUserErr != nil {
		helpers.ResponseErrorSender(w, helpers.USER_CREATE_FAILED, helpers.FAILED, http.StatusBadRequest)
		return
	}

	getUser, getUserErr := services.GetUserById(createUser.Hex())
	if getUserErr != nil {
		helpers.ResponseErrorSender(w, helpers.USER_FETCH_FAILED, helpers.FAILED, http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccess(w, helpers.USER_CREATE_SUCCESS, helpers.SUCCESS, http.StatusCreated, getUser, 0)

}
