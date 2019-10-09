package datastore

import (
	"context"
	"errors"
	"fmt"
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

// CreateUser function stores user related information in DB
func (c UserDetails) CreateUser() error {
	client := getConnHbase()
	if client == nil {
		return errors.New("Error while connecting to HBase")
	}
	defer client.Close()

	id := genUUID()
	c.ID = id.String()

	//Can optimize following
	values := map[string]map[string][]byte{FAMILYUSERS: map[string][]byte{"ID": []byte(c.ID),
		"username": []byte(c.Username), "email": []byte(c.Email), "fullname": []byte(c.FullName), "password": []byte(c.Password),
		"sourceIPAddr": []byte(c.SourceIPAddr)}}
	putRequest, err := hrpc.NewPutStr(context.Background(), "gomessenger", string(os.Getpid())+string(time.Now().Unix()), values)
	if err != nil {
		return err
	}
	_, err = client.Put(putRequest)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
