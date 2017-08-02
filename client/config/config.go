package config

import (
	"io/ioutil"
	"encoding/json"
	"github.com/Greckas/SSU-Golang-252-Chat/loger"
)

type Server struct {
	Connection string	`json:"connection"`
}

type Config struct {
	Server Server 		`json:"server"`
}

var config *Config

func init(){
	data, err := ioutil.ReadFile("../../config.json")
	if err != nil {
		loger.Log.Panicf("Can`t read config file. %s", err.Error())
	}
	config = &Config{}

	if err := json.Unmarshal(data, &config); err != nil {
		loger.Log.Panicf("Corrupted data in connfig file. %s", err.Error())
	}
}

func GetConfig() Config{
	return *config
}