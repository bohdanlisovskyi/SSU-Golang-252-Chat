package modules

import (
	"encoding/json"

	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/auth"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/customers"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/message"
	"github.com/8tomat8/SSU-Golang-252-Chat/contactService"
	"github.com/gorilla/websocket"
)

func EmptyType() {
	loger.Log.Errorf("Message Header Type Empty")
}

func Message(message *messageService.Message, messageType int, conn *websocket.Conn) {

	err := coremessage.SendMessage(message, messageType)
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
	customers.Clients[message.Header.UserName] = customers.Client{Conn: conn, Token: tok}
}

func AddContact(message *messageService.Message, messageType int, conn *websocket.Conn) {
	var newContact *contactService.Contacts
	var err error
	newContact, err = contactService.UnmarshalContacts(*message)

	if _, err = contactService.AddContact(*newContact); err != nil {
		loger.Log.Errorf("adding contact into contacts table failed server error: ", err)
		byteError, _ := json.Marshal(
			messageService.Message{
				Header:
				messageService.MessageHeader{
					Type_:   coremessage.AddContactCommand,
					Command: "Contact was not added",
				},
				Body: nil,
			})
		conn.WriteMessage(messageType, byteError)
		return
	}

	byteSuccess, _ := json.Marshal(
		messageService.Message{
			Header:
			messageService.MessageHeader{
				Type_:   coremessage.AddContactCommand,
				Command: coremessage.StatusOk,
			},
			Body: nil,
		})
	conn.WriteMessage(messageType, byteSuccess)
	return
}

func DeleteContact(message *messageService.Message, messageType int, conn *websocket.Conn) {
	var deleteContact *contactService.Contacts
	var err error
	deleteContact, err = contactService.UnmarshalContacts(*message)

	if _, err = contactService.RemoveContact(deleteContact.MainUser, deleteContact.ContactUser); err != nil {
		loger.Log.Errorf("deleting contact from contacts table failed server error: ", err)
		byteError, _ := json.Marshal(
			messageService.Message{
				Header:
				messageService.MessageHeader{
					Type_:   coremessage.DeleteContactCommand,
					Command: "Contact was not deleted",
				},
				Body: nil,
			})
		conn.WriteMessage(messageType, byteError)
		return
	}

	byteSuccess, _ := json.Marshal(
		messageService.Message{
			Header:
			messageService.MessageHeader{
				Type_:   coremessage.DeleteContactCommand,
				Command: coremessage.StatusOk,
			},
			Body: nil,
		})
	conn.WriteMessage(messageType, byteSuccess)
	return
}

func DeleteUser(message *messageService.Message, messageType int, conn *websocket.Conn) {
	var deleteContact *contactService.Contacts
	var err error
	deleteContact, err = contactService.UnmarshalContacts(*message)

	if _, err = contactService.RemoveUser(deleteContact.ContactUser); err != nil {
		loger.Log.Errorf("deleting user from contacts, users and authintications tables failed server error: ", err)
		byteError, _ := json.Marshal(
				messageService.Message{
				Header:
				messageService.MessageHeader{
					Type_:   coremessage.DeleteUserCommand,
					Command: "User was not deleted",
				},
				Body: nil,
			})
		conn.WriteMessage(messageType, byteError)
		return
	}

	byteSuccess, _ := json.Marshal(
		messageService.Message{
			Header:
			messageService.MessageHeader{
				Type_:   coremessage.DeleteUserCommand,
				Command: coremessage.StatusOk,
			},
			Body: nil,
		})
	conn.WriteMessage(messageType, byteSuccess)
	return
}

func ShowUserInfo(message *messageService.Message, messageType int, conn *websocket.Conn) {
	var userInfo *contactService.Users
	var err error
	userInfo, err = contactService.UnmarshalUsers(*message)

	if _, err = contactService.GetUserInfo(userInfo.UserName); err != nil {
		loger.Log.Errorf("selecting user information from users tables failed server error: ", err)
		byteError, _ := json.Marshal(
			messageService.Message{
				Header:
				messageService.MessageHeader{
					Type_:   coremessage.ContactType,
					Command: coremessage.ShowUserInfoCommand,
				},
				Body: nil,
			})
		conn.WriteMessage(messageType, byteError)
		return
	}

	byteVal, _ := json.Marshal(userInfo)
	body, _ := json.RawMessage.MarshalJSON(byteVal)
	byteSuccess, _ := json.Marshal(
		messageService.Message{
			Header:
			messageService.MessageHeader{
				Type_:   coremessage.ContactType,
				Command: coremessage.ShowUserInfoCommand,
			},
			Body: body,
		})
	conn.WriteMessage(messageType, byteSuccess)
	return
}

func ShowAllContacts(message *messageService.Message, messageType int, conn *websocket.Conn) {
	var userInfo *contactService.Users
	var userContacts []contactService.Users
	var err error
	userInfo, err = contactService.UnmarshalUsers(*message)

	if userContacts, err = contactService.GetContacts(userInfo.UserName); err != nil {
		loger.Log.Errorf("selecting all contacts failed server error: ", err)
		byteError, _ := json.Marshal(
			messageService.Message{
				Header:
				messageService.MessageHeader{
					Type_:   coremessage.ContactType,
					Command: coremessage.ShowAllContactsCommand,
				},
				Body: nil,
			})
		conn.WriteMessage(messageType, byteError)
		return
	}

	byteVal, _ := json.Marshal(userContacts)
	body, _ := json.RawMessage.MarshalJSON(byteVal)
	byteSuccess, _ := json.Marshal(
		messageService.Message{
			Header:
			messageService.MessageHeader{
				Type_:   coremessage.ContactType,
				Command: coremessage.ShowAllContactsCommand,
			},
			Body: body,
		})
	conn.WriteMessage(messageType, byteSuccess)
	return
}

func SearchContact(message *messageService.Message, messageType int, conn *websocket.Conn) {
	var search *contactService.Search
	var userContacts []contactService.Users
	var err error
	search, err = contactService.UnmarshalSearch(*message)

	if userContacts, err = contactService.SearchContacts(search.SearchContact); err != nil {
		loger.Log.Errorf("search contacts failed server error: ", err)
		byteError, _ := json.Marshal(
			messageService.Message{
				Header: messageService.MessageHeader{
					Type_:   coremessage.ContactType,
					Command: coremessage.SearchContactCommand,
				},
				Body: nil,
			})
		conn.WriteMessage(messageType, byteError)
		return
	}

	byteVal, _ := json.Marshal(userContacts)
	body, _ := json.RawMessage.MarshalJSON(byteVal)
	byteSuccess, _ := json.Marshal(
		messageService.Message{
			Header: messageService.MessageHeader{
				Type_:   coremessage.ContactType,
				Command: coremessage.SearchContactCommand,
			},
			Body: body,
		})
	conn.WriteMessage(messageType, byteSuccess)
	return
}
