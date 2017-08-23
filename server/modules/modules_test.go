package modules

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"flag"

	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/gorilla/websocket"
)

var (
	UserName = "Vasya"
	NickName = "Vas"
	Token    = "333"
	header   = messageService.MessageHeader{Type_: "contcts", Command: "contactsrequest", UserName: UserName, Token: Token}

	body = messageService.Authentification{UserName: UserName, NickName: NickName}
	addr = flag.String("addr", "localhost:3006", "http service address")
)

func TestSendContacts(t *testing.T) {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/message"}
	t.Logf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		t.Logf("dial:", err)
	}

	defer c.Close()

	go func() {

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				t.Logf("read:", err)
				return
			}
			t.Logf("recv: %s", message)
		}
	}()
	byteUser, err := json.Marshal(body)
	test_user := messageService.Message{
		Header: header,
		Body:   byteUser,
	}
	SendContacts(&test_user, c)
	fmt.Println("finish")

}
