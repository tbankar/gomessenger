package main

import (
	"net/http"

	"github.com/gomessenger/apig/api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/create", api.CreateUser).Methods("POST")
	http.ListenAndServe(":8000", router)
}
