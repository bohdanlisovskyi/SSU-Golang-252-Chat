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
	)

func TestMarshalMessage(t *testing.T) {
	header := MessageHeader{Type_, Command, Sender, Receiver, CurrentTime() , Auth}
	body := MessageBody{Text}

	message := Message{header, body}

	//fmt.Println("Message before marshaling", message)
	//MarshalMessage(message)
	fmt.Println("Message after marshaling",message)

}
