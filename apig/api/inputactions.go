package api

const (
	CREATE    = "create"
	SEND      = "send"
	LISTUSERS = "listusers"
	ACTION    = "action"
)

type CreateInputReq struct {
	Username     string `json:"username"`
	UserFullname string `json:"fullname"`
	UserEmail    string `json:"email"`
	Password     string `json:"password"`
}

type LoginResp struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statuscode"`
}

type LoginInputReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
