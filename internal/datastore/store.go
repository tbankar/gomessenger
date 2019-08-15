package datastore

type UserID string

type UserDetails struct {
	UserID    UserID
	Useremail string
	Username  string
	Name      string
}

type LoginDetails struct {
	UserID      UserID
	LoginTS     int64
	LogoutTS    int64
	LoginIPAddr string
}

type Chat struct {
	FromUserID UserID
	Message    string
	TimeStamp  int64
	Status     bool
	ToUserID   UserID
}

type UserServerMap struct {
	UserID   UserID
	ServerIP string
}
