package auth

type UserInfo struct {
	User string `json:"user"`
	Org  int    `json:"org"`
}
