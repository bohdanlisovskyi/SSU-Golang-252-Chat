package messageService

import (
	"encoding/json"
	"log"
)

// Message is a structure for message which is sending between users
type Message struct {
	Header MessageHeader `json:"header"`
	Body   MessageBody `json:"body"`
}

// MessageHeader is a structure of header for Message
type MessageHeader struct {
	Type_    string `json:"type"`
	Command  string `json:"command"`
	UserName string `json:"userName"`
	Token    string `json:"token"`
}

// MessageBody is a structure of body for Message
type MessageBody struct {
	ReceiverName string `json:"receiverName"`
	Time         int `json:"time"` // unix time will be used http://www.unixtimestamp.com/
	Text         string `json:"text"`
}

// UnmarshalMessage is a function for unmarshaling message (from [] byte JSON to Message structure)
func UnmarshalMessage(byteMessage [] byte) (*Message, error) {
	var MessageStructure *Message
	err := json.Unmarshal(byteMessage, &MessageStructure)
	if err != nil {
		log.Println("Error has occured: ", err) // it will be synchronized with Bohdan's logger-lib after merging initial PR's
		return MessageStructure, err
	}
	return MessageStructure, err
}

// MarshalMessage is a function for marshaling message (Message -> [] byte JSON)
func MarshalMessage(message Message) ([] byte, error) {
	msgJSON, err := json.Marshal(message)
	if err != nil {
		log.Println("Error has occured: ", err) // it will be synchronized with Bohdan's logger-lib after merging initial PR's
		return nil, err
	}
	return msgJSON, err
}

// UnmarshalRequest is a function for unmarshaling request(without info about structure) in []byte  which come from UI into map[string]interface{}
func UnmarshalRequest(byteRequest [] byte) (map[string]interface{}, error) {
	var unmarshaledRequest map[string]interface{}
	err := json.Unmarshal(byteRequest, &unmarshaledRequest)
	if err != nil {
		log.Println("Error has occured: ", err) // it will be synchronized with Bohdan's logger-lib after merging initial PR's
		return nil, err
	}
	return unmarshaledRequest, err
}