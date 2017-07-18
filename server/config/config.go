package config

import (
	"io/ioutil"
	"encoding/json"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
)

type Storage struct {
	Driver string `json:"driver"`
	Name string `json:"name"`
}

type Config struct {
	Storage []Storage `json:"storage"`
}

var config Config

func init() {

	data, err := ioutil.ReadFile("./config.json")

	if err != nil {
		loger.Log.Panicf("Can't read config. %s", err.Error())
	}

	config = Config{}

	if err := json.Unmarshal(data, &config); err != nil {

		loger.Log.Panicf("Can't read config. %s", err.Error())
	}
}

func GetConfig() Config {
	return config
}