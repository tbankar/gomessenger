package internal

import (
	"encoding/json"
	"gomessenger/server/internal/datastore"
)

func stopOnError(err error, errMsg string) {

}
func CreateUser(userinfo []byte) {
	cu := datastore.UserDetails{}
	err := json.Unmarshal(userinfo, &cu)
	stopOnError(err, "createUser: json unmarshal error")
	datastore.SendCreateUser(&cu)
}
