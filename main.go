package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shaikdev/groshop-server/routes"
)

func main() {

	r := mux.NewRouter()
	routes.UserRoute(r.PathPrefix("/api/v1").Subrouter())
	http.Handle("/", r)
	http.ListenAndServe(":8081", nil)

}
