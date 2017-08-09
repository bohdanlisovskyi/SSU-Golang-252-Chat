package userinfo

var CurrentUserInfo *UserInfo

type UserInfo struct {
	UserName string
	Token    string
}
