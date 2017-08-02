// messager
package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/Greckas/SSU-Golang-252-Chat/messageService"
	"github.com/gorilla/websocket"
	"github.com/Greckas/SSU-Golang-252-Chat/loger"
	"time"
)

//binding to qml
var qmlMessage *QmlMessage

//represent message struct as QObject
type QmlMessage struct {
	core.QObject

	_ func(message string) 			`slot:"sendMessage"`
	_ func(messageWasSent bool)		`signal:"messageSent"`
}

//initialize QmlMessage and its slots
func initQmlMessage(quickWidget *quick.QQuickWidget) {
	qmlMessage = NewQmlMessage(nil)
	quickWidget.RootContext().SetContextProperty("qmlMessage", qmlMessage)
	qmlMessage.ConnectSendMessage(func(message string) {
		newMessageHeader := messageService.MessageHeader{
			Type_:"message", //type will be added to config file in near future
			Command:"sendMessage", //commands will be added to config file in near future
			UserName:"Kitler", //still haven`t user preferences
			Token:"123456789009876", //still haven`t authorization
		}
		newMessageBody := messageService.MessageBody{
			ReceiverName: "Little Kitty", //have to be replaced with Contacts data
			Time:int(time.Now().Unix()),
			Text:message,
		}
		newMessage := messageService.Message{
			Header:newMessageHeader,
			Body:newMessageBody,
		}
		marshaledMessage, err := messageService.MarshalMessage(&newMessage)
		if err != nil {
			qmlMessage.MessageSent(false)
			loger.Log.Errorf("Can`t marshal message. %s", err)
			qmlStatus.SendStatus("Message wasn`t sent")
		} else {
			if err := connection.WriteMessage(websocket.TextMessage, marshaledMessage); err != nil {
				loger.Log.Errorf("Can not send message. %s", err)
				qmlMessage.MessageSent(false)
				qmlStatus.SendStatus("Message wasn`t sent. It`s not your fall. Life is just a PAIN")
				return
			}
			loger.Log.Info("Message was sent successfully")
			qmlMessage.MessageSent(true)
			qmlStatus.SendStatus("Message sent")
		}
	})
}