package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shaikdev/groshop-server/helpers"
	menuresponse "github.com/shaikdev/groshop-server/helpers/menu_response"
	"github.com/shaikdev/groshop-server/models"
	"github.com/shaikdev/groshop-server/services"
)

func CreateMenu(w http.ResponseWriter, r *http.Request) {
	helpers.Header(w, "POST")

	// decode
	var menu models.Menu
	json.NewDecoder(r.Body).Decode(&menu)

	//TODO: check body
	if menuBody, menuBodyErr := menu.MenuBodyCheck(); menuBody {
		helpers.ResponseErrorSender(w, menuBodyErr, helpers.FAILED, http.StatusBadRequest)
		return
	}

	//TODO: check menu exist or not
	getMenu, _ := services.GetMenu("", menu.Name)
	if getMenu.Name == menu.Name {
		helpers.ResponseErrorSender(w, menuresponse.MENU_ALREADY_EXIST, helpers.FAILED, http.StatusBadRequest)
		return
	}

	menu.CreatedAt = time.Now()
	menu.ModifiedAt = time.Now()
	menu.IsDeleted = false

	createMenu, err := services.CreateMenu(menu)
	if err != nil {
		helpers.ResponseErrorSender(w, menuresponse.CREATE_FAILED, helpers.FAILED, http.StatusUnprocessableEntity)
	}

	getCreatedMenu, createdMenuErr := services.GetMenu(createMenu.Hex(), "")
	if createdMenuErr != nil {
		helpers.ResponseErrorSender(w, menuresponse.GET_FAILED, helpers.FAILED, http.StatusNotFound)
		return
	}
	helpers.ResponseSuccess(w, menuresponse.CREATE_SUCCESS, helpers.SUCCESS, http.StatusCreated, map[string]interface{}{"data": getCreatedMenu})

}

func GetMenu(w http.ResponseWriter, r *http.Request) {
	helpers.Header(w, "GET")
	params := mux.Vars(r)
	if getMenu, err := services.GetMenu(params["menuId"], ""); err != nil {
		helpers.ResponseErrorSender(w, menuresponse.GET_FAILED, helpers.FAILED, http.StatusNotFound)
		return
	} else {
		helpers.ResponseSuccess(w, menuresponse.GET_SUCCESS, helpers.SUCCESS, http.StatusOK, map[string]interface{}{"data": getMenu})
		return
	}
}

func GetMenus(w http.ResponseWriter, r *http.Request) {
	helpers.Header(w, "GET")
	if getMenus, err := services.GetMenus(); err != nil {
		helpers.ResponseErrorSender(w, menuresponse.GET_MANY_FAILED, helpers.FAILED, http.StatusNotFound)
	} else {
		helpers.ResponseSuccess(w, menuresponse.GET_SUCCESS, helpers.SUCCESS, http.StatusOK, map[string]interface{}{"data": getMenus})
		return
	}
}

func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	helpers.Header(w, "PUT")
	params := mux.Vars(r)
	//TODO: check menu exist or not
	if _, menuErr := services.GetMenu(params["menuId"], ""); menuErr != nil {
		helpers.ResponseErrorSender(w, menuresponse.GET_FAILED, helpers.FAILED, http.StatusNotFound)
		return
	}
	// decode menu
	var menu models.Menu
	json.NewDecoder(r.Body).Decode(&menu)

	//TODO: check body
	if menuBody, menuBodyErr := menu.MenuBodyCheck(); menuBody {
		helpers.ResponseErrorSender(w, menuBodyErr, helpers.FAILED, http.StatusBadRequest)
		return
	}

	if _, updateErr := services.UpdateMenu(params["menuId"], menu); updateErr != nil {
		helpers.ResponseErrorSender(w, menuresponse.EDIT_FAILED, helpers.FAILED, http.StatusUnprocessableEntity)
		return
	}
	getMenu, getMenuErr := services.GetMenu(params["menuId"], "")
	if getMenuErr != nil {
		helpers.ResponseErrorSender(w, menuresponse.EDIT_FAILED, helpers.FAILED, http.StatusUnprocessableEntity)
		return
	}
	helpers.ResponseSuccess(w, menuresponse.EDIT_SUCCESS, helpers.SUCCESS, http.StatusOK, map[string]interface{}{"data": getMenu})

}

func DeleteMenu(w http.ResponseWriter, r *http.Request) {
	helpers.Header(w, "DELETE")
	params := mux.Vars(r)

	// TODO: check menu exist or not
	if _, menuErr := services.GetMenu(params["menuId"], ""); menuErr != nil {
		helpers.ResponseErrorSender(w, menuresponse.GET_FAILED, helpers.FAILED, http.StatusNotFound)
		return
	}

	if _, deleteErr := services.DeleteMenu(params["menuId"]); deleteErr != nil {
		helpers.ResponseErrorSender(w, menuresponse.DELETE_FAILED, helpers.FAILED, http.StatusUnprocessableEntity)
		return
	}
	helpers.ResponseSuccess(w, menuresponse.DELETE_SUCCESS, helpers.SUCCESS, http.StatusOK, map[string]interface{}{})

}

func DeleteMenus(w http.ResponseWriter, r *http.Request) {
	helpers.Header(w, "DELETE")
	if _, deleteErr := services.DeleteMenus(); deleteErr != nil {
		helpers.ResponseErrorSender(w, menuresponse.DELETE_MANY_FAILED, helpers.FAILED, http.StatusUnprocessableEntity)
		return
	}

	helpers.ResponseSuccess(w, menuresponse.DELETE_MANY_SUCCESS, helpers.SUCCESS, http.StatusOK, map[string]interface{}{})
}
