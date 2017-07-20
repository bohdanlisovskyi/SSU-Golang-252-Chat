package messageService

import (
	"testing"
	"SSU-Golang-252-Chat/loger"
)

var (
	Type_        = "messageService"
	Command      = "send message"
	UserName     = "userOne"
	ReceiverName = "userTwo"
	CurrentTime  = 1500373713
	Token        = "token"
	Text         = "msg from user 1 to user 2"
	byteMessage  = []byte(`{"header":{"type":"messageService","command":"send message","sender":1,"receiver":2,"time":1500373713,"auth":"token"},"body":{"text":"msg from user 1 to user 2"}}`)
	byteHeader   = []byte(`{"type": "message", "command": "sentMessage", "sender": 1, "receiver": 2, "time": "hereAndNow", "auth": "token"}`)
	byteBody     = []byte(`{"text": "msg from user 1 to user 2"}`)
)

func TestMarshalMessage(t *testing.T) {
	header := MessageHeader{Type_, Command, UserName, Token}
	body := MessageBody{ReceiverName, CurrentTime, Text}
	message := Message{header, body}
	mrsMsg, err := MarshalMessage(&message)
	if err != nil {
		loger.Log.Fatal("Error has occured: ", err)
	}
	loger.Log.Printf("Message after marshaling in []byte: ", mrsMsg)
}

func TestUnmarshalMessage(t *testing.T) {
	ThisMessage, err := UnmarshalMessage(byteMessage)
	if err != nil {
		t.Fatal("Error has occured: ", err)
	}
	loger.Log.Printf("Message after unmarshaling has type %T", ThisMessage)
	loger.Log.Printf("Message after unmarshaling : %+v \n ", ThisMessage)
}

func Test1UnmarshalRequest(t *testing.T) {
	MessageInByte, err := UnmarshalRequest(byteMessage)
	if err != nil {
		t.Fatal("Error has occured: ", err)
	}
	loger.Log.Printf("Message after unmarshaling has type %T /n", MessageInByte)
	loger.Log.Printf("Message after unmarshaling : %+v ", MessageInByte)
}

func Test2UnmarshalRequest(t *testing.T) {
	HeaderInByte, err := UnmarshalRequest(byteHeader)
	if err != nil {
		t.Fatal("Error has occured: ", err)
	}
	loger.Log.Printf("Message after unmarshaling has type %T /n", HeaderInByte)
	loger.Log.Printf("Message after unmarshaling : %+v  ", HeaderInByte)
}

func Test3UnmarshalRequest(t *testing.T) {
	BodyInByte, err := UnmarshalRequest(byteBody)
	if err != nil {
		t.Fatal("Error has occured: ", err)
	}
	loger.Log.Printf("Message after unmarshaling has type %T \n", BodyInByte)
	loger.Log.Printf("Message after unmarshaling : %+v  ", BodyInByte)
}