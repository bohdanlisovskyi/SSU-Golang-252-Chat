package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Greckas/SSU-Golang-252-Chat/loger"
	"github.com/Greckas/SSU-Golang-252-Chat/server/core"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/message", core.MessageHandler) // listen message

	err := http.ListenAndServe(":3002", r)

	if err != nil {

		loger.Log.Panicf("Cannot run server %s", err.Error())
	}
}
