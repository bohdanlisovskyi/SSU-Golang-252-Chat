package core

import (
	"strconv"
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/8tomat8/SSU-Golang-252-Chat/chat-log"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	conn *websocket.Conn
	message chan []byte
}

var clients = map[string]Client{}

func MessageHandler(w http.ResponseWriter, r *http.Request) {

	conn := addNewClient(w, r)

	chatlog.Log(chatlog.Fields{}, "info", "Add new connection: ", conn)

	go func() {

		for {
			messageType, text, err := conn.ReadMessage()

			chatlog.HandleError(err, "Read message error: ")

			validateMessage(messageService.UnmarshalMessage(text), messageType)
		}
	}()
}

func sendMessage(message messageService.Message, messageType int) {

	byteMessage := messageService.MarshalMessage(message)

	writeMsg(byteMessage, strconv.Itoa(message.Header.Receiver), messageType) //I send this text []byte to receiver
}

func validateMessage(message messageService.Message, messageType int) {

	//Why i don't use case? because if working faster

	if message.Header.Type_ == "msg" {

		sendMessage(message, messageType)
	}

	if message.Header.Type_ == "register" {

		//run register function
	}

	if message.Header.Type_ == "auth" {

		//run auth function
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
}

func addNewClient(w http.ResponseWriter, r *http.Request) *websocket.Conn{

	conn, err := upgrader.Upgrade(w, r, nil)

	chatlog.HandleError(err, "Connect new user Error: ")

	clients[strconv.Itoa(len(clients)+1)] = Client{conn:conn}

	return conn
}


func writeMsg(text []byte, receiver_id string, messageType int) {

	chatlog.HandleError(clients[receiver_id].conn.WriteMessage(messageType, text), "Write message error: ")
}