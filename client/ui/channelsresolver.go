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
		case <-listener.SettingsChannel:
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
		connection.Close()
		loger.Log.Infof("User %s login wasn't successfully.", message.Header.UserName)
		qmlLogin.LoginDataIsValid(false)
		qmlStatus.SendStatus("Loging was unsuccessfully. Please check username or password again.")
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
		sender.MessageHistory += messageBody.Text
		qmlContacts.SendLastMessage(messageBody.Text, contacts.ContactsList.IndexByUserName(sender.UserName))
	}
}
