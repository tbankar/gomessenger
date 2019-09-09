package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"gomessenger/apig/internal/datastore"
)

type InputReq struct {
	Username     string `json:"username"`
	UserFullname string `json:"fullname"`
	UserEmail    string `json:"email"`
	Password     string `json:"password"`
}

type ValidateUserInput struct {
	Username     string
	Password     string
	Email        string
	UserFullname string
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	created := make(chan bool)
	errChan := make(chan error)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	var userDetails InputReq
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}
	json.Unmarshal(reqBody, &userDetails)

	ok, err := datastore.IsUserExists(userDetails.Username, "username")
	if err != nil {
		fmt.Fprintf(w, "Error while checking existing user:%v", err)
	}
	if ok {
		w.Write([]byte("User already exists"))
	} else {
		host := datastore.MapUserToServer(userDetails.Username)
		host = "localhost"
		go CallCreateUser(&userDetails, host, created, errChan)
		select {
		case <-created:
			w.WriteHeader(201)
		case <-errChan:
			w.Write([]byte(fmt.Sprintf("%s", err)))
		case <-time.After(3 * time.Second):
			w.Write([]byte("No response received"))
		}
	}
}

func DoLogin(w http.ResponseWriter, r *http.Request) {
	var userDetails InputReq
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}
	json.Unmarshal(reqBody, &userDetails)

	ok, err := datastore.IsUserExists(userDetails.Password, "password")
	if err != nil {
		fmt.Fprintf(w, "Error while checking existing user:%v", err)
	}
	if ok {
		//Check password here
	}

}
