package datastore

type DstoreOps interface {
	CreateUser() error
}

var dstoreOps DstoreOps
