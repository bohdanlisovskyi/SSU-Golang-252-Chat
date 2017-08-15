package modules

import (
	"encoding/json"

	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/customers"
	"github.com/gorilla/websocket"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/message"
)

func EmptyType() {
	loger.Log.Errorf("Message Header Type Empty")
}

func Message(message *messageService.Message, messageType int, conn *websocket.Conn) {

	err := coremessage.SendMessage(message, messageType)
	byteError, er := json.Marshal(err)

	if er != nil {
		loger.Log.Errorf("Marshal Error Message")
	}

	if err.Header.Command != "Ok" {
		err := conn.WriteMessage(messageType, byteError)

		if err != nil {
			loger.Log.Errorf("Write Message Error")
		}
	}
}

func Register(message *messageService.Message, conn *websocket.Conn) {

	if _, ok := customers.Clients[message.Header.UserName]; ok {
		loger.Log.Warn("User already exist")
		return

	}
	customers.Clients[message.Header.UserName] = customers.Client{Conn: conn}
}

func Auth(message *messageService.Message, conn *websocket.Conn) {

	if _, ok := customers.Clients[message.Header.UserName]; ok {
		loger.Log.Warn("User already exist")
		return

	}
	customers.Clients[message.Header.UserName] = customers.Client{Conn: conn}
}
