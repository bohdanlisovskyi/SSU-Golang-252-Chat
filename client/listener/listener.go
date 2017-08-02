package listener

import (
	"github.com/8tomat8/SSU-Golang-252-Chat/client/resolvers"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/gorilla/websocket"
)

func ListenToServer(conn *websocket.Conn) {
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			loger.Log.Warningf("Can not read message from server. %s", err)
			continue
		}
		unmarshalMessage, err := messageService.UnmarshalMessage(data)
		if err != nil {
			loger.Log.Warningf("Can not unmarshal message. %s", err)
			continue
		}
		ValidateAndRedirectMessage(unmarshalMessage)
		//
	}
}

func ValidateAndRedirectMessage(message *messageService.Message) {
	if message.Header.Type_ == "message" {
		resolvers.ReceiveMessage(message)
	} else if message.Header.Type_ == "settings" {

	}
}
