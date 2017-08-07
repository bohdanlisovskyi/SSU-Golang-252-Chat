package core

import (
	"encoding/json"
	"net/http"
	"testing"

	"fmt"

	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
)

var (
	Type_       = "messageService"
	Command     = "send message"
	UserName    = "userOne"
	NickName    = "nick"
	Password    = "1234324234"
	CurrentTime = 1500373713
	Token       = "token"
	Text        = "msg from user 1 to user 2"
	header      = messageService.MessageHeader{Type_: "authorization", Command: "registrissucc", UserName: UserName, Token: Token}

	body = messageService.User{UserName: UserName, Password: Password, NickName: NickName}

	byteMessage = []byte(`{"header":{"type":"messageService","command":"send message","sender":1,"receiver":2,"time":1500373713,"auth":"token"},"body":{"text":"msg from user 1 to user 2"}}`)
	byteHeader  = []byte(`{"type": "message", "command": "sentMessage", "sender": 1, "receiver": 2, "time": "hereAndNow", "auth": "token"}`)
	byteBody    = []byte(`{"text": "msg from user 1 to user 2"}`)
	conn, err   = addNewConnect(*new(http.ResponseWriter), new(http.Request))
)

func TestValidateMessage(t *testing.T) {
	byteUser, err := json.Marshal(body)
	if err != nil {

	}
	message := &messageService.Message{Header: header, Body: byteUser}
	ValidateMessage(message, 1, conn)
	fmt.Println("ok")

}
