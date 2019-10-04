package internal

import (
	"encoding/json"
	"gomessenger/server/internal/datastore"
)

func CreateUser(userinfo []byte) error {
	cu := datastore.UserDetails{}
	err := json.Unmarshal(userinfo, &cu)
	if err != nil {
		return err
	}
	var d datastore.DstoreOps
	d = cu
	err = d.CreateUser()
	return err
}
