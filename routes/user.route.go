package routes

import (
	"github.com/gorilla/mux"
	"github.com/shaikdev/groshop-server/controllers"
)

func UserRoute(r *mux.Router) {

	userRouter := r.PathPrefix("/auth").Subrouter()

	userRouter.HandleFunc("/register", controllers.UserRegister).Methods("POST")

	userRouter.HandleFunc("/login", controllers.UserLogin).Methods("POST")

}
