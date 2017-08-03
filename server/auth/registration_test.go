package auth

import (
	"testing"

	"fmt"

	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
)

var (
	UserName    = "userOne"
	Password    = "password"
	Token       = "token"
	byteMessage = []byte(`{"header":{"type":"messageService","command":"send message","sender":1,"receiver":2,"time":1500373713,"auth":"token"},"body":{"text":"msg from user 1 to user 2"}}`)
	byteHeader  = []byte(`{"type": "message", "command": "sentMessage", "sender": 1, "receiver": 2, "time": "hereAndNow", "auth": "token"}`)
	byteBody    = []byte(`{"text": "msg from user 1 to user 2"}`)
)

func TestRegisterNewUser(t *testing.T) {

	test_user := messageService.User{
		UserName: "ia",
		Password: "ia",
		NickName: "ioa",
	}
	user, token, err := RegisterNewUser(&test_user)
	fmt.Println(user, token, err)
}
