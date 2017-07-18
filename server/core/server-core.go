package core

import (
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
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

//TODO uncommented this when the message module wil be done

func MessageHandler(w http.ResponseWriter, r *http.Request) {

	/*conn := addNewClient(w, r)

	loger.Log.Infof("Add new connection: ", conn)

	go func() {

		for {
			messageType, text, err := conn.ReadMessage()

			if err != nil {

				loger.Log.Warningf("Read message error: ", err.Error())
			}

			validateMessage(messageService.UnmarshalMessage(text), messageType)
		}
	}()*/
}

/*func sendMessage(message messageService.Message, messageType int) {

	byteMessage := messageService.MarshalMessage(message)

	writeMsg(byteMessage, strconv.Itoa(message.Header.Receiver), messageType) //I send this text []byte to receiver
}

func validateMessage(message messageService.Message, messageType int) {

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
}*/

func addNewClient(w http.ResponseWriter, r *http.Request) *websocket.Conn{

	client_id := "" // get client_id from request

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {

		loger.Log.Errorf("Connect new user Error: ", err.Error())
	}

	clients[client_id] = Client{conn:conn}

	return conn
}


func writeMsg(text []byte, receiver_id string, messageType int) {

	err := clients[receiver_id].conn.WriteMessage(messageType, text)

	if err != nil {

		loger.Log.Errorf("Write message error: ", err.Error())
	}
}