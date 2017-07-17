package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/8tomat8/SSU-Golang-252-Chat/chat-log"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/core"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/message", core.MessageHandler) // listen message

	err := http.ListenAndServe(":3002", r)

	chatlog.HandleError(err, "Can't connect to server error: ")
}
