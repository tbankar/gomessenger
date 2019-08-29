package datastore

import (
	"context"
	"errors"
	"io"

	"github.com/tsuna/gohbase/filter"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

const FAMILYUSERS = "user_details"

func IsUserExists(uname string) (bool, error) {
	//Change IP address to hostname hbasedb
	client := gohbase.NewClient("172.17.0.2")
	if client == nil {
		return false, errors.New("Error while connecting to HBase")
	}
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
