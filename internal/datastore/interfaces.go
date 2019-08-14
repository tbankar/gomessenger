package datastore
import 

type DstoreOps interface {
	create_user(email string, username string, name string) (bool, error)
	get_user_details(userid string) (error, UserDetails)
}

var dbops DstoreOps
