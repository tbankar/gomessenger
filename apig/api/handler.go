package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gomessenger/apig/internal/datastore"
)

type InputReq struct {
	Username     string `json:"username"`
	UserFullname string `json:"fullname"`
	UserEmail    string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var userDetails InputReq
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}
	json.Unmarshal(reqBody, &userDetails)

	ok, err := datastore.IsUserExists(userDetails.Username)
	if err != nil {
		fmt.Fprintf(w, "Error while checking existing user:%v", err)
	}
	if !ok {
		w.Write([]byte("User already exists"))
	} else {
		host := datastore.MapUserToServer(userDetails.Username)
		msg, err := CallCreateUser(&userDetails, host)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("%s", err)))
		} else if msg != "" {
			w.Write([]byte(fmt.Sprintf("%s", msg)))
		} else {
			w.WriteHeader(200)
		}
	}
}

func SendMsg() (bool, error) {

}
