package datastore

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/tsuna/gohbase/filter"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

const FAMILYUSERS = "user_details"

func IsUserExists(uname string) (bool, error) {
	// client := gohbase.NewClient("hbasedb")
	client := gohbase.NewClient("172.17.0.2")
	if client == nil {
		return false, errors.New("Error while connecting to HBase")
	}
	fmt.Println("Connected")
	defer client.Close()
	b := filter.NewByteArrayComparable([]byte(uname))
	comparator := filter.NewBinaryComparator(b)
	bFilter := filter.NewSingleColumnValueFilter([]byte(FAMILYUSERS), []byte("username"), filter.Equal, comparator, false, true)
	scanReq, err := hrpc.NewScanStr(context.Background(), "gomessenger", hrpc.Filters(bFilter))
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
