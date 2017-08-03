package listener

import (
	"github.com/8tomat8/SSU-Golang-252-Chat/client/config"
	"github.com/8tomat8/SSU-Golang-252-Chat/client/resolvers"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/gorilla/websocket"
)

var LoginStatusChannel chan messageService.ResponseFromServerBody
var MessageSentStatusChannel chan messageService.ResponseFromServerBody
var RegisterStatusChannel chan messageService.ResponseFromServerBody
var QuitChannel chan bool

//var MessageChannel chan messageService.MessageBody

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
	switch message.Header.Type_ {
	case config.GetConfig().MessageType.Message:
		resolvers.ReceiveMessage(message, MessageSentStatusChannel)
	case config.GetConfig().MessageType.Settings:
	case config.GetConfig().MessageType.Auth:
	}
}
