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

/*type ValidateUserInput struct {
	Username     string
	Password     string
	Email        string
	UserFullname string
}*/

type LoginResponse struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statuscode"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	created := make(chan bool)
	errChan := make(chan error)
	defer close(created)

	var userDetails InputReq
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
	var LResp LoginResponse
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
		LResp = LoginResponse{
			StatusCode: http.StatusUnauthorized,
			Status:     "Failed",
		}
	} else {
		LResp = LoginResponse{
			StatusCode: http.StatusOK,
			Status:     "Success",
		}
	}
	json.NewEncoder(w).Encode(LResp)
}
