package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shaikdev/groshop-server/controllers"
)

func UserRoute(r *mux.Router) {

	userRouter := r

	userRouter.HandleFunc("/register", controllers.UserRegister).Methods("POST")

	userRouter.HandleFunc("/login", controllers.UserLogin).Methods("POST")

	userRouter.Handle("/user", controllers.VerifyToken(http.HandlerFunc(controllers.GetUser))).Methods("GET")

	userRouter.Handle("/users", controllers.VerifyToken(http.HandlerFunc(controllers.GetUsers))).Methods("GET")

	userRouter.Handle("/user", controllers.VerifyToken(http.HandlerFunc(controllers.UpdateUser))).Methods("PUT")

	userRouter.Handle("/users", controllers.VerifyToken(http.HandlerFunc(controllers.DeleteUsers))).Methods("DELETE")

	userRouter.Handle("/user", controllers.VerifyToken(http.HandlerFunc(controllers.DeleteUserById))).Methods("DELETE")

}
