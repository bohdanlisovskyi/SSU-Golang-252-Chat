package modules

import (
	"encoding/json"

	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/auth"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/customers"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/message"
	"github.com/8tomat8/SSU-Golang-252-Chat/settingService"
	"github.com/gorilla/websocket"
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

	errMessage := coremessage.SendMessage(message, messageType)
	byteError, er := json.Marshal(errMessage)

	if er != nil {
		loger.Log.Errorf("Marshal Error Message")
	}

	if errMessage.Header.Command != "Ok" {
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
	us, tok, err := auth.Login(user.UserName, user.Password)
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