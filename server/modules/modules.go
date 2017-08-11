package modules

import (
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"encoding/json"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/gorilla/websocket"

	"github.com/8tomat8/SSU-Golang-252-Chat/server/customers"
	"github.com/8tomat8/SSU-Golang-252-Chat/settingService"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/message"
)

func EmptyType() {
	loger.Log.Errorf("Message Header Type Empty")
}

func Message(message *messageService.Message, messageType int, conn *websocket.Conn) {

	isBlocked, err := settingService.IsUserBlocked(message)
	if err != nil {
		loger.Log.Errorf(" Error has occurred : ", err)
		byteError, _ := json.Marshal(changeNicknameError) // this for UI
		conn.WriteMessage(messageType, byteError)         // this for UI
		return
	}
	if isBlocked {
		loger.Log.Errorf("This contact is blocked")
		return
	}

	err = coremessage.SendMessage(message, messageType)
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

func ChangePass(message *messageService.Message, messageType int, conn *websocket.Conn) {
	ok, err := settingService.ChangePass(message)
	if err != nil {
		loger.Log.Errorf(" Error has occurred : ", err)
		byteError, _ := json.Marshal(changePassError) // this for UI
		conn.WriteMessage(messageType, byteError)     // this for UI
		return
	}
	if !ok {
		loger.Log.Warnf(" Password has not changed. Please try again")
		byteError, _ := json.Marshal(changePassError) // this for UI
		conn.WriteMessage(messageType, byteError)     // this for UI
		return
	}
	byteOk, _ := json.Marshal(changePassOk) // this for UI
	conn.WriteMessage(messageType, byteOk)  // this for UI
	return
}

func ChangeNickName(message *messageService.Message, messageType int, conn *websocket.Conn) {
	ok, err := settingService.ChangeNickName(message)
	if err != nil {
		loger.Log.Errorf(" Error has occurred : ", err)
		byteError, _ := json.Marshal(changeNicknameError) // this for UI
		conn.WriteMessage(messageType, byteError)         // this for UI
		return
	}
	if !ok {
		loger.Log.Warnf(" Nickname has not changed. Please try again")
		byteError, _ := json.Marshal(changeNicknameError) // this for UI
		conn.WriteMessage(messageType, byteError)         // this for UI
		return
	}
	byteOk, _ := json.Marshal(changeNicknameOk) // this for UI
	conn.WriteMessage(messageType, byteOk)      // this for UI
	return
}

func ChangeBirthday(message *messageService.Message, messageType int, conn *websocket.Conn) {
	ok, err := settingService.ChangeBirthday(message)
	if err != nil {
		loger.Log.Errorf(" Error has occurred : ", err)
		byteError, _ := json.Marshal(changeBirthdayError) // this for UI
		conn.WriteMessage(messageType, byteError)         // this for UI
		return
	}
	if !ok {
		loger.Log.Warnf(" Birthday has not changed. Please try again")
		byteError, _ := json.Marshal(changeBirthdayError) // this for UI
		conn.WriteMessage(messageType, byteError)         // this for UI
		return
	}
	byteOk, _ := json.Marshal(changeBirthdayOk) // this for UI
	conn.WriteMessage(messageType, byteOk)      // this for UI
	return
}

func ChangeAboutUserInfo(message *messageService.Message, messageType int, conn *websocket.Conn) {
	ok, err := settingService.ChangeAboutUserInfo(message)
	if err != nil {
		loger.Log.Errorf(" Error has occurred : ", err)
		byteError, _ := json.Marshal(changeUserInfoError) // this for UI
		conn.WriteMessage(messageType, byteError)         // this for UI
		return
	}
	if !ok {
		loger.Log.Warnf(" Info about user has not changed. Please try again")
		byteOk, _ := json.Marshal(changeUserInfoOk) // this for UI
		conn.WriteMessage(messageType, byteOk)      // this for UI
		return
	}
	return
}

func BlockUnblockUser(message *messageService.Message, messageType int, conn *websocket.Conn) {
	ok, err := settingService.BlockUnblockUser(message)
	if err != nil {
		loger.Log.Errorf(" Error has occurred : ", err)
		byteError, _ := json.Marshal(blockUserError) // this for UI
		conn.WriteMessage(messageType, byteError)    // this for UI
		return
	}
	if !ok {
		loger.Log.Warnf(" Error. Please try again")
		byteOk, _ := json.Marshal(blockUserOk) // this for UI
		conn.WriteMessage(messageType, byteOk) // this for UI
		return
	}
	return
}
