package datastore

type DstoreOps interface {
	CreateUser() (bool, error)
}

func SendCreateUser(d DstoreOps) {
	d.CreateUser()
}
