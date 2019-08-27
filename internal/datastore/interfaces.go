package datastore

type DstoreOps interface {
	CreateUser() (string, bool, error)
}

func SendCreateUser(d DstoreOps) {
	d.CreateUser()
}
