package datastore

import (
	"context"
	"fmt"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
)

const (
	FAMILYUSERS        = "user_details"
	FAMILYACTUSERS     = "active_users"
	FAMILYLOGINDETAILS = "login_details"
	FAMILYMSGS         = "messages"
	FAMILYUIDMAP       = "user_server_mapping"
)

func getConnHbase() gohbase.Client {
	client := gohbase.NewClient("172.17.0.2")
	return client
}

func isUserExists(uname string, cli gohbase.Client) (bool, error) {
	defer cli.Close()
	f := filter.ByteArrayComparable{}
	binComp := filter.NewBinaryComparator()
	filter.NewSingleColumnValueFilter(FAMILYUSERS, "userid", filter.Equal, 2, false, true)
	scanRequest, err := hrpc.NewScanStr(context.Background(), "gomessenger",
		hrpc.Filters(pFilter))
	scanRsp, err := cli.Scan(scanRequest)
	return true, nil

}

func createUser(c *UserDetails) (bool, error) {
	cli := getConnHbase()
	defer cli.Close()
	uID := string(c.UserID)
	values := map[string]map[string][]byte{FAMILYUSERS: map[string][]byte{"username": []byte(uID)}}
	putRequest, err := hrpc.NewPutStr(context.Background(), "gomessenger", "row1", values)
	if err != nil {
		return false, err
	}
	rsp, err := cli.Put(putRequest)
	if err != nil {
		return false, err
	}
	fmt.Println(rsp)
	return true, nil

}

func (c *UserDetails) CreateUser() (bool, error) {
	client := getConnHbase()
	ok, err := isUserExists(c.Username, client)
	if err != nil {
		return false, err
	}
	if ok {
		createUser(c)
	}
	if client == nil {
		return false, nil
	}
	//TODO: Generate UserID
	return true, nil
}
