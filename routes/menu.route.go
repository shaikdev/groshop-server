package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shaikdev/groshop-server/controllers"
)

func MenuRoute(r *mux.Router) {
	menuRoute := r

	menuRoute.Handle("/create_menu", controllers.VerifyToken(http.HandlerFunc(controllers.CreateMenu))).Methods("POST")

	menuRoute.Handle("/get_menu/{menuId}", controllers.VerifyToken(http.HandlerFunc(controllers.GetMenu))).Methods("GET")

	menuRoute.Handle("/get_many_menus", controllers.VerifyToken(http.HandlerFunc(controllers.GetMenus))).Methods("GET")

	menuRoute.Handle("/edit_menu/{menuId}", controllers.VerifyToken(http.HandlerFunc(controllers.UpdateMenu))).Methods("PUT")

	menuRoute.Handle("/delete_menu/{menuId}", controllers.VerifyToken(http.HandlerFunc(controllers.DeleteMenu))).Methods("DELETE")

	menuRoute.Handle("/delete_many_menus", controllers.VerifyToken(http.HandlerFunc(controllers.DeleteMenus))).Methods("DELETE")
}
