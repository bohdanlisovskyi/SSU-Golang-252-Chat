package core

import (
	"net/http"

	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/config"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/modules"
	"github.com/Greckas/SSU-Golang-252-Chat/server/message"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := addNewConnect(w, r)

	if err != nil {
		loger.Log.Error("Add new connection error: %s", err)
		return
	}

	go func() {
		defer conn.Close()
		for {
			messageType, text, err := conn.ReadMessage()
			if err != nil {
				loger.Log.Warningf("Read message error: %s", err.Error())
				break
			}

			msg, err := messageService.UnmarshalMessage(text)
			if err != nil {
				loger.Log.Warnf("Unmarshal message error: %s", err.Error())
				continue
			}

			validateMessage(msg, messageType, conn)
		}
	}()
}

func validateMessage(message *messageService.Message, messageType int, conn *websocket.Conn) {

	switch message.Header.Type_ {
	case coremessage.EmptyType:
		modules.EmptyType()
	case coremessage.MessageType:
		modules.Message(message, messageType, conn)
	case coremessage.RegisterType:
		modules.Register(message, conn)
	case coremessage.AuthType:
		modules.Auth(message, conn)
	}
}

func addNewConnect(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		loger.Log.Errorf("Connect new user Error: %s", err.Error())
	}

	return conn, err
}

func ReturnPort() string {

	port := config.GetConfig().Server.Port
	if port == "" {
		loger.Log.Panic("Port is Empty")
	}

	return port
}
