package datastore

type DstoreOps interface {
	CreateUser() error
	LoginUser() error
}

var dops DstoreOps

func GetDStoreOps() DstoreOps {
	return dops
}
