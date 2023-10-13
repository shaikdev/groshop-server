package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shaikdev/groshop-server/helpers"
	addressresponse "github.com/shaikdev/groshop-server/helpers/address_response"
	userresponse "github.com/shaikdev/groshop-server/helpers/user_response"
	"github.com/shaikdev/groshop-server/models"
	"github.com/shaikdev/groshop-server/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAddress(w http.ResponseWriter, r *http.Request) {
	helpers.Header(w, "PUT")
	userId := r.Header.Get("id")
	defer r.Body.Close()

	//TODO: check user exist or not
	user, err := services.GetUser(userId, "")
	if err != nil {
		helpers.ResponseErrorSender(w, userresponse.USER_GET_FAILED, helpers.FAILED, http.StatusNotFound)
		return
	}
	// decode address
	var address models.Address
	json.NewDecoder(r.Body).Decode(&address)
	if addressBodyCheck, addressBodyCheckErr := address.AddressBodyCheck(); addressBodyCheck {
		helpers.ResponseErrorSender(w, addressBodyCheckErr, helpers.FAILED, http.StatusBadRequest)
		return
	}
	// TODO: Add _id in address
	address.Id = primitive.NewObjectID()

	// TODO: Removed previous addresses for avoiding multiple addresses updated
	user.Address = []*models.Address{}
	user.Address = append(user.Address, &address)
	_, createAddressErr := services.CreateAddress(userId, user)
	if createAddressErr != nil {
		helpers.ResponseErrorSender(w, addressresponse.CREATE_ADDRESS_FAILED, helpers.FAILED, http.StatusInternalServerError)
		return
	}
	getUpdateUser, updateUserErr := services.GetUser(userId, "")
	if updateUserErr != nil {
		helpers.ResponseErrorSender(w, userresponse.USER_GET_FAILED, helpers.FAILED, http.StatusInternalServerError)
		return
	}
	helpers.ResponseSuccess(w, addressresponse.CREATE_ADDRESS_FAILED, helpers.SUCCESS, http.StatusCreated, map[string]interface{}{"data": getUpdateUser})
}

func UpdateAddress(w http.ResponseWriter, r *http.Request) {
	helpers.Header(w, "PUT")
	userId := w.Header().Get("id")
	params := mux.Vars(r)
	addressId := params["address_id"]
	defer r.Body.Close()

	_, err := services.GetUser(userId, "")
	if err != nil {
		helpers.ResponseErrorSender(w, userresponse.USER_GET_FAILED, helpers.FAILED, http.StatusNotFound)
		return
	}
	// decode the address
	var address models.Address
	json.NewDecoder(r.Body).Decode(&address)
	if addressBodyCheck, addressBodyCheckErr := address.AddressBodyCheck(); addressBodyCheck {
		helpers.ResponseErrorSender(w, addressBodyCheckErr, helpers.FAILED, http.StatusBadRequest)
		return
	}

	_, updateAddressErr := services.UpdateAddress(addressId, address)
	if updateAddressErr != nil {
		helpers.ResponseErrorSender(w, addressresponse.UPDATE_ADDRESS_FAILED, helpers.FAILED, http.StatusInternalServerError)
		return
	}
	updateUser, updateUserErr := services.GetUser(userId, "")
	if updateUserErr != nil {
		helpers.ResponseErrorSender(w, userresponse.USER_FETCH_FAILED, helpers.FAILED, http.StatusNotFound)
		return
	}
	helpers.ResponseSuccess(w, addressresponse.UPDATE_ADDRESS_SUCCESSFULLY, helpers.FAILED, http.StatusOK, map[string]interface{}{"data": updateUser})

}

func DeleteAddressById(w http.ResponseWriter, r *http.Request) {
	helpers.Header(w, "PUT")
	userId := w.Header().Get("id")
	params := mux.Vars(r)
	addressId := params["address_id"]
	defer r.Body.Close()

	user, err := services.GetUser(userId, "")
	if err != nil {
		helpers.ResponseErrorSender(w, userresponse.USER_GET_FAILED, helpers.FAILED, http.StatusNotFound)
		return
	}
	deleteAddress, deleteAddressErr := services.DeleteAddressById(user.Id.Hex(), addressId)
	if deleteAddressErr != nil {
		helpers.ResponseErrorSender(w, addressresponse.DELETE_ADDRESS_FAILED, helpers.FAILED, http.StatusInternalServerError)
		return
	}
	if deleteAddress == 0 {
		helpers.ResponseErrorSender(w, addressresponse.DELETE_ADDRESS_FAILED, helpers.FAILED, http.StatusUnprocessableEntity)
		return
	}

	getUser, getUserErr := services.GetUser(user.Id.Hex(), "")
	if getUserErr != nil {
		helpers.ResponseErrorSender(w, userresponse.USER_FETCH_FAILED, helpers.FAILED, http.StatusNotFound)
		return
	}
	helpers.ResponseSuccess(w, addressresponse.DELETE_ADDRESS_SUCCESSFULLY, helpers.SUCCESS, http.StatusOK, map[string]interface{}{"data": getUser})

}
