package routes

import (
	"github.com/gorilla/mux"
	"github.com/shaikdev/groshop-server/controllers"
)

func UserRoute(r *mux.Router) {

	userRouter := r.PathPrefix("/auth").Subrouter()

	userRouter.HandleFunc("/create_user", controllers.UserCreate).Methods("POST")

}
