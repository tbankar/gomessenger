package datastore

import (
	"log"

	"github.com/tsuna/gohbase"
)

func getConnHbase() gohbase.Client {
	client := gohbase.NewClient("172.17.0.2")
	return client

}

func (c *UserDetails) CreateUser() (bool, error) {
	client := getConnHbase()
	if client != nil {
		log.Fatal(client)
	}
	//TODO: Generate UserID
	return true, nil
}
