package main

import (
	"log"
	"net/http"

	"gomessenger/apig/api"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/create", api.CreateUser).Methods("POST")
	router.HandleFunc("/login", api.DoLogin).Methods("POST")
	router.HandleFunc("/onlineusers", api.GetOnlineUsers).Methods("GET")
	router.HandleFunc("/sendmsg", api.SendMsg).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowedHeaders: []string{
			"*",
		},
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe("0.0.0.0:8000", handler))
}
