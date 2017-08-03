// login
package ui

import (
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
var qmlLogin *QmlLogin

//connection will be used during all session
var connection *websocket.Conn

//represent login as QObject
type QmlLogin struct {
	core.QObject
	_ func(userName, password string) `slot:"checkLoginData"`
	_ func()                          `slot:"logOut"`
	_ func(isLoginValid bool)         `signal:"loginDataIsValid"`
	_ func(isLogOutSuccessfull bool)  `signal:"logOutIsSuccessfull"`
}

func initQmlLogin(quickWidget *quick.QQuickWidget) {
	qmlLogin = NewQmlLogin(nil)
	quickWidget.RootContext().SetContextProperty("qmlLogin", qmlLogin)
	qmlLogin.ConnectCheckLoginData(checkLoginDataAndConnect)
	qmlLogin.ConnectLogOut(logOut)
}

func checkLoginDataAndConnect(userName, password string) {
	if userName != "" && password != "" {
		//err has separate declarations
		// because in next statement we need to use global "connection" variable
		//instead declarate new one
		var err error
		//another separate declaration/definition for connUrl
		//because compiler emit error for declaration->definition->method_call in one statement
		connUrl := url.URL{
			Scheme: "ws",
			Host:   config.GetConfig().Server.Host,
			Path:   config.GetConfig().Server.Path,
		}
		connection, _, err = websocket.DefaultDialer.Dial(connUrl.String(), nil)
		if err != nil {
			loger.Log.Errorf("Can not dial. %s", err)
			qmlLogin.LoginDataIsValid(false)
			qmlStatus.SendStatus("Loging failed")
			return
		}
		//there we need to send to server loging data
		//
		qmlStatus.SendStatus("Waiting for server response.")
		//after get connection, we need to listen for all responses
		listener.LoginStatusChannel = make(chan messageService.ResponseFromServerBody)
		listener.MessageSentStatusChannel = make(chan messageService.ResponseFromServerBody)
		listener.RegisterStatusChannel = make(chan messageService.ResponseFromServerBody)
		listener.QuitChannel = make(chan bool)
		go listener.ListenToServer(connection)
		go channelsResolver()
		//go CheckForMessageResponse()
	} else {
		qmlStatus.SendStatus("Please, fill the fields")
		qmlLogin.LoginDataIsValid(false)
	}
}

func logOut() {
	err := connection.Close()
	if err != nil {
		loger.Log.Errorf("Can not close connection. %s", err)
		qmlStatus.SendStatus("Can not log out")
		qmlLogin.LogOutIsSuccessfull(false)
		return
	}
	loger.Log.Info("Connection closed")
	qmlLogin.LogOutIsSuccessfull(true)
}
