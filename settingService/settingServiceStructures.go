package settingService

// RequestHeader is a standard structure for header
type RequestHeader struct {
	Type_    string `json:"type"`
	Command  string `json:"command"`
	UserName string `json:"userName"`
	Token    string `json:"token"`
}

// UserInfo is a structure for passing info related to contacts
type UserInfo struct {
	UserName  string `json:"user_name"`
	NickName  string `json:"nick_name"`
	Birthday  int `json:"birthday"`
	IsBlocked int `json:"is_blocked"`
}
