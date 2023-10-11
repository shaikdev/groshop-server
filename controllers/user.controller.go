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
	defer r.Body.Close()
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
	defer r.Body.Close()

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
	comparePasswordErr := services.ComparePasswords(getUser.Password, user.Password)
	if comparePasswordErr != nil {
		helpers.ResponseErrorSender(w, helpers.PASSWORD_DOES_NOT_MATCH, helpers.FAILED, http.StatusBadRequest)
		return
	}
	token, tokenErr := services.GenerateJwtToken(getUser)
	if tokenErr != nil {
		helpers.ResponseErrorSender(w, helpers.FAILED_TOKEN_CREATION, helpers.FAILED, http.StatusBadRequest)
		return
	}
	helpers.ResponseSuccess(w, helpers.USER_LOGIN_SUCCESS, helpers.SUCCESS, http.StatusOK, map[string]interface{}{"data": getUser, "token": "Bearer " + token})
}

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helpers.Header(w, "POST")
		// get Header from request
		var authorization = r.Header.Get("Authorization")
		if authorization == "" {
			helpers.ResponseErrorSender(w, helpers.TOKEN_NOT_FOUND, helpers.FAILED, http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authorization, "Bearer ") {
			helpers.ResponseErrorSender(w, helpers.INVALID_TOKEN, helpers.FAILED, http.StatusUnauthorized)
			return
		}

		token := strings.Split(authorization, " ")[1]

		accessToken, validErr := services.VerifyToken(token)
		if validErr != nil {
			helpers.ResponseErrorSender(w, validErr.Error(), helpers.FAILED, http.StatusUnauthorized)
			return
		}

		r.Header.Set("id", accessToken["_id"].(string))
		next.ServeHTTP(w, r)
	})

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("id")

	defer r.Body.Close()
	if user, err := services.GetUser(userId, ""); err != nil {
		helpers.ResponseErrorSender(w, helpers.USER_FETCH_FAILED, helpers.FAILED, http.StatusNotFound)
		return
	} else {
		helpers.ResponseSuccess(w, helpers.USER_FETCH_SUCCESS, helpers.SUCCESS, http.StatusOK, map[string]interface{}{"data": user})
		return
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	helpers.Header(w, "GET")
	defer r.Body.Close()
	users, err := services.GetUsers()
	if err != nil {
		helpers.ResponseErrorSender(w, helpers.USER_GET_FAILED, helpers.FAILED, http.StatusInternalServerError)
		return
	}
	helpers.ResponseSuccess(w, helpers.USER_GET_SUCCESSFULLY, helpers.SUCCESS, http.StatusOK, map[string]interface{}{"data": users})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	helpers.Header(w, "PUT")
	userId := r.Header.Get("id")
	defer r.Body.Close()
	// check user is exist or not
	if _, isUserErr := services.GetUser(userId, ""); isUserErr != nil {
		helpers.ResponseErrorSender(w, helpers.USER_FETCH_FAILED, helpers.FAILED, http.StatusNotFound)
		return
	}
	// decode the body
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	if bodyCheck, bodyCheckErr := user.EditUser(); bodyCheck {
		helpers.ResponseErrorSender(w, bodyCheckErr, helpers.FAILED, http.StatusBadRequest)
		return
	}
	if isUpdate, err := services.UpdateUser(userId, user); err != nil {
		helpers.ResponseErrorSender(w, helpers.USER_EDIT_FAILED, helpers.FAILED, http.StatusInternalServerError)
		return
	} else if !isUpdate {
		helpers.ResponseErrorSender(w, helpers.USER_EDIT_FAILED, helpers.FAILED, http.StatusBadRequest)
		return
	} else {
		if getUser, getUserErr := services.GetUser(userId, ""); getUserErr != nil {
			helpers.ResponseErrorSender(w, helpers.USER_FETCH_FAILED, helpers.FAILED, http.StatusNotFound)
			return
		} else {
			helpers.ResponseSuccess(w, helpers.USER_EDIT_SUCCESSFULLY, helpers.SUCCESS, http.StatusOK, map[string]interface{}{"data": getUser})
		}
	}
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	helpers.Header(w, "DELETE")
	userId := r.Header.Get("id")
	defer r.Body.Close()
	// check if user exists or not
	if _, err := services.GetUser(userId, ""); err != nil {
		helpers.ResponseErrorSender(w, helpers.USER_GET_FAILED, helpers.FAILED, http.StatusNotFound)
		return
	} else {
		if deleteUser, deleteFailed := services.DeleteUser(userId); deleteFailed != nil {
			helpers.ResponseErrorSender(w, helpers.USER_DELETE_FAILED, helpers.FAILED, http.StatusInternalServerError)
			return
		} else if deleteUser == 0 {
			helpers.ResponseErrorSender(w, helpers.USER_DELETE_FAILED, helpers.FAILED, http.StatusBadRequest)
			return
		} else {
			helpers.ResponseSuccess(w, helpers.USER_DELETED_SUCCESSFULLY, helpers.SUCCESS, http.StatusOK, map[string]interface{}{})
		}
	}
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	helpers.Header(w, "DELETE")
	defer r.Body.Close()

	if delete, err := services.DeleteAllUser(); err != nil {
		helpers.ResponseErrorSender(w, helpers.USERS_DELETE_FAILED, helpers.FAILED, http.StatusInternalServerError)
		return
	} else if delete == 0 {
		helpers.ResponseErrorSender(w, helpers.USERS_DELETE_FAILED, helpers.FAILED, http.StatusUnprocessableEntity)
		return
	} else {
		helpers.ResponseSuccess(w, helpers.USERS_DELETED_SUCCESSFULLY, helpers.SUCCESS, http.StatusOK, map[string]interface{}{})
	}

}
