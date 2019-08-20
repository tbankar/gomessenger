package datastore

import (
	"context"
	"errors"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

const (
	FAMILYUSERS        = "user_details"
	FAMILYACTUSERS     = "active_users"
	FAMILYLOGINDETAILS = "login_details"
	FAMILYMSGS         = "msgs"
	FAMILYUIDMAP       = "user_to_server"
)

func getConnHbase() gohbase.Client {
	client := gohbase.NewClient("172.17.0.2")
	return client
}
func isUserExists(uname string, cli gohbase.Client) (bool, error) {
	family := map[string][]string{FAMILYUSERS: []string{"username"}}
	req, err := hrpc.NewGetStr(context.Background(), "gomessenger", uname, hrpc.Families(family))
	if err != nil {
		return false, err
	}
	resp, err := cli.Get(req)
	if err != nil {
		return false, err
	}
	if *resp.Exists {
		return false, errors.New("User exists")
	}
	return true, nil

}

func (c *UserDetails) CreateUser() (bool, error) {
	client := getConnHbase()
	isUserExists(c.Username, client)
	defer client.Close()
	if client == nil {
		return false, nil
	}
	//TODO: Generate UserID
	return true, nil
}
