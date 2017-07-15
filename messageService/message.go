package messageService

import (
	"time"
	"encoding/json"
	"log"
)

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

// body for Message
type MessageBody struct {
	Text string `json:"text"`
	//links []link -> slice of formatted URL's
}

//func for unmarshalling message ([] byte JSON -> Message struct)
func UnmarshalMessage(byteMessage [] byte ) (Message, error){
	var MessageStructure Message
	err := json.Unmarshal(byteMessage, &MessageStructure)
	if err != nil{
		log.Println("error has occured: ", err)
		return MessageStructure, err
	}
	return MessageStructure, err
}

//func for marshaling message (Message -> [] byte JSON)
func MarshalMessage(message Message) ([] byte, error ){
	msgJSON, err := json.Marshal(message)
	if err != nil {
		log.Println("error has occured: ", err)
		return nil, err
	}
	return msgJSON, err
}

//func for unmarshalling []byte into map[string]interface{} without knowing data's structure
func UnmarshalAnyDataStructure(byteSlice [] byte) (map[string]interface{}, error) {
	var dataStructure interface{}
	err := json.Unmarshal(byteSlice, &dataStructure)
	mapStructure := dataStructure.(map[string]interface{})
	if err != nil{
		log.Println("error has occured: ", err)
		return nil, err
	}
	return mapStructure, err
}

//function for retrieving "type" value from JSON requests
func GetType(byteSlice [] byte) (string, error) {
	JSONMap, err := UnmarshalAnyDataStructure(byteSlice)
	if err != nil {
		log.Println("error has occured: ", err)
		return "", err
	}
	var value string
	for k, v := range JSONMap {
		//value := string(v)
		//fmt.Println(k , v)
		if k == "type" {
			value := v.(string)

			//v := string(v)
			return value, err
		}
	}
	return value, err
}

//function for retrieving "command" value from JSON requests
func GetCommand(byteSlice [] byte) (string, error) {
	JSONMap, err := UnmarshalAnyDataStructure(byteSlice)
	if err != nil {
		log.Println("error has occured: ", err)
		return "", err
	}
	var value string
	for k, v := range JSONMap {
		if k == "command" {
			value := v.(string)
			return value, err
		}
	}
	return value, err
}