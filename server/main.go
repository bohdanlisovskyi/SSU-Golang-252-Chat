package main

import (
	"net/http"
	"os"

	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/core"
	"github.com/gorilla/mux"
)

func main() {

	port := os.Getenv("PORT")
	loger.Log.Infof("Server run with port: " + port)
	r := mux.NewRouter()
	r.HandleFunc("/message", core.MessageHandler) // listen message
	err := http.ListenAndServe(":" + port, r)

	if err != nil {
		loger.Log.Panicf("Cannot run server %s", err.Error())
	}
}
