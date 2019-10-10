package datastore

import (
	"context"
	"errors"
	"io"

	//log "github.com/sirupsen/logrus"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
)

const FAMILYUSERS = "user_details"

func IsUserExists(uname, password string) (bool, error) {
	client := gohbase.NewClient("hbasedb")
	//log.SetLevel(5)

	if client == nil {
		return false, errors.New("Error while connecting to HBase")
	}
	defer client.Close()

	var err error
	var bFilter1 filter.SingleColumnValueFilter
	var scanReq *hrpc.Scan
	if password != "" {
		b1 := filter.NewByteArrayComparable([]byte(password))
		comparator1 := filter.NewBinaryComparator(b1)
		bFilter1 = *filter.NewSingleColumnValueFilter([]byte(FAMILYUSERS), []byte("password"), filter.Equal, comparator1, false, true)
	}

	b := filter.NewByteArrayComparable([]byte(uname))
	comparator := filter.NewBinaryComparator(b)
	bFilter := filter.NewSingleColumnValueFilter([]byte(FAMILYUSERS), []byte("username"), filter.Equal, comparator, false, true)
	if password != "" {
		scanReq, err = hrpc.NewScanStr(context.Background(), "gomessenger", hrpc.Filters(bFilter), hrpc.Filters(&bFilter1))
	} else {
		scanReq, err = hrpc.NewScanStr(context.Background(), "gomessenger", hrpc.Filters(bFilter))
	}

	if err != nil {
		return false, err
	}

	_, err = client.Scan(scanReq).Next()
	if err == io.EOF {
		return false, nil
	}
	return true, err
}

// Idea is to map perticular set of Users to speicific server
func MapUserToServer(username string) string {
	if username[0] < 'i' {
		return "useratoi"
	} else if username[0] < 's' {
		return "userjtos"
	}
	return "userttoz"

}
