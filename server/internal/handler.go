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
	d := datastore.GetDStoreOps()
	err = d.CreateUser()
	return err
}

func LoginUser(userinfo []byte) error {
	lu := datastore.LoginDetails{}
	err := json.Unmarshal(userinfo, &lu)
	if err != nil {
		return err
	}
	d := datastore.GetDStoreOps()
	err = d.LoginUser()
	return err

}
