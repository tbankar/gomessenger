package api

import (
	"encoding/json"
	"fmt"
	"gomessenger/common"
	"io/ioutil"
	"net/http"
	"time"

	"gomessenger/apig/internal/datastore"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

// CreateUser will create a user in a system
func CreateUser(w http.ResponseWriter, r *http.Request) {
	respChan := make(chan string)
	defer close(respChan)

	var userDetails CreateInputReq
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.ResponseToClient(400, "Please enter valid data", w)
		return
	}
	json.Unmarshal(reqBody, &userDetails)

	ok, err := datastore.IsUserExists(userDetails.Username, "")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
	}
	if !ok {
		common.ResponseToClient(400, "User already exists", w)
	} else {
		go CallCreateUser(&userDetails, respChan)
		select {
		case s := <-respChan:
			if s == "1" {
				common.ResponseToClient(201, "Success", w)
			} else {
				common.ResponseToClient(200, s, w)
			}
		case <-time.After(5 * time.Second):
			common.ResponseToClient(503, "Time exceeded while creating a user...Exited", w)
			return
		}
	}
}

//DoLogin function will allow user to login to the system
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
	go datastore.IsUserExists(userDetails.Username, userDetails.Password)
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
