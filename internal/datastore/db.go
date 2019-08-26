package datastore

import (
	"context"
	"sync/atomic"

	"github.com/google/uuid"

	"github.com/tsuna/gohbase/filter"

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
	client := gohbase.NewClient("172.17.0.2")
	return client
}

func isUserExists(uname string, cli gohbase.Client) (bool, error) {
	defer cli.Close()
	b := filter.NewByteArrayComparable([]byte(uname))
	comparator := filter.NewBinaryComparator(b)
	bFilter := filter.NewSingleColumnValueFilter([]byte(FAMILYUSERS), []byte("username"), filter.Equal, comparator, false, true)
	scanReq, err := hrpc.NewScanStr(context.Background(), "gomessenger", hrpc.Filters(bFilter))
	if err != nil {
		return false, err
	}
	scanResp, err := cli.Scan(scanReq).Next()
	if err != nil {
		return false, err
	}
	scanLen := len(scanResp.Cells)
	if scanLen == 0 {
		return true, nil
	}
	return false, nil

}

func (c *UserDetails) CreateUser() (bool, error) {
	client := getConnHbase()
	defer client.Close()

	ok, err := isUserExists(c.Username, client)
	if err != nil {
		return false, err
	}
	if ok {
		rowCnt := atomic.AddInt64(globalCounter, 1)
		id := genUUID()
		c.UserID = id.String()
		values := map[string]map[string][]byte{FAMILYUSERS: map[string][]byte{"userid": []byte(c.UserID), "username": []byte(c.Username), "email": []byte(c.Useremail), "fullname": []byte(c.Name)}}
		putRequest, err := hrpc.NewPutStr(context.Background(), "gomessenger", string(rowCnt), values)
		if err != nil {
			return false, err
		}
		_, err = client.Put(putRequest)
		if err != nil {
			return false, err
		}
	}
	if client == nil {
		return false, nil
	}
	return true, nil
}
