package auth

//UserInfo 用户信息
type UserInfo struct {
	User string `json:"user"`
	Org  int    `json:"org"`
}
