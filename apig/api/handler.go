package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"gomessenger/apig/internal/datastore"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	created := make(chan bool)
	errChan := make(chan error)
	defer close(created)
	defer close(errChan)

	var userDetails CreateInputReq
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}
	json.Unmarshal(reqBody, &userDetails)

	ok, err := datastore.IsUserExists(userDetails.Username, "")
	if err != nil {
		fmt.Fprintf(w, "Error while checking existing user:%v", err)
	}
	if !ok {
		w.Write([]byte("User already exists"))
	} else {
		go CallCreateUser(&userDetails, created, errChan)
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
	var userDetails CreateInputReq
	var LResp LoginResp
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}
	json.Unmarshal(reqBody, &userDetails)

	ok, err := datastore.IsUserExists(userDetails.Username, userDetails.Password)
	if err != nil {
		fmt.Fprintf(w, "Error while checking existing user:%v", err)
	}
	if ok {
		LResp = LoginResp{
			StatusCode: http.StatusUnauthorized,
			Status:     "Failed",
		}
	} else {
		LResp = LoginResp{
			StatusCode: http.StatusOK,
			Status:     "Success",
		}
	}
	json.NewEncoder(w).Encode(LResp)
}
