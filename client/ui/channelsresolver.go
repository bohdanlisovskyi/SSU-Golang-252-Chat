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
			receiveAuth(msg)
		case msg := <-listener.MessageChannel:
			receiveMessage(msg)
		case msg := <-listener.SettingsChannel:
			receiveSettings(msg)
		case <-listener.ContactsChannel:
		}
	}
}

func receiveAuth(message messageService.Message) {
	switch message.Header.Command {
	case config.GetConfig().MessageCommand.LoginIsSucc:
		//login data is valid, so we open other channels to listen to server
		listener.MessageChannel = make(chan messageService.Message)
		listener.SettingsChannel = make(chan messageService.Message)
		listener.ContactsChannel = make(chan messageService.Message)
		//save received username and token
		userinfo.CurrentUserInfo.UserName = message.Header.UserName
		userinfo.CurrentUserInfo.UserName = message.Header.Token
		//now we just inform UI that login was successfully
		loger.Log.Infof("User %s login successfully.", message.Header.UserName)
		qmlLogin.LoginDataIsValid(true)
		qmlStatus.SendStatus("Loging successfully. Be happy and write messages :)")
	case config.GetConfig().MessageCommand.LoginIsNotSucc:
		//login data is invalid, so we close conn and inform UI
		listener.QuitChannel <- struct{}{}
		//we blocked Auth channel until user want to send login data again
		listener.AuthorizationChannel = nil
		//just close connection
		err := connection.Close()
		if err != nil {
			loger.Log.Errorf("Cannot close connection after login is not successfully. %s", err)
			//for client it doesn`t matter. connection can`t close == login is not successfully
		}
		loger.Log.Infof("User %s login wasn't successfully.", message.Header.UserName)
		qmlLogin.LoginDataIsValid(false)
		qmlStatus.SendStatus("Loging was unsuccessfully. Please check username or password again.")
	case config.GetConfig().MessageCommand.RegisterIsSucc:
		//however register successfully we need to close connection
		//because all register actions have ended
		listener.QuitChannel <- struct{}{}
		//we blocked Auth channel until user want to send Auth data again
		listener.AuthorizationChannel = nil
		err := connection.Close()
		if err != nil {
			loger.Log.Errorf("Cannot close connection after register is successfully. %s", err)
		}
		loger.Log.Infof("User %s register successfully.", message.Header.UserName)
		qmlRegister.RegisterDataIsValid(true)
		qmlStatus.SendStatus("Register successfully. One more step and you`ll be happy!")
	case config.GetConfig().MessageCommand.RegisterIsNotSucc:
		//register is not successfully we need to close connection
		listener.QuitChannel <- struct{}{}
		//we blocked Auth channel until user want to send Auth data again
		listener.AuthorizationChannel = nil
		err := connection.Close()
		if err != nil {
			loger.Log.Errorf("Cannot close connection after register is not successfully. %s", err)
			//for client it doesn`t matter. connection can`t close == register is not successfully
		}
		loger.Log.Infof("User %s register not successfully.", message.Header.UserName)
		qmlRegister.RegisterDataIsValid(false)
		qmlStatus.SendStatus("Register not successfully. Please, try again later!")
	}
}

func receiveMessage(message messageService.Message) {
	switch message.Header.Command {
	case config.GetConfig().MessageCommand.ReceiveMessage:
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
	case config.GetConfig().MessageCommand.MessageSent:
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
	case config.GetConfig().MessageCommand.MessageWasntSent:
		//we need to unmarshal body to get whom message was sent
		var messageBody *messageService.MessageBody
		err := json.Unmarshal(message.Body, &messageBody)
		if err != nil {
			loger.Log.Warningf("Cannot unmarshal received message. %s", err)
			return
		}
		qmlStatus.SendStatus("Message wasn't sent to " + messageBody.ReceiverName)
		qmlMessage.MessageSent(true)
		loger.Log.Infof("Message wasn't sent to %s .", messageBody.ReceiverName)
	}
}

func receiveSettings(message messageService.Message) {

}
