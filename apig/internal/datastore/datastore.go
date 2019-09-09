package datastore

import (
	"context"
	"errors"
	"io"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
)

const FAMILYUSERS = "user_details"

func IsUserExists(uname, password string) (bool, error) {
	// client := gohbase.NewClient("hbasedb")
	client := gohbase.NewClient("172.17.0.3")

	if client == nil {
		return false, errors.New("Error while connecting to HBase")
	}
	defer client.Close()

	b := filter.NewByteArrayComparable([]byte(uname))
	b1 := filter.NewByteArrayComparable([]byte(password))
	comparator := filter.NewBinaryComparator(b)
	comparator1 := filter.NewBinaryComparator(b1)
	bFilter := filter.NewSingleColumnValueFilter([]byte(FAMILYUSERS), []byte("username"), filter.Equal, comparator, false, true)
	bFilter1 := filter.NewSingleColumnValueFilter([]byte(FAMILYUSERS), []byte("password"), filter.Equal, comparator1, false, true)

	scanReq, err := hrpc.NewScanStr(context.Background(), "gomessenger", hrpc.Filters(bFilter), hrpc.Filters(bFilter1))
	if err != nil {
		return false, err
	}

	_, err = client.Scan(scanReq).Next()
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

func MapUserToServer(username string) string {
	if username[0] < 'i' {
		return "useratoi"
	} else if username[0] < 's' {
		return "userjtos"
	}
	return "userttoz"

}
