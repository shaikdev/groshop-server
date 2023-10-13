package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shaikdev/groshop-server/controllers"
)

func AddressRoute(r *mux.Router) {
	addressRouter := r
	addressRouter.Handle("/create_address", controllers.VerifyToken(http.HandlerFunc(controllers.CreateAddress))).Methods("PUT")
	addressRouter.Handle("/update_address/{address_id}", controllers.VerifyToken(http.HandlerFunc(controllers.UpdateAddress))).Methods("PUT")
	addressRouter.Handle("/delete_address/{address_id}", controllers.VerifyToken(http.HandlerFunc(controllers.DeleteAddressById))).Methods("DELETE")
}
