package ui

import (
	"encoding/json"

	"github.com/8tomat8/SSU-Golang-252-Chat/client/config"
	"github.com/8tomat8/SSU-Golang-252-Chat/client/contacts"
	"github.com/8tomat8/SSU-Golang-252-Chat/client/listener"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/8tomat8/SSU-Golang-252-Chat/userinfo"
)

var responseWaiterChannel chan struct{}
var pingerChannel chan struct{}

func channelsResolver() {
	for {
		select {
		case <-listener.QuitChannel:
			close(pingerChannel)
			return
		case msg := <-listener.AuthorizationChannel:
			//close this channel to end response-timer
			close(responseWaiterChannel)
			if msg.Header.Command == config.GetConfig().MessageCommand.LoginIsSucc {
				loginIsSuccessfully(msg)
			} else {
				loginIsNotSuccessfully(msg)
			}

		case msg := <-listener.RegistrationChannel:
			close(responseWaiterChannel)
			if msg.Header.Command == config.GetConfig().MessageCommand.RegisterIsSucc {
				registerIsSuccessfully(msg)
			} else {
				registerIsNotSuccessfully(msg)
			}
		case msg := <-listener.MessageChannel:
			switch msg.Header.Command {
			case config.GetConfig().MessageCommand.ReceiveMessage:
				receiveNewMessage(msg)
			case config.GetConfig().MessageCommand.MessageSent:
				messageSent(msg)
			default:
				messageWasntSent(msg)
			}
		case msg := <-listener.SettingsChannel:
			receiveSettings(msg)
		case msg := <-listener.ContactsChannel:
			receiveContacts(msg)
		}
	}
}

func receiveSettings(message messageService.Message) {

}

func receiveContacts(message messageService.Message) {
	err := json.Unmarshal(message.Body, contacts.ContactsList)
	if err != nil {
		loger.Log.Warningf("Can not unmarshal list of contacts. %s", err)
		qmlStatus.SendStatus("Contacts list didn't download.")
		//TODO: send contacts request again
		return
	}
	updateContactsList()
	qmlStatus.SendStatus("Contacts list updated")
}

func loginIsSuccessfully(message messageService.Message) {
	//login data is valid, so we open other channels to listen to server
	listener.MessageChannel = make(chan messageService.Message)
	listener.SettingsChannel = make(chan messageService.Message)
	listener.ContactsChannel = make(chan messageService.Message)
	//save received username and token
	userinfo.CurrentUserInfo.UserName = message.Header.UserName
	userinfo.CurrentUserInfo.Token = message.Header.Token
	//TODO: send request every n seconds until we get response
	//now we just inform UI that login was successfully and send contacts request
	if err := sendContactsRequestToServer(); err != nil {
		//it's no need to print error message,
		//it has already printed in sendContactsRequestToServer function
		qmlLogin.LoginDataIsValid(false)
		return
	}
	loger.Log.Infof("User %s login successfully.", message.Header.UserName)
	qmlLogin.LoginDataIsValid(true)
	qmlStatus.SendStatus("Login successfully. Waiting for contacts list.")
	//test contacts. will be removed after implement Contacts service
	contacts.ContactsList.ContactsList = append(contacts.ContactsList.ContactsList, messageService.ClientContact{UserName: "aome1", NickName: "1", IsBlocked: 1})
	contacts.ContactsList.ContactsList = append(contacts.ContactsList.ContactsList, messageService.ClientContact{UserName: "aome2", NickName: "2", IsBlocked: 1})
	contacts.ContactsList.ContactsList = append(contacts.ContactsList.ContactsList, messageService.ClientContact{UserName: "aome3", NickName: "3", IsBlocked: 0})
	contacts.ContactsList.ContactsList = append(contacts.ContactsList.ContactsList, messageService.ClientContact{UserName: "aome4", NickName: "4", IsBlocked: 1})
	updateContactsList()
}

func loginIsNotSuccessfully(message messageService.Message) {
	//just close connection. in listener after closing connection excutes QuitChannel <- struct{}{}
	err := connection.Close()
	if err != nil {
		loger.Log.Warningf("Cannot close connection after login is not successfully. %s", err)
		//for client it doesn`t matter. connection can`t close == login is not successfully
	}
	//in this case message.Header.Command contains error message
	loger.Log.Infof("User %s login wasn't successfully. %s", message.Header.UserName, message.Header.Command)
	qmlLogin.LoginDataIsValid(false)
	qmlStatus.SendStatus("Loging was unsuccessfully. Please check username or password again.")
}

