package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"gomessenger/apig/internal/datastore"

	"github.com/gorilla/sessions"
)

const (
	LOGIN  = "login"
	LOGOUT = "logout"
	CREATE = "create"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	published := make(chan bool)
	errChan := make(chan error)
	defer close(published)
	defer close(errChan)

	var userDetails CreateInputReq
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}
	json.Unmarshal(reqBody, &userDetails)

	ok, err := datastore.IsUserExists(userDetails.Username, "")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
	}
	if !ok {
		w.Write([]byte("User already exists"))
	} else {
		go CallCreateUser(&userDetails, published, errChan)
		select {
		case <-published:
			go listenToMQ(CREATE)
			close(errChan)
			close(published)
		case <-errChan:
			w.Write([]byte(fmt.Sprintf("%s", err)))
		case <-time.After(5 * time.Second):
			w.Write([]byte("Time exceeded while creating a user...Exited"))
		}
	}
}

func DoLogin(w http.ResponseWriter, r *http.Request) {
	var userDetails CreateInputReq
	var LResp LoginResp
	success := make(chan bool)
	errChan := make(chan error)
	defer close(success)
	defer close(errChan)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}
	json.Unmarshal(reqBody, &userDetails)
	go datastore.IsUserExists(userDetails.Username, userDetails.Password, success, errChan)
	select {
	case <-success:
		w.WriteHeader(200)
	case <-errChan:
		w.Write([]byte(fmt.Sprintf("%s", err)))
	case <-time.After(5 * time.Second):
		w.Write([]byte("Time exceeded while creating a user"))
	}
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
