package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

//func for add new routes

func Routes() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/", NewRoute)

	//add even more routes
	return route
}

type Mess struct {
}

func NewRoute(w http.ResponseWriter, r *http.Request) {

}
