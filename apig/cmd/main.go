package main

import (
	"net/http"

	"gomessenger/apig/api"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/create", api.CreateUser).Methods("POST")
	router.HandleFunc("/sendmsg", api.SendMsg).Methods("POST")
	http.ListenAndServe(":8000", router)
}
