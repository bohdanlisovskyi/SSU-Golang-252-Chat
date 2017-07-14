package messageService

import (
	"time"
	"encoding/json"
	"fmt"
	"log"
)

//server address, IP and port will fill later
const server = "IP:port"

//structure for message(should be in JSON format) which consist with header and body
type Message struct {
	Header MessageHeader `json:"header"`
	Body MessageBody `json:"body"`
}

// header for Message
type MessageHeader struct {
	Type_ string `json:"type"`// this named "Type_" because "type" - keyword
	Command string `json:"command"` //depend on action from user (e.g. sent message, set info e.t.c.)
	Sender int `json:"sender"`
	Receiver int `json:"receiver"`
	Time time.Time `json:"time"`
	Auth string `json:"auth"`// value of Auth is "token" at the beginning (maybe use JSON web token?)
}

// body for message
type MessageBody struct {
	Text string `json:"text"`
	//links []link -> slice of formatted URL's
}

//func unmarshal message ([] byte JSON -> Message)
func UnmarshalMessage(byteMessage [] byte ) Message {
	var message Message
	err := json.Unmarshal(byteMessage, &message)
	if err != nil{
		log.Println("Can't unmarshal this array of bytes :", err)
	}
	return message
}

//func for marshaling message ([] byte or Message ??? -> [] byte JSON)
func MarshalMessage(message Message) [] byte {
	msgJSON, err := json.Marshal(message) // marshal message to JSON
	if err != nil {
		fmt.Println("error:", err)
	}
	return msgJSON
}

//func from parsing type from JSON
func GetTypeFromHeader(message Message) string {
	//implement this
	return "type"
}

//func for parsing command from JSON
func GetCommandFromHeader(message Message) string {
	//implement this
	return "command"
}

//function for creating massages
//from where data about SenderId, ReceiverId should come?Maybe, from UI, when user pick-up contact from contact-list
//func  MakeMessage(SenderId, ReceiverId int, Text string)  []byte{
	//header := MessageHeader{, SenderId, ReceiverId, time.Now(), "token"}//creating MessageHeader

	//Text = "message for Go252Lv team" // this hardcoded just for initial checking
	//body := MessageBody{Text}//text should be read from UI and passed here

	//msg := Message{header, body}
	//msgJSON, err := json.Marshal(msg) // marshal message to JSON
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//return msgJSON
//}

//func validate message



// function for sending(using WebSocket) messages to server
func SentMessage(message Message, addr string) bool  {

	return false
}