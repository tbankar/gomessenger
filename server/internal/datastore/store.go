package datastore

type UserDetails struct {
	ID           string
	Email        string
	Username     string
	FullName     string
	Password     string
	SourceIPAddr string
}

type LoginDetails struct {
	UserName     string
	LoginStatus  bool
	LogoutStatus bool
	SourceIPAddr string
}

type Chat struct {
	FromUserID   string
	Message      string
	TimeStamp    int64
	Status       bool
	ToUserID     string
	SourceIPAddr string
}

type UserServerMap struct {
	UserID   string
	ServerIP string
}
