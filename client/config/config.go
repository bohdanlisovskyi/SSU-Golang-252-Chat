package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
)

type Server struct {
	Host string `json:"host"`
	Path string `json:"path"`
}

type MessageType struct {
	Message  string `json:"message"`
	Auth     string `json:"authorization"`
	Contacts string `json:"contacts"`
	Settings string `json:"settings"`
}

type MessageCommand struct {
	SendMessage string `json:"sendmessage"`
	MessageSent string `json:"messagesent"`
}

type Config struct {
	Server         Server         `json:"server"`
	MessageType    MessageType    `json:"messagetype"`
	MessageCommand MessageCommand `json:"messagecommand"`
}

var config *Config

func init() {
	data, err := ioutil.ReadFile("../../config.json")
	if err != nil {
		loger.Log.Panicf("Can`t read config file. %s", err.Error())
	}
	config = &Config{}

	if err := json.Unmarshal(data, &config); err != nil {
		loger.Log.Panicf("Corrupted data in connfig file. %s", err.Error())
	}
}

func GetConfig() Config {
	return *config
}
