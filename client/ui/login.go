// login
package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/gorilla/websocket"
	"github.com/Greckas/SSU-Golang-252-Chat/loger"
	"github.com/Greckas/SSU-Golang-252-Chat/client/config"
)

//binding to qml
var qmlLogin *QmlLogin
//connection will be used during all session
var connection *websocket.Conn
//represent login as QObject
type QmlLogin struct {
	core.QObject
	_ func(userName, password string) 	`slot:"checkLoginData"`
	_ func()							`slot:"logOut"`
	_ func(isLoginValid bool)           `signal:"loginDataIsValid"`
}

func initQmlLogin(quickWidget *quick.QQuickWidget) {
	qmlLogin = NewQmlLogin(nil)
	quickWidget.RootContext().SetContextProperty("qmlLogin", qmlLogin)
	// TODO: move validation/connection slot to separate function
	qmlLogin.ConnectCheckLoginData(func(userName, password string) {
		if userName != "" && password != "" {
			//err has separate declarations
			// because in next statement we need to use global "connection" variable
			//instead declarate new one
			var err error
			connection, _, err = websocket.DefaultDialer.Dial(config.GetConfig().Server.Connection, nil)
			if err != nil {
				loger.Log.Errorf("Can not dial. %s", err)
				qmlLogin.LoginDataIsValid(false)
				qmlStatus.SendStatus("Loging failed")
				return
			}
			//there we need to send to server loging data
			//
			loger.Log.Info("Connected to server")
			qmlStatus.SendStatus("Login successfull")
			qmlLogin.LoginDataIsValid(true)
		} else {
			qmlLogin.LoginDataIsValid(false)
		}
	})
	qmlLogin.ConnectLogOut(func(){
		loger.Log.Info("Connection closed")
		connection.Close()
	})
}