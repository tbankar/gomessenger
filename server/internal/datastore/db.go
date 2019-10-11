package datastore

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

var globalCounter *int64 = new(int64)

const (
	FAMILYUSERS        = "user_details"
	FAMILYACTUSERS     = "active_users"
	FAMILYLOGINDETAILS = "login_details"
	FAMILYMSGS         = "messages"
	FAMILYUIDMAP       = "user_server_mapping"
)

func genUUID() uuid.UUID {
	return uuid.New()
}

func getConnHbase() gohbase.Client {
	//We can always get "hbasedb" either from config or from ENV variable
	client := gohbase.NewClient("hbasedb")
	return client
}

func putRequestToHbase(data map[string]map[string][]byte) error {
	client := getConnHbase()
	if client == nil {
		return errors.New("Error while connecting to HBase")
	}
	defer client.Close()

	putRequest, err := hrpc.NewPutStr(context.Background(), "gomessenger", string(os.Getpid())+string(time.Now().Unix()), data)
	if err != nil {
		return err
	}
	_, err = client.Put(putRequest)
	return err
}

// CreateUser function stores user related information in DB
func (c UserDetails) CreateUser() error {

	id := genUUID()
	c.ID = id.String()

	//Can optimize following
	values := map[string]map[string][]byte{FAMILYUSERS: map[string][]byte{"ID": []byte(c.ID),
		"username": []byte(c.Username), "email": []byte(c.Email), "fullname": []byte(c.FullName), "password": []byte(c.Password),
		"sourceIPAddr": []byte(c.SourceIPAddr)}}
	err := putRequestToHbase(values)
	return err
}

func (l LoginDetails) LoginUser() error {
	client := getConnHbase()
	if client == nil {
		return errors.New("Error while connecting to HBase")
	}
	defer client.Close()
	var loginstat string
	if l.LoginStatus {
		loginstat = "SUCCESS"
	} else {
		loginstat = "FAILED"
	}

	values := map[string]map[string][]byte{FAMILYLOGINDETAILS: map[string][]byte{"username": []byte(l.UserName),
		"login_status": []byte(loginstat), "source_ipaddr": []byte(l.SourceIPAddr)}}

	err := putRequestToHbase(values)
	if err != nil {
		return err
	}
	if l.LoginStatus {
		values = map[string]map[string][]byte{FAMILYACTUSERS: map[string][]byte{"username": []byte(l.UserName)}}
		err = putRequestToHbase(values)
	}
	return err
}
