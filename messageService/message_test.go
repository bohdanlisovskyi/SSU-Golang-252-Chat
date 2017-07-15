package messageService

import (
	"testing"
	"time"
	"fmt"
)

var (
	Type_ = "messageService"
	Command = "send message"
	Sender = 1
	Receiver = 2
	CurrentTime = time.Now
	Auth = "token"
	Text = "msg from user 1 to user 2"
	byteMessage = []byte(`{"header":{"type":"messageService","command":"send message","sender":1,"receiver":2,"time":"2017-07-15T01:11:04.134765+03:00","auth":"token"},"body":{"text":"msg from user 1 to user 2"}}`)
	byteHeader = []byte(`{"type": "message", "command": "sentMessage", "sender": 1, "receiver": 2, "time": "hereAndNow", "auth": "token"}`)
	byteBody = []byte(`{"text": "msg from user 1 to user 2"}`)
	)

func TestMarshalMessage(t *testing.T) {
	header := MessageHeader{Type_, Command, Sender, Receiver, CurrentTime() , Auth}
	body := MessageBody{Text}
	message := Message{header, body}
	mrsMsg, err := MarshalMessage(message)
	if err == nil {
		fmt.Println("TestMarshalMessage: PASS, Error-> ", err)
		fmt.Println("Message after marshaling in []byte: \n", mrsMsg)
		fmt.Println("Message after marshaling in string: \n",string(mrsMsg) )
		fmt.Println("========================================================")
	}
}

func TestUnmarshalMessage(t *testing.T) {
	var ThisMessage Message
	ThisMessage, err := UnmarshalMessage(byteMessage)
	if err == nil {
		fmt.Println("TestUnmarshalMessage: PASS, Error-> ", err)
		fmt.Printf("Message after unmarshaling has type %T \n",ThisMessage )
		fmt.Printf("Message after unmarshaling : %+v \n ",ThisMessage )
		fmt.Println("========================================================")
	}
}

func Test1UnmarshalAnyDataStructure(t *testing.T) {
	MessageInByte, err := UnmarshalAnyDataStructure(byteMessage)
	if err == nil {
		fmt.Println("Test1UnmarshalAnyDataStructure: PASS, Error-> ", err)
		fmt.Printf("Message after unmarshaling has type %T \n", MessageInByte)
		fmt.Printf("Message after unmarshaling : %+v \n ", MessageInByte)
		fmt.Println("========================================================")
	}
}

func Test2UnmarshalAnyDataStructure(t *testing.T) {
	HeaderInByte, err := UnmarshalAnyDataStructure(byteHeader)
	if err == nil {
		fmt.Println("Test2UnmarshalAnyDataStructure: PASS, Error-> ", err)
		fmt.Printf("Message after unmarshaling has type %T \n", HeaderInByte)
		fmt.Printf("Message after unmarshaling : %+v \n ", HeaderInByte)
		fmt.Println("========================================================")
	}
}

func Test3UnmarshalAnyDataStructure(t *testing.T) {
	BodyInByte, err := UnmarshalAnyDataStructure(byteBody)
	if err == nil {
		fmt.Println("Test3UnmarshalAnyDataStructure: PASS, Error-> ", err)
		fmt.Printf("Message after unmarshaling has type %T \n", BodyInByte)
		fmt.Printf("Message after unmarshaling : %+v \n ", BodyInByte)
		fmt.Println("========================================================")
	}
}

func TestGetType(t *testing.T) {
	Value, err := GetType(byteHeader)
	if err == nil {
		fmt.Println("TestGetType: PASS, Error-> ", err)
		fmt.Printf("Type is %T \n", Value)
		if Value == "message" {
			fmt.Printf("Correct, Value is -> %v", Value)
		}else {
			fmt.Printf("Wrong, \"message\" != %v", Value)
		}
	}
	fmt.Println("\n ========================================================")
}

func TestGetCommand(t *testing.T) {
	Command, err := GetCommand(byteHeader)
	if err == nil {
		fmt.Println("TestGetCommand: PASS, Error-> ", err)
		fmt.Printf("Command has type %T \n", Command)
		if Command == "sentMessage" {
			fmt.Printf("Correct, Value is -> %v", Command)
		}else {
			fmt.Printf("Wrong, \"sentMessage\" != %v", Command)
		}
	}
	fmt.Println("\n ========================================================")
}


