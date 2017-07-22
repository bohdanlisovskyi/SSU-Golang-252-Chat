// messager
package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	//"github.com/8tomat8/SSU-Golang-252-Chat/client/config"
	"github.com/gorilla/websocket"
	//"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"time"
	//"net/url"
	//"flag"
	"fmt"
)

//binding to qml
var qmlMessage *QmlMessage

//represent message struct as QObject
type QmlMessage struct {
	core.QObject

	_ func(message string) 			`slot:"sendMessage"`
	_ func(messageWasSent bool)		`slot:"messageSent"`
	_ func(errorMessage string)		`signal:"errorOccured"`
}

//initialize QmlMessage and its slots
func initQmlMessage(quickWidget *quick.QQuickWidget) {
	qmlMessage = NewQmlMessage(nil)
	quickWidget.RootContext().SetContextProperty("qmlMessage", qmlMessage)
	qmlMessage.ConnectSendMessage(func(message string) {
		newMessageHeader := messageService.MessageHeader{
			Type_:"msg", //or messsage
			Command:"sendMessage", //commands will be added to config file in near future
			UserName:"Kitler", //still haven`t user preferences
			Token:"123456789009876", //still haven`t authorization
		}
		newMessageBody := messageService.MessageBody{
			ReceiverName: "Little Kitty", //contactsList.CurrentContactUsername(), //struct Contacts has not full implementation
			Time:int(time.Now().Unix()),
			Text:message,
		}
		newMessage := messageService.Message{
			Header:newMessageHeader,
			Body:newMessageBody,
		}
		conn, resp, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:3002/message", nil)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", resp)
		conn.WriteJSON(newMessage)
		//conn.WriteJSON(msg)

		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 1)
			newMessage.Body.Text = fmt.Sprintf("Message №%d", i)
			err := conn.WriteJSON(newMessage)
			if err != nil {
				panic(err)
			}
		}
		conn.Close()
		/*marshaledMessage, err := messageService.MarshalMessage(newMessage)
		if err != nil {
			qmlMessage.MessageSent(false)
			qmlMessage.ErrorOccured(err.Error())
			loger.Log.Panicf("Can`t marshal message. %s", err)
		} else {
			var addr = flag.String("addr", config.GetConfig().Server.Connection, "http server address")
			//u := url.URL{Scheme:“ws”, Host: *addr, Path: “/message”}

			u := url.URL{Scheme:"ws", Host:*addr, Path:"/message"}
			connection, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			//connection, _, err := websocket.DefaultDialer.Dial(config.GetConfig().Server.Connection, nil)
			if err != nil {
				loger.Log.Fatal("dial:", err)
			}

			if err := connection.WriteMessage(websocket.TextMessage, marshaledMessage); err != nil {
				loger.Log.Panicf("Can not send message. %s", err)
			}

			select {

			}
		}*/
	})
}
