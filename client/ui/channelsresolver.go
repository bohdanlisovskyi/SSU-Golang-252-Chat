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

func channelsResolver() {
	for {
		select {
		case <-listener.QuitChannel:
			return
		case msg := <-listener.AuthorizationChannel:
			switch msg.Header.Command {
			case config.GetConfig().MessageCommand.LoginIsSucc:
				loginIsSuccessfully(msg)
			default:
				loginIsNotSuccessfully(msg)
			}
		case msg := <-listener.RegistrationChannel:
			switch msg.Header.Command {
			case config.GetConfig().MessageCommand.RegisterIsSucc:
				registerIsSuccessfully(msg)
			default:
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
		case <-listener.ContactsChannel:
		}
	}
}

func receiveSettings(message messageService.Message) {

}

func loginIsSuccessfully(message messageService.Message) {
	loger.Log.Info("Open channels in loginIsSuccessfully")
	//login data is valid, so we open other channels to listen to server
	listener.MessageChannel = make(chan messageService.Message)
	listener.SettingsChannel = make(chan messageService.Message)
	listener.ContactsChannel = make(chan messageService.Message)
	loger.Log.Info("Set userinfo in loginIsSuccessfully")
	//save received username and token
	userinfo.CurrentUserInfo.UserName = message.Header.UserName
	userinfo.CurrentUserInfo.Token = message.Header.Token
	//now we just inform UI that login was successfully
	loger.Log.Infof("User %s login successfully.", message.Header.UserName)
	qmlLogin.LoginDataIsValid(true)
	qmlStatus.SendStatus("Login successfully. Waiting for contacts list.")
	sendContactsRequestToServer()
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
	sender := contacts.ContactsList.GetContactByUserName(message.Header.UserName)
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
	//TODO: change div-alignment property to display message properly
	sender.MessageHistory += messageBody.Text
	qmlContacts.SendLastMessage(messageBody.Text, contacts.ContactsList.IndexByUserName(sender.UserName))
}

func messageSent(message messageService.Message) {
	//we need to unmarshal body to get whom message was sent
	var messageBody *messageService.MessageBody
	err := json.Unmarshal(message.Body, &messageBody)
	if err != nil {
		loger.Log.Warningf("Cannot unmarshal received message. %s", err)
		return
	}
	receiver := contacts.ContactsList.GetContactByUserName(messageBody.ReceiverName)
	receiver.MessageHistory += messageBody.Text
	qmlContacts.SendLastMessage(messageBody.Text, contacts.ContactsList.IndexByUserName(receiver.UserName))
	qmlStatus.SendStatus("Message sent to " + messageBody.ReceiverName + " successfully")
	qmlMessage.MessageSent(true)
	loger.Log.Infof("Message was sent to %s .", messageBody.ReceiverName)
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
