// login
package ui

import (
	"net/url"

	"encoding/json"

	"time"

	"github.com/8tomat8/SSU-Golang-252-Chat/client/config"
	"github.com/8tomat8/SSU-Golang-252-Chat/client/listener"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/8tomat8/SSU-Golang-252-Chat/userinfo"
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
	_ func()                          `signal:"startBusyIndicator"`
}

func initQmlLogin(quickWidget *quick.QQuickWidget) {
	qmlLogin = NewQmlLogin(nil)
	quickWidget.RootContext().SetContextProperty("qmlLogin", qmlLogin)
	qmlLogin.ConnectCheckLoginData(checkLoginDataAndConnect)
	qmlLogin.ConnectLogOut(logOut)
}

func checkLoginDataAndConnect(userName, password string) {
	if userName == "" || password == "" {
		loger.Log.Info("Login failed. userName and/or password field are empty")
		qmlStatus.SendStatus("Please, fill the fields")
		qmlLogin.LoginDataIsValid(false)
		return
	}
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
		loger.Log.Warningf("Can not dial. %s", err)
		qmlLogin.LoginDataIsValid(false)
		qmlStatus.SendStatus("Loging failed")
		return
	}
	loger.Log.Infof("Connection created. Server url - %s", connUrl.String())
	//there we need to send to server loging data
	newMessageHeader := messageService.MessageHeader{
		Type_:    config.GetConfig().MessageType.Auth,
		Command:  config.GetConfig().MessageCommand.SendLoginData,
		UserName: "",
		Token:    "",
	}
	newMessageBody := messageService.Authentification{
		UserName: userName,
		NickName: "",
		Password: password,
	}
	newRawMessageBody, err := json.Marshal(newMessageBody)
	if err != nil {
		qmlLogin.LoginDataIsValid(false)
		loger.Log.Warningf("Can`t marshal login message. %s", err)
		qmlStatus.SendStatus("Login data can't be sent")
		return
	}
	newMessage := messageService.Message{
		Header: newMessageHeader,
		Body:   newRawMessageBody,
	}
	marshaledMessage, err := messageService.MarshalMessage(&newMessage)
	if err := connection.WriteMessage(websocket.TextMessage, marshaledMessage); err != nil {
		qmlLogin.LoginDataIsValid(false)
		loger.Log.Warningf("Can not send login data. %s", err)
		qmlStatus.SendStatus("Login data can't be sent")
		return
	}
	qmlStatus.SendStatus("Waiting for server response.")
	loger.Log.Info("Login data sent to server. Waiting for response.")
	//start listening for login response
	//if all right - we need to start listening other channels
	listener.AuthorizationChannel = make(chan messageService.Message)
	listener.QuitChannel = make(chan struct{})
	responseWaiterChannel = make(chan struct{})
	userinfo.CurrentUserInfo = &userinfo.UserInfo{}
	go listener.ListenToServer(connection)
	go channelsResolver()
	go loginResponseWaiter(8)
}

func logOut() {
	err := connection.Close()
	if err != nil {
		loger.Log.Warningf("Can not close connection. %s", err)
		//if it`s not already closed - we have no idea what happened
		if !websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseAbnormalClosure) {
			qmlStatus.SendStatus("Can not log out")
			qmlLogin.LogOutIsSuccessfull(false)
			return
		}
		//if it`s closed, error spawns, but connection in stay we want
		//so just signal UI to change window
	}
	loger.Log.Info("Connection closed")
	qmlLogin.LogOutIsSuccessfull(true)
}

func loginResponseWaiter(seconds time.Duration) {
	timer := time.NewTimer(time.Second * seconds)
	qmlLogin.StartBusyIndicator()
	for {
		select {
		case <-timer.C:
			//we haven't got response from server in time
			qmlLogin.LoginDataIsValid(false)
			qmlStatus.SendStatus("Server doesn't response. Pleas, try again later.")
			connection.Close()
		case <-responseWaiterChannel:
			//we have got response in time
			return
		}
	}
}
