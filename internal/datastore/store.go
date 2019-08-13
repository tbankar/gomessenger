package datastore

type UserDetails struct {
	UserID    string
	Useremail string
	Username  string
	Name      string
}

type LoginDetails struct {
	LoginTS  int64
	LogoutTS int64
}
