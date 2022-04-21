package server

import (
	"log"
	"net/http"
)

func RunServer() error {
	router := Routes()
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal("Server internal error")
	}
	return err
}
