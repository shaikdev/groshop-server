package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/shaikdev/groshop-server/helpers"
	"github.com/shaikdev/groshop-server/models"
	"github.com/shaikdev/groshop-server/services"
)

func UserRegister(w http.ResponseWriter, r *http.Request) {
	helpers.Header(w, "POST")
	// decode
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	user.Email = strings.ToLower(user.Email)
	isBodyCheck, err := user.UserBodyCheck()
	if isBodyCheck {
		helpers.ResponseErrorSender(w, err, helpers.FAILED, http.StatusBadRequest)
		return
	}
	existUser, _ := services.GetUser("", user.Email)
	if existUser.Email == user.Email {
		helpers.ResponseErrorSender(w, helpers.EMAIL_ALREADY_EXIST, helpers.FAILED, http.StatusBadRequest)
		return
	}
	passwordHash, hasError := services.HashPassword(user.Password)
	if hasError != nil {
		helpers.ResponseErrorSender(w, helpers.PASSWORD_HASH_FAILED, helpers.FAILED, http.StatusBadRequest)
		return
	}
	user.Password = passwordHash
	createUser, createUserErr := services.CreateUser(user)
	if createUserErr != nil {
		helpers.ResponseErrorSender(w, helpers.USER_CREATE_FAILED, helpers.FAILED, http.StatusBadRequest)
		return
	}

	getUser, getUserErr := services.GetUser(createUser.Hex(), "")
	if getUserErr != nil {
		helpers.ResponseErrorSender(w, helpers.USER_FETCH_FAILED, helpers.FAILED, http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccess(w, helpers.USER_CREATE_SUCCESS, helpers.SUCCESS, http.StatusCreated, map[string]interface{}{"data": getUser})

}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	helpers.Header(w, "POST")

	// decode
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	user.Email = strings.ToLower(user.Email)
	isBodyCheck, err := user.LoginBodyCheck()
	if isBodyCheck {
		helpers.ResponseErrorSender(w, err, helpers.FAILED, http.StatusBadRequest)
		return
	}
	//TODO: check user exist or not
	getUser, getUserErr := services.GetUser("", user.Email)
	if getUserErr != nil {
		helpers.ResponseErrorSender(w, helpers.USER_FETCH_FAILED, helpers.FAILED, http.StatusNotFound)
		return
	}
	token, tokenErr := services.GenerateJwtToken(getUser)
	if tokenErr != nil {
		helpers.ResponseErrorSender(w, helpers.FAILED_TOKEN_CREATION, helpers.FAILED, http.StatusBadRequest)
		return
	}
	helpers.ResponseSuccess(w, helpers.USER_LOGIN_SUCCESS, helpers.SUCCESS, http.StatusOK, map[string]interface{}{"data": getUser, "token": token})

}
