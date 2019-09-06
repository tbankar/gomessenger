package main

import (
	"log"
	"net/http"
	"os"

	"gomessenger/apig/api"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/create", api.CreateUser).Methods("POST")
	router.HandleFunc("/login", api.DoLogin).Methods("POST")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// start server listen
	// with error handling
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
