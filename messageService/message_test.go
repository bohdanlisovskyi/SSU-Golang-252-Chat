package messageService

import (
	"testing"
	"time"
)

var (
	Type_         = "messageService"
	Command       = "send message"
	Sender        = 1
	Receiver      = 2
	CurrentTime   = time.Now
	Auth          = "token"
	Text          = "msg from user 1 to user 2"
	byteMessage   = []byte(`{"header":{"type":"messageService","command":"send message","sender":1,"receiver":2,"time":"2017-07-15T01:11:04.134765+03:00","auth":"token"},"body":{"text":"msg from user 1 to user 2"}}`)
	byteHeader    = []byte(`{"type": "message", "command": "sentMessage", "sender": 1, "receiver": 2, "time": "hereAndNow", "auth": "token"}`)
	byteBody      = []byte(`{"text": "msg from user 1 to user 2"}`)
	mappedRequest = map[string]interface{}{"key1": "wrong", "key2": 1, "type": "message", "typed": "typed", "etype": "band", "comand": "wrong", "command": "sent message"}
)

func TestMarshalMessage(t *testing.T) {
	header := MessageHeader{Type_, Command, Sender, Receiver, CurrentTime(), Auth}
	body := MessageBody{Text}
	message := Message{header, body}
	mrsMsg, err := MarshalMessage(message)
	if err != nil {
		t.Fatal("Error has occured: ", err)
	}
	t.Log("Message after marshaling in []byte: \n", mrsMsg)
	t.Log("Message after marshaling in string: \n", string(mrsMsg))
}

func TestUnmarshalMessage(t *testing.T) {
	var ThisMessage Message
	ThisMessage, err := UnmarshalMessage(byteMessage)
	if err != nil {
		t.Fatal("Error has occured: ", err)
	}
	t.Logf("Message after unmarshaling has type %T \n", ThisMessage)
	t.Logf("Message after unmarshaling : %+v \n ", ThisMessage)
}

func Test1UnmarshalRequest(t *testing.T) {
	MessageInByte, err := UnmarshalRequest(byteMessage)
	if err != nil {
		t.Fatal("Error has occured: ", err)
	}
	t.Logf("Message after unmarshaling has type %T \n", MessageInByte)
	t.Logf("Message after unmarshaling : %+v \n ", MessageInByte)
}

func Test2UnmarshalRequest(t *testing.T) {
	HeaderInByte, err := UnmarshalRequest(byteHeader)
	if err != nil {
		t.Fatal("Error has occured: ", err)
	}
	t.Logf("Message after unmarshaling has type %T \n", HeaderInByte)
	t.Logf("Message after unmarshaling : %+v \n ", HeaderInByte)
}

func Test3UnmarshalRequest(t *testing.T) {
	BodyInByte, err := UnmarshalRequest(byteBody)
	if err != nil {
		t.Fatal("Error has occured: ", err)
	}
	t.Logf("Message after unmarshaling has type %T \n", BodyInByte)
	t.Logf("Message after unmarshaling : %+v \n ", BodyInByte)
}

func TestGetType(t *testing.T) {
	Value, err := GetType(mappedRequest)
	if err != nil || Value != "message" {
		t.Fatal("Error has occured: ", err)
	}
	t.Logf("Type is %T \n", Value)
	t.Logf("Type is %v", Value)
}

func TestGetCommand(t *testing.T) {
	Command, err := GetCommand(mappedRequest)
	if err != nil || Command != "sent message" {
		t.Fatal("Error has occured: ", err)
	}
	t.Logf("Command is %T \n", Command)
	t.Logf("Command is %v", Command)
}
