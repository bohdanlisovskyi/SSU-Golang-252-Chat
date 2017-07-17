package messageService

import (
	"time"
	"encoding/json"
	"log"
	"errors"
)

//structure for message which sent between users
type Message struct {
	Header MessageHeader `json:"header"`
	Body   MessageBody `json:"body"`
}

//structure of header for Message
type MessageHeader struct {
	Type_    string `json:"type"`    // this named "Type_" because "type" - keyword
	Command  string `json:"command"` //depend on action from user (e.g. sent message, set info e.t.c.)
	Sender   int `json:"sender"`
	Receiver int `json:"receiver"`
	Time     time.Time `json:"time"`
	Auth     string `json:"auth"` // value of Auth is "token" at the beginning (maybe use JSON web token?)
}

//structure of body for Message
type MessageBody struct {
	Text string `json:"text"`
	//links []link -> slice of formatted URL's
}

//function for unmarshaling message (from [] byte JSON to Message structure)
func UnmarshalMessage(byteMessage [] byte) (Message, error) {
	var MessageStructure Message
	err := json.Unmarshal(byteMessage, &MessageStructure)
	if err != nil {
		log.Fatal("Error has occured: ", err)
		return MessageStructure, err
	}
	return MessageStructure, err
}

//function for marshaling message (Message -> [] byte JSON)
func MarshalMessage(message Message) ([] byte, error) {
	msgJSON, err := json.Marshal(message)
	if err != nil {
		log.Fatal("Error has occured: ", err)
		return nil, err
	}
	return msgJSON, err
}

//function for unmarshaling request(without info about structure) in []byte  which come from UI into map[string]interface{}
func UnmarshalRequest(byteRequest [] byte) (map[string]interface{}, error) {
	var unmarshaledRequest map[string]interface{}
	err := json.Unmarshal(byteRequest, &unmarshaledRequest)
	if err != nil {
		log.Fatal("Error has occured: ", err)
		return nil, err
	}
	return unmarshaledRequest, err
}

//function for server to retrieving "type" value from unmarshaled request
func GetType(unmarshaledRequest map[string]interface{}) (string, error) {
	typeInterface := unmarshaledRequest["type"]
	if typeInterface == nil {
		err := errors.New("Can't find value for type")
		log.Fatal("Error has occured: ", err)
		return "", err
	}
	typeValue := typeInterface.(string)
	return typeValue, nil
}

//function for server to retrieving "command" value from unmarshaled request
func GetCommand(unmarshaledRequest map[string]interface{}) (string, error) {
	valInterface := unmarshaledRequest["command"]
	if valInterface == nil {
		err := errors.New("Can't find value for command")
		log.Fatal("Error has occured: ", err)
		return "", err
	}
	value := valInterface.(string)
	return value, nil
}
