package core

import (
	"encoding/json"
	"testing"

	"fmt"

	"log"
	"net/url"

	"flag"

	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/gorilla/websocket"
)

var (
	Type_       = "messageService"
	Command     = "send message"
	UserName    = "use90ddgOne"
	NickName    = "nick"
	Password    = "1234324234"
	CurrentTime = 1500373713
	Token       = "333"
	Text        = "msg from user 1 to user 2"
	header      = messageService.MessageHeader{Type_: "auth", Command: "registrissucc", UserName: UserName, Token: Token}

	body = messageService.User{UserName: UserName, Password: Password, NickName: NickName}
	addr = flag.String("addr", "localhost:3002", "http service address")
)

func TestValidateMessage(t *testing.T) {

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/message"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Fatal("dial:", err)
	}

	defer c.Close()

	go func() {

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()
	byteUser, err := json.Marshal(body)
	if err != nil {

	}

	message := &messageService.Message{Header: header, Body: byteUser}
	ValidateMessage(message, 1, c)
	fmt.Println("ok")

}
