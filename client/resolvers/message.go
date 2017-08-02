package resolvers

import (
	"encoding/json"

	"github.com/8tomat8/SSU-Golang-252-Chat/client/config"
	"github.com/8tomat8/SSU-Golang-252-Chat/client/contacts"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
)

var HaveNewMessageResponse bool
var MessageResponse messageService.ResponseFromServerBody

func ReceiveMessage(message *messageService.Message) {
	switch message.Header.Command {
	case config.GetConfig().MessageCommand.SendMessage:
		{
			AddNewMessageFromSender(message)
			break
		}
	case config.GetConfig().MessageCommand.MessageSent:
		{
			ReceiveMessageResponse(message)
			break
		}
	}
}

func AddNewMessageFromSender(message *messageService.Message) {
	//userName in Message structure is sender name for us
	senderName := message.Header.UserName
	//get sender index in contacts list
	indexOfSender := contacts.ContactsList.IndexByUserName(senderName)
	//check if this user is not blocked
	if contacts.ContactsList.ContactsList[indexOfSender].IsBlocked != true {
		var messageBody *messageService.MessageBody
		err := json.Unmarshal(message.Body, &messageBody)
		if err != nil {
			loger.Log.Errorf("Cannot unmarshal received message. %s", err)
			return
		}
		contacts.ContactsList.ContactsList[indexOfSender].MessageHistory += messageBody.Text
		// TODO: fix SendLastMessage signal
		//ui.GetQmlContacts().SendLastMessage(messageBody.Text, indexOfSender)
	}
}

func ReceiveMessageResponse(message *messageService.Message) {
	var responseBody *messageService.ResponseFromServerBody
	err := json.Unmarshal(message.Body, &responseBody)
	if err != nil {
		loger.Log.Errorf("Cannot unmarshal received message. %s", err)
		return
	}
	HaveNewMessageResponse = true
	MessageResponse = *responseBody
}

func GetMessageResponse() (bool, messageService.ResponseFromServerBody) {
	return HaveNewMessageResponse, MessageResponse
}

func SetMessageResponseReaded() {
	HaveNewMessageResponse = false
	MessageResponse = messageService.ResponseFromServerBody{}
}
