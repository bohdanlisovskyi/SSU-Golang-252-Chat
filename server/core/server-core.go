package core

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/8tomat8/SSU-Golang-252-Chat/settingService"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	conn *websocket.Conn
}

var clients = map[string]Client{}

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
				loger.Log.Warningf("Read message error: ", err.Error())
				break
			}

			msg, err := messageService.UnmarshalMessage(text)
			if err != nil {
				loger.Log.Warnf("Unmarshal message error: ", err.Error())
				continue
			}

			validateMessage(msg, messageType, conn)
		}
	}()
}

func validateMessage(message *messageService.Message, messageType int, conn *websocket.Conn) {

	if message.Header.Type_ == "" {
		loger.Log.Errorf("Message Header Type Empty")
		return
	}

	if message.Header.Type_ == "message" {
		err := sendMessage(message, messageType)
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

		return
	}

	if message.Header.Type_ == "register" {
		if _, ok := clients[message.Header.UserName]; ok {
			loger.Log.Warn("User already exist")
			return

		} else {

			clients[message.Header.UserName] = Client{conn: conn}
		}
		//run register function
		return
	}

	if message.Header.Type_ == "auth" {
		if _, ok := clients[message.Header.UserName]; ok {
			loger.Log.Warn("User already exist")
			return

		} else {

			clients[message.Header.UserName] = Client{conn: conn}
		}
		//run auth function
		return
	}

	if message.Header.Type_ == "search" {

		//run search contact function
	}

	if message.Header.Type_ == "list" {

		//run list contact function
	}

	if message.Header.Type_ == "change_pass" {

		//run change_pass function
	}

	if message.Header.Type_ == "change_info" {

		//run change_user_info function
	}

	if message.Header.Type_ == "change_aboutUser" {
		ok, err := settingService.ChangeAboutUserInfo(message)
		if err != nil {
			loger.Log.Errorf(" Error has occurred : ", err)
		}
		if !ok {
			loger.Log.Warnf(" Info about user has not changed. Please try again")
		}
		return
	}
}

func addNewConnect(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		loger.Log.Errorf("Connect new user Error: ", err.Error())
	}

	return conn, err
}

func sendMessage(message *messageService.Message, messageType int) messageService.Message {

	byteMessage, err := messageService.MarshalMessage(message)
	if err != nil {
		loger.Log.Errorf("Marshal message error: ", err.Error())
		return messageService.Message{
			Header: messageService.MessageHeader{
				Type_:   "message",
				Command: "Marshal message error",
			},
			Body: message.Body,
		}
	}

	return writeMsg(byteMessage, message, messageType) //I send this text []byte to receiver
}

func writeMsg(text []byte, message *messageService.Message, messageType int) messageService.Message {
	msgBody := messageService.MessageBody{}
	err := json.Unmarshal(message.Body, &msgBody)
	if err != nil {
		loger.Log.Errorf("Unmarshal message error: ", err.Error())

		return messageService.Message{
			Header: messageService.MessageHeader{
				Type_:   "message",
				Command: "Unmarshal message error",
			},
			Body: message.Body,
		}
	}

	client, ok := clients[msgBody.ReceiverName]
	if !ok {
		loger.Log.Warn("Receiver not found")
		return messageService.Message{
			Header: messageService.MessageHeader{
				Type_:   "message",
				Command: "Receiver not found",
			},
			Body: message.Body,
		}
	}

	err = client.conn.WriteMessage(messageType, text)
	if err != nil {
		loger.Log.Errorf("Write message error: ", err.Error())
		return messageService.Message{
			Header: messageService.MessageHeader{
				Type_:   "message",
				Command: "Write message error",
			},
			Body: message.Body,
		}
	}

	return messageService.Message{
		Header: messageService.MessageHeader{
			Type_:   "message",
			Command: "Ok",
		},
		Body: message.Body,
	}
}
