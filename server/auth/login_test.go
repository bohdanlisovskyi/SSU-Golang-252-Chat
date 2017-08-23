package auth

import (
	"fmt"
	"testing"

	"flag"

	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
)

func TestLogin(t *testing.T) {
	test_data := messageService.Authentification{
		UserName: "vasya",
		Password: "vasya2vasya444",
	}
	user, token, err := Login(test_data.UserName, test_data.Password)
	fmt.Println(user, token, err)
}

var (
	Type_    = "messageService"
	Command  = "send message"
	UserName = "num0ddgOne"
	//NickName    = "nick"
	Password    = "1234324234"
	CurrentTime = 1500373713
	Token       = "333"
	Text        = "msg from user 1 to user 2"
	header      = messageService.MessageHeader{Type_: "auth", Command: "authissucc", UserName: UserName, Token: Token}

	//body = messageService.Authentification{UserName: UserName, Password: Password, NickName: NickName}
	addr = flag.String("addr", "localhost:3006", "http service address")
)

//func TestVerifyToken(t *testing.T) {
//	u := url.URL{Scheme: "ws", Host: *addr, Path: "/message"}
//	t.Logf("connecting to %s", u.String())
//
//	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
//
//	if err != nil {
//		t.Logf("dial:", err)
//	}
//
//	defer c.Close()
//
//	go func() {
//
//		for {
//			_, message, err := c.ReadMessage()
//			if err != nil {
//				t.Logf("read:", err)
//				return
//			}
//			t.Logf("recv: %s", message)
//		}
//	}()
//	byteUser, err := json.Marshal(body)
//	if err != nil {
//
//	}
//	clientStruc := customers.Client{Conn: c, Token: "333"}
//	message := &messageService.Message{Header: header, Body: byteUser}
//	res, err := VerifyToken(message, clientStruc)
//	if err != nil {
//		t.Log(err)
//	}
//
//	t.Log("Verification result: ", res)
//}

func TestSendContacts(t *testing.T) {
	test_user := messageService.Authentification{
		UserName: "Yuriy",
	}
	SendContacts(&test_user)

}
