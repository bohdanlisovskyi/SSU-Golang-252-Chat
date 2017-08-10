package modules

import (
	"encoding/json"

	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/auth"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/customers"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/message"
	"github.com/gorilla/websocket"
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
		conn.Close()
		return
	}
	var user *messageService.Authentification
	err := json.Unmarshal(message.Body, &user)
	if err != nil {
		loger.Log.Warn("failed to unmarshal body")
		conn.Close()
		return
	}
	us, tok, err := auth.RegisterNewUser(user)
	if err != nil {
		loger.Log.Errorf("failed to register user", err)
		conn.Close()
		return
	}
	newMessageHeader := messageService.MessageHeader{
		Type_:    "authorization",
		Command:  "registrissucc",
		UserName: us.UserName,
		Token:    tok,
	}
	newMessage := messageService.Message{
		Header: newMessageHeader,
	}
	marshaledMessage, err := messageService.MarshalMessage(&newMessage)
	if err != nil {
		loger.Log.Errorf("Can`t marshal message. %s", err)
		conn.Close()
		return
	}
	if err := conn.WriteMessage(websocket.TextMessage, marshaledMessage); err != nil {
		loger.Log.Errorf("Can not send message. %s", err)
		conn.Close()
		return
	}
	customers.Clients[message.Header.UserName] = customers.Client{Conn: conn, Token: tok}
}

func Auth(message *messageService.Message, conn *websocket.Conn) {
	if _, ok := customers.Clients[message.Header.UserName]; ok {
		loger.Log.Warn("User already exist")
		conn.Close()
		return
	}
	var user *messageService.Authentification
	err := json.Unmarshal(message.Body, &user)
	if err != nil {
		loger.Log.Warn("failed to unmarshal body")
		conn.Close()
		return
	}
	us, tok, err := auth.RegisterNewUser(user)
	if err != nil {
		loger.Log.Errorf("failed to register user", err)
		conn.Close()
		return
	}
	newMessageHeader := messageService.MessageHeader{
		Type_:    "authorization",
		Command:  "authissucc",
		UserName: us.UserName,
		Token:    tok,
	}
	newMessage := messageService.Message{
		Header: newMessageHeader,
	}
	marshaledMessage, err := messageService.MarshalMessage(&newMessage)
	if err != nil {
		loger.Log.Errorf("Can`t marshal message. %s", err)
		conn.Close()
		return
	}
	if err := conn.WriteMessage(websocket.TextMessage, marshaledMessage); err != nil {
		loger.Log.Errorf("Can not send message. %s", err)
		conn.Close()
		return
	}
	customers.Clients[message.Header.UserName] = customers.Client{Conn: conn, Token: tok}
}
