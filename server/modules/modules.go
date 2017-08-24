package modules

import (
	"encoding/json"

	"errors"

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
		var messageError = messageService.Message{
			Header:
			messageService.MessageHeader{
				Type_:   coremessage.MessageType,
				Command: "Message has not been sent. Please try again",
			},
			Body: message.Body,
		}
		byteError, _ := json.Marshal(messageError) // this for UI
		conn.WriteMessage(messageType, byteError)  // this for UI
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
func sendMsg(conn *websocket.Conn, newHeader messageService.MessageHeader, newBody json.RawMessage) {
	newMessage := messageService.Message{
		Header: newHeader,
		Body:   newBody,
	}
	marshaledMessage, err := messageService.MarshalMessage(&newMessage)
	if err != nil {
		loger.Log.Infof("Can`t marshal message. %s", err)
		return
	}
	if err := conn.WriteMessage(websocket.TextMessage, marshaledMessage); err != nil {
		loger.Log.Infof("Can not send message. %s", err)
		conn.Close()
		return
	}

}
func Register(message *messageService.Message, conn *websocket.Conn) {
	if _, ok := customers.Clients[message.Header.UserName]; ok {
		err := errors.New("User already exist")
		newMessageHeader := messageService.MessageHeader{
			Type_:   coremessage.RegisterType,
			Command: err.Error(),
		}
		sendMsg(conn, newMessageHeader, nil)
		conn.Close()
		return
	}
	var user *messageService.Authentification
	err := json.Unmarshal(message.Body, &user)
	if err != nil {
		loger.Log.Infof(" failed to unmarshal body : ", err)
		newMessageHeader := messageService.MessageHeader{
			Type_:   coremessage.RegisterType,
			Command: err.Error(),
		}
		newMessageBody := messageService.Authentification{}
		newRawMessageBody, err := json.Marshal(newMessageBody)
		if err != nil {
			loger.Log.Warningf("Can`t marshal login message. %s", err)
			return
		}
		sendMsg(conn, newMessageHeader, newRawMessageBody)
		conn.Close()
		return
	}
	us, tok, err := auth.RegisterNewUser(user)
	if err != nil {
		loger.Log.Infof("failed to register user", err)
		newMessageHeader := messageService.MessageHeader{
			Type_:   coremessage.RegisterType,
			Command: err.Error(),
		}
		newMessageBody := messageService.Authentification{}
		newRawMessageBody, err := json.Marshal(newMessageBody)
		if err != nil {
			loger.Log.Warningf("Can`t marshal login message. %s", err)
			return
		}
		sendMsg(conn, newMessageHeader, newRawMessageBody)
		conn.Close()
		return
	}
	newMessageHeader := messageService.MessageHeader{
		Type_:    coremessage.RegisterType,
		Command:  coremessage.RegisterSuccComm,
		UserName: us.UserName,
		Token:    tok,
	}
	newMessageBody := messageService.Authentification{}
	newRawMessageBody, err := json.Marshal(newMessageBody)
	if err != nil {
		loger.Log.Warningf("Can`t marshal login message. %s", err)
		return
	}
	sendMsg(conn, newMessageHeader, newRawMessageBody)
}

func Auth(message *messageService.Message, conn *websocket.Conn) {
	if _, ok := customers.Clients[message.Header.UserName]; ok {
		err := errors.New("User already loged in")
		newMessageHeader := messageService.MessageHeader{
			Type_:   coremessage.AuthType,
			Command: err.Error(),
		}
		newMessageBody := messageService.Authentification{}
		newRawMessageBody, err := json.Marshal(newMessageBody)
		if err != nil {
			loger.Log.Warningf("Can`t marshal login message. %s", err)
			return
		}
		sendMsg(conn, newMessageHeader, newRawMessageBody)
		conn.Close()
		return
	}
	var user *messageService.Authentification
	err := json.Unmarshal(message.Body, &user)
	if err != nil {
		loger.Log.Infof(" failed to unmarshal body : ", err)
		newMessageHeader := messageService.MessageHeader{
			Type_:   coremessage.AuthType,
			Command: err.Error(),
		}
		newMessageBody := messageService.Authentification{}
		newRawMessageBody, err := json.Marshal(newMessageBody)
		if err != nil {
			loger.Log.Warningf("Can`t marshal login message. %s", err)
			return
		}
		sendMsg(conn, newMessageHeader, newRawMessageBody)
		conn.Close()
		return
	}
	us, tok, err := auth.Login(user.UserName, user.Password)
	if err != nil {
		loger.Log.Infof("failed to login user", err)
		newMessageHeader := messageService.MessageHeader{
			Type_:   coremessage.AuthType,
			Command: err.Error(),
		}
		newMessageBody := messageService.Authentification{}
		newRawMessageBody, err := json.Marshal(newMessageBody)
		if err != nil {
			loger.Log.Warningf("Can`t marshal login message. %s", err)
			return
		}
		sendMsg(conn, newMessageHeader, newRawMessageBody)
		conn.Close()
		return
	}
	newMessageHeader := messageService.MessageHeader{
		Type_:    coremessage.AuthType,
		Command:  coremessage.AuthSuccComm,
		UserName: us.UserName,
		Token:    tok,
	}
	nick, err := auth.SendNickName(user)
	if err != nil {
		loger.Log.Warningf("Can`t get user nick_name. %s", err)
		return
	}
	newMessageBody := messageService.Authentification{
		NickName: nick,
	}
	newRawMessageBody, err := json.Marshal(newMessageBody)
	if err != nil {
		loger.Log.Warningf("Can`t marshal login message. %s", err)
		return
	}
	customers.Clients[user.UserName] = customers.Client{Conn: conn, Token: tok}
	sendMsg(conn, newMessageHeader, newRawMessageBody)
}

func SendContacts(message *messageService.Message, conn *websocket.Conn) {
	ok, err := auth.VerifyToken(message, customers.Clients[message.Header.UserName])
	if err != nil {
		loger.Log.Warningf("Error in token verification! %s", err)
		return
	}
	if !ok {
		loger.Log.Warningf("token verification failed! %s", err)
		return
	}
	cont, err := auth.SendContacts(message)
	if err != nil {
		loger.Log.Warningf("Failed to send user contacts %s", err)
		return
	}
	newMessageHeader := messageService.MessageHeader{
		Type_:   coremessage.ContactsType,
		Command: coremessage.ContactsRequest,
	}
	newMessageBody := messageService.ClientContacts{
		ContactsList: cont.ContactsList,
	}
	newRawMessageBody, err := json.Marshal(newMessageBody)
	if err != nil {
		loger.Log.Warningf("Can`t marshal login message. %s", err)
		return
	}
	sendMsg(conn, newMessageHeader, newRawMessageBody)

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
