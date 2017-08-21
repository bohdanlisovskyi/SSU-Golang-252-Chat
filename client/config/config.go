package config

import (
	"encoding/json"
	"io/ioutil"
	"sync"

	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
)

type Server struct {
	Host string `json:"host"`
	Path string `json:"path"`
}

type MessageType struct {
	Message  string `json:"message"`
	Auth     string `json:"authorization"`
	Register string `json:"registration"`
	Contacts string `json:"contacts"`
	Settings string `json:"settings"`
}

type MessageCommand struct {
	SendMessage      string `json:"sendmessage"`
	ReceiveMessage   string `json:"receivemessage"`
	MessageSent      string `json:"messagesent"`
	SendLoginData    string `json:"sendlogindata"`
	LoginIsSucc      string `json:"loginissucc"`
	SendRegisterData string `json:"sendregisterdata"`
	RegisterIsSucc   string `json:"registerissucc"`
	SendSettings     string `json:"sendsettings"`
	SettingsChanged  string `json:"settingschsnged"`
	ContactsRequest  string `json:"contactsrequest"`
}

type Config struct {
	Server         Server         `json:"server"`
	MessageType    MessageType    `json:"messagetype"`
	MessageCommand MessageCommand `json:"messagecommand"`
}

var config *Config
var once sync.Once

func GetConfig() Config {
	once.Do(func() {
		data, err := ioutil.ReadFile("../../config.json")
		if err != nil {
			loger.Log.Panicf("Can`t read config file. %s", err.Error())
		}
		config = &Config{}

		if err := json.Unmarshal(data, &config); err != nil {
			loger.Log.Panicf("Corrupted data in connfig file. %s", err.Error())
		}
	})
	return *config
}
