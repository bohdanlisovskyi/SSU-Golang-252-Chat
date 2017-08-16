// messager
package ui

import (
	"encoding/json"
	"time"

	"github.com/8tomat8/SSU-Golang-252-Chat/client/config"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/8tomat8/SSU-Golang-252-Chat/userinfo"
	"github.com/gorilla/websocket"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

//binding to qml
var qmlMessage *QmlMessage

//represent message struct as QObject
type QmlMessage struct {
	core.QObject

	_ func(message string)      `slot:"sendMessage"`
	_ func(messageWasSent bool) `signal:"messageSent"`
}

//initialize QmlMessage and its slots
func initQmlMessage(quickWidget *quick.QQuickWidget) {
	qmlMessage = NewQmlMessage(nil)
	quickWidget.RootContext().SetContextProperty("qmlMessage", qmlMessage)
	qmlMessage.ConnectSendMessage(sendMessage)
}

func sendMessage(message string) {
	newMessageHeader := messageService.MessageHeader{
		Type_:    config.GetConfig().MessageType.Message,
		Command:  config.GetConfig().MessageCommand.SendMessage,
		UserName: userinfo.CurrentUserInfo.UserName,
		Token:    userinfo.CurrentUserInfo.Token,
	}
	newMessageBody := messageService.MessageBody{
		ReceiverName: "Little Kitty", //have to be replaced with Contacts data
		Time:         int(time.Now().Unix()),
		Text:         message,
	}
	newRawMessageBody, err := json.Marshal(newMessageBody)
	if err != nil {
		qmlMessage.MessageSent(false)
		loger.Log.Errorf("Can not marshal message body. %s", err)
		qmlStatus.SendStatus("Message wasn`t sent")
		return
	}
	newMessage := messageService.Message{
		Header: newMessageHeader,
		Body:   newRawMessageBody,
	}
	marshaledMessage, err := messageService.MarshalMessage(&newMessage)
	if err != nil {
		qmlMessage.MessageSent(false)
		loger.Log.Errorf("Can`t marshal message. %s", err)
		qmlStatus.SendStatus("Message wasn`t sent")
		return
	}
	if err := connection.WriteMessage(websocket.TextMessage, marshaledMessage); err != nil {
		qmlMessage.MessageSent(false)
		loger.Log.Errorf("Can not send message. %s", err)
		qmlStatus.SendStatus("Message wasn`t sent. It`s not your fall. Life is just a PAIN")
		return
	}
	qmlStatus.SendStatus("Waiting for server response.")
	//now we just waiting for server response
}
