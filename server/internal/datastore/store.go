package datastore

type UserDetails struct {
	ID       string
	Email    string
	Username string
	FullName string
	Password string
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
