package listener

import (
	"github.com/8tomat8/SSU-Golang-252-Chat/client/config"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/gorilla/websocket"
)

var AuthorizationChannel chan messageService.Message
var MessageChannel chan messageService.Message
var ContactsChannel chan messageService.Message
var SettingsChannel chan messageService.Message
var QuitChannel chan struct{}

func ListenToServer(conn *websocket.Conn) {
	defer conn.Close()
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			loger.Log.Errorf("Can not read message from server. %s", err)
			break
		}
		unmarshaledMessage, err := messageService.UnmarshalMessage(data)
		if err != nil {
			loger.Log.Warningf("Can not unmarshal message. %s", err)
			continue
		}
		ValidateAndRedirectMessage(unmarshaledMessage)
	}
	//when we get connection error, this code executes
	//also when we close connection from another part of application
	//so for close channels enough to close connection
	close(QuitChannel)
}

func ValidateAndRedirectMessage(message *messageService.Message) {
	switch message.Header.Type_ {
	case config.GetConfig().MessageType.Message:
		MessageChannel <- *message
	case config.GetConfig().MessageType.Settings:
		SettingsChannel <- *message
	case config.GetConfig().MessageType.Auth:
		AuthorizationChannel <- *message
	case config.GetConfig().MessageType.Contacts:
		ContactsChannel <- *message
	default:
		loger.Log.Warningf("Can not recognize message type of received message. Received %s type.", message.Header.Type_)
	}
}
