package datastore

type DstoreOps interface {
	create_user(email string, username string, name string) (error, bool)
	get_user_details(userid string) (error, UserDetails)
}
