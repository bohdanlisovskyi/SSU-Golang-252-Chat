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

	_ func(message string, currentIndex int) `slot:"sendMessage"`
	_ func(messageWasSent bool)              `signal:"messageSent"`
}

//initialize QmlMessage and its slots
func initQmlMessage(quickWidget *quick.QQuickWidget) {
	qmlMessage = NewQmlMessage(nil)
	quickWidget.RootContext().SetContextProperty("qmlMessage", qmlMessage)
	qmlMessage.ConnectSendMessage(sendMessage)
}

func sendMessage(message string, currentIndex int) {
	const MessageWasntSent = "Message wasn`t sent"
	newMessageHeader := messageService.MessageHeader{
		Type_:    config.GetConfig().MessageType.Message,
		Command:  config.GetConfig().MessageCommand.SendMessage,
		UserName: userinfo.CurrentUserInfo.UserName,
		Token:    userinfo.CurrentUserInfo.Token,
	}
	var iData, exists = listOfContacts.Get(currentIndex)
	if !exists {
		qmlMessage.MessageSent(false)
		loger.Log.Warningf("List has no index %d.", currentIndex)
		qmlStatus.SendStatus(MessageWasntSent)
		return
	}
	var data = iData.(*ContactObject)
	newMessageBody := messageService.MessageBody{
		ReceiverName: data.UserName,
		Time:         int(time.Now().Unix()),
		Text:         message,
	}
	newRawMessageBody, err := json.Marshal(newMessageBody)
	if err != nil {
		qmlMessage.MessageSent(false)
		loger.Log.Warningf("Can not marshal message body. %s", err)
		qmlStatus.SendStatus(MessageWasntSent)
		return
	}
	newMessage := messageService.Message{
		Header: newMessageHeader,
		Body:   newRawMessageBody,
	}
	marshaledMessage, err := messageService.MarshalMessage(&newMessage)
	if err != nil {
		qmlMessage.MessageSent(false)
		loger.Log.Warningf("Can`t marshal message. %s", err)
		qmlStatus.SendStatus(MessageWasntSent)
		return
	}
	if err := connection.WriteMessage(websocket.TextMessage, marshaledMessage); err != nil {
		qmlMessage.MessageSent(false)
		loger.Log.Warningf("Can not send message. %s", err)
		qmlStatus.SendStatus(MessageWasntSent + ". It`s not your fall. Life is just a PAIN")
		return
	}
	qmlStatus.SendStatus("Waiting for server response.")
	loger.Log.Info("Message sent to server. Waiting response.")
	//now we just waiting for server response
}
