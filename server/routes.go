package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

//func for add new routes

func Routes() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/", NewHandler)

	//add even more routes
	return route
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	//
}
