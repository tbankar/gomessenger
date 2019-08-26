package datastore

type UserDetails struct {
	UserID    string
	Useremail string
	Username  string
	Name      string
}

type LoginDetails struct {
	UserID      string
	LoginTS     int64
	LogoutTS    int64
	LoginIPAddr string
}

type Chat struct {
	FromUserID string
	Message    string
	TimeStamp  int64
	Status     bool
	ToUserID   string
}

type UserServerMap struct {
	UserID   string
	ServerIP string
}
