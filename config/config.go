package config

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"path"
	"os"
	"github.com/8tomat8/SSU-Golang-252-Chat/chat-log"
)

type Storage struct {
	Driver string `json:"driver"`
	Name string `json:"name"`
}

type Config struct {
	Storage Storage `json:"storage"`
}

var config Config

func init() {

	data, err := ioutil.ReadFile(GetFilePath("src/github.com/8tomat8/SSU-Golang-252-Chat/config/config.json"))

	chatlog.HandleError(err, "Can't read config.")

	config = Config{}

	if err := json.Unmarshal(data, &config); err != nil {

		chatlog.HandleError(err, "Can't read config.")
	}

}

func GetConfig() Config {
	return config
}

func GetFilePath(file string) string {
	return path.Join(os.Getenv("GOPATH"), file)
}