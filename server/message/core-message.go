package coremessage

import (
	"encoding/json"

	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/customers"
)

const (
	EmptyType          = ""
	MessageType        = "message"
	RegisterType       = "register"
	RegisterSuccComm   = "registerissucc"
	AuthType           = "auth"
	AuthSuccComm       = "loginissucc"
	ContactsType       = "contacts"
	ContactsRequest    = "contactsrequest"
	SettingType        = "settings"
	ChangePassComm     = "change_pass"
	BlockUserComm      = "block_user"
	ChangeNicknameComm = "change_nickname"
	ChangeBirthdayComm = "change_birthday"
	ChangeUserInfoComm = "change_userInfo"
	SentMessageComm    = "sent_message"
	StatusOk           = "Ok"
	MarshalError       = "Marshal message error"
	UnmarshalError     = "Unmarshal message error"
	ReceiverNotFound   = "Receiver not found"
	WriteMsgError      = "Write message error"
)

func SendMessage(message *messageService.Message, messageType int) messageService.Message {

	byteMessage, err := messageService.MarshalMessage(message)
	if err != nil {
		loger.Log.Errorf("Marshal message error: %s", err.Error())
		return messageService.Message{
			Header: messageService.MessageHeader{
				Type_:   MessageType,
				Command: MarshalError,
			},
			Body: message.Body,
		}
	}

	return writeMsg(byteMessage, message, messageType) //I send this text []byte to receiver
}

func writeMsg(text []byte, message *messageService.Message, messageType int) messageService.Message {
	msgBody := messageService.MessageBody{}
	err := json.Unmarshal(message.Body, &msgBody)
	if err != nil {
		loger.Log.Errorf("Unmarshal message error: %s", err.Error())

		return messageService.Message{
			Header: messageService.MessageHeader{
				Type_:   MessageType,
				Command: UnmarshalError,
			},
			Body: message.Body,
		}
	}

	client, ok := customers.Clients[msgBody.ReceiverName]
	if !ok {
		loger.Log.Warn("Receiver not found")
		return messageService.Message{
			Header: messageService.MessageHeader{
				Type_:   MessageType,
				Command: ReceiverNotFound,
			},
			Body: message.Body,
		}
	}

	err = client.Conn.WriteMessage(messageType, text)
	if err != nil {
		loger.Log.Errorf("Write message error: %s", err.Error())
		return messageService.Message{
			Header: messageService.MessageHeader{
				Type_:   MessageType,
				Command: WriteMsgError,
			},
			Body: message.Body,
		}
	}

	return messageService.Message{
		Header: messageService.MessageHeader{
			Type_:   MessageType,
			Command: StatusOk,
		},
		Body: message.Body,
	}
}
