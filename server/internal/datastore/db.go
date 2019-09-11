package datastore

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync/atomic"

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
	// client := gohbase.NewClient("hbasedb")
	client := gohbase.NewClient("172.17.0.2")
	return client
}

func (c *UserDetails) CreateUser() (bool, error) {
	client := getConnHbase()
	if client == nil {
		return false, errors.New("Error while connecting to HBase")
	}
	defer client.Close()

	rowCnt := strconv.FormatInt(atomic.AddInt64(globalCounter, 1), 10)
	id := genUUID()
	c.UserID = id.String()
	values := map[string]map[string][]byte{FAMILYUSERS: map[string][]byte{"userid": []byte(c.UserID), "username": []byte(c.Username), "email": []byte(c.Useremail), "fullname": []byte(c.Name), "password": []byte(c.Password)}}
	putRequest, err := hrpc.NewPutStr(context.Background(), "gomessenger", rowCnt, values)
	if err != nil {
		return false, err
	}
	_, err = client.Put(putRequest)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}
