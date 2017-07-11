package com_github_yawal

import (
	"time"
	"encoding/json"
)

//server address, IP and port will fill later
const server = "IP:port"

//structure for message(should be in JSON format) which consist with header and body
type Message struct {
	Header MessageHeader
	Body MessageBody
}

// header for Message
type MessageHeader struct {
	Type_ string// this named "Type_" because "type" - keyword
	Sender int
	Receiver int
	Time time.Time
	Auth string // value of Auth is "token" at the beginning (maybe use JSON web token?)
}

// body for message
type MessageBody struct {
	Text string
	//links []link -> slice of formatted URL's
}

//function for creating massages
//from where data about SenderId, ReceiverId should come?Maybe, from UI, when user pick-up contact from contact-list
func  MakeMessage(SenderId, ReceiverId int, Text string)  []byte{
	header := MessageHeader{"message", SenderId, ReceiverId, time.Now(), "token"}//creating MessageHeader

	Text = "message for Go252Lv team" // this hardcoded just for initial checking
	body := MessageBody{Text}//text should be read from UI and passed here

	msg := Message{header, body}
	msgJSON, _ := json.Marshal(msg) // convert message to JSON
	return msgJSON
}

// function for sending(using WebSocket) messages to server
func SentMessage(message Message, addr string) bool  {

	return false
}