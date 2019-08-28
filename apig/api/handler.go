package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type InputReq struct {
	Username  string `json:"username"`
	Fullname  string `json:"fullname`
	UserEmail string `json:"email`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	inp := InputReq{
		Username:  params["username"],
		UserEmail: params["email"],
		Fullname:  params["fullname"],
	}

	return

}
