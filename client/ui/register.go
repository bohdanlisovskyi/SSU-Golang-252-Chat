// register
package ui

import (
	"encoding/json"
	"net/url"

	"github.com/8tomat8/SSU-Golang-252-Chat/client/config"
	"github.com/8tomat8/SSU-Golang-252-Chat/client/listener"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/gorilla/websocket"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

//binding to qml
var qmlRegister *QmlRegister

//represent register struct as QObject
type QmlRegister struct {
	core.QObject
	_ func(userName, nickName, password string) `slot:"checkRegisterData"`
	_ func(isRegisterValid bool)                `signal:"registerDataIsValid"`
}

//initialize QmlRegister and its slots
func initQmlRegister(quickWidget *quick.QQuickWidget) {
	qmlRegister = NewQmlRegister(nil)
	quickWidget.RootContext().SetContextProperty("qmlRegister", qmlRegister)
	qmlRegister.ConnectCheckRegisterData(func(userName, nickName, password string) {
		if userName == "" || nickName == "" || password == "" {
			loger.Log.Info("Register failed. userName and/or nickName and/or password field are empty")
			qmlStatus.SendStatus("Please, fill all fields.")
			qmlRegister.RegisterDataIsValid(false)
			return
		}
		var err error
		connUrl := url.URL{
			Scheme: "ws",
			Host:   config.GetConfig().Server.Host,
			Path:   config.GetConfig().Server.Path,
		}
		//this connection should be closed just after server response
		connection, _, err = websocket.DefaultDialer.Dial(connUrl.String(), nil)
		if err != nil {
			loger.Log.Warningf("Can not dial. %s", err)
			qmlLogin.LoginDataIsValid(false)
			qmlStatus.SendStatus("Register failed")
			return
		}
		newMessageHeader := messageService.MessageHeader{
			Type_:   config.GetConfig().MessageType.Register,
			Command: config.GetConfig().MessageCommand.SendRegisterData,
		}
		newMessageBody := messageService.Authentification{
			UserName: userName,
			NickName: nickName,
			Password: password,
		}
		newRawMessageBody, err := json.Marshal(newMessageBody)
		if err != nil {
			qmlLogin.LoginDataIsValid(false)
			loger.Log.Warningf("Can`t marshal register message. %s", err)
			qmlStatus.SendStatus("Register data can't be sent")
			return
		}
		newMessage := messageService.Message{
			Header: newMessageHeader,
			Body:   newRawMessageBody,
		}
		marshaledMessage, err := messageService.MarshalMessage(&newMessage)
		if err := connection.WriteMessage(websocket.TextMessage, marshaledMessage); err != nil {
			qmlLogin.LoginDataIsValid(false)
			loger.Log.Warningf("Can not send register data. %s", err)
			qmlStatus.SendStatus("Register data can't be sent")
			return
		}
		qmlStatus.SendStatus("Waiting for server response.")
		listener.RegistrationChannel = make(chan messageService.Message)
		listener.QuitChannel = make(chan struct{})
		responseWaiterChannel = make(chan struct{})
		go listener.ListenToServer(connection)
		go channelsResolver()
	})
}