func registerIsSuccessfully(message messageService.Message) {
	//however register successfully we need to close connection
	//because all register actions have ended
	err := connection.Close()
	if err != nil {
		loger.Log.Errorf("Cannot close connection after register is successfully. %s", err)
	}
	loger.Log.Infof("User %s register successfully.", message.Header.UserName)
	qmlRegister.RegisterDataIsValid(true)
	qmlStatus.SendStatus("Register successfully. One more step and you`ll be happy!")
}

func registerIsNotSuccessfully(message messageService.Message) {
	//register is not successfully we need to close connection
	err := connection.Close()
	if err != nil {
		loger.Log.Errorf("Cannot close connection after register is not successfully. %s", err)
		//for client it doesn`t matter. connection can`t close == register is not successfully
	}
	//in this case message.Header.Command contains error message
	loger.Log.Infof("User %s register not successfully. %s", message.Header.UserName, message.Header.Command)
	qmlRegister.RegisterDataIsValid(false)
	qmlStatus.SendStatus("Register not successfully. Please, try again later!")
}

func receiveNewMessage(message messageService.Message) {
	//sender := contacts.ContactsList.GetContactByUserName(message.Header.UserName)
	var sender *ContactObject
	for i := 0; i < listOfContacts.Size(); i++ {
		var iData, exists = listOfContacts.Get(i)
		if !exists {
			return
		}
		var data = iData.(*ContactObject)
		if data.UserName == message.Header.UserName {
			sender = data
			break
		}
	}
	if sender == nil {
		loger.Log.Infof("Received message from user not in list. %s", message.Header.UserName)
		return
	}
	if sender.IsBlocked {
		loger.Log.Infof("Received message from blocked user - %s", sender.UserName)
		return
	}
	var messageBody *messageService.MessageBody
	err := json.Unmarshal(message.Body, &messageBody)
	if err != nil {
		loger.Log.Warningf("Cannot unmarshal received message. %s", err)
		return
	}
	senderIndex := addToHistory(message.Header.UserName, messageBody.Text)
	if senderIndex != -1 {
		qmlContacts.SendLastMessage(messageBody.Text, senderIndex)
		qmlStatus.SendStatus("Received message from " + message.Header.UserName)
		loger.Log.Infof("Message received from %s .", message.Header.UserName)
		return
	}
	loger.Log.Infof("Message received from %s . But he is not in your list.", message.Header.UserName)
	return
}

func messageSent(message messageService.Message) {
	//we need to unmarshal body to get whom message was sent
	var messageBody *messageService.MessageBody
	err := json.Unmarshal(message.Body, &messageBody)
	if err != nil {
		loger.Log.Warningf("Cannot unmarshal received message. %s", err)
		return
	}
	receiverIndex := addToHistory(messageBody.ReceiverName, messageBody.Text)
	if receiverIndex != -1 {
		qmlContacts.SendLastMessage(messageBody.Text, receiverIndex)
		qmlStatus.SendStatus("Message sent to " + messageBody.ReceiverName + " successfully")
		qmlMessage.MessageSent(true)
		loger.Log.Infof("Message was sent to %s .", messageBody.ReceiverName)
		return
	}
	qmlStatus.SendStatus("Message sent to " + messageBody.ReceiverName + ", but now he's offline.")
	qmlMessage.MessageSent(true)
	loger.Log.Infof("Message was sent to offline contact %s .", messageBody.ReceiverName)
}

func messageWasntSent(message messageService.Message) {
	//we need to unmarshal body to get whom message was sent
	var messageBody *messageService.MessageBody
	err := json.Unmarshal(message.Body, &messageBody)
	if err != nil {
		loger.Log.Warningf("Cannot unmarshal received message. %s", err)
		return
	}
	qmlStatus.SendStatus("Message wasn't sent to " + messageBody.ReceiverName)
	qmlMessage.MessageSent(true)
	//in this case message.Header.Command contains error message
	loger.Log.Infof("Message wasn't sent to %s . %s", messageBody.ReceiverName, message.Header.Command)
}
