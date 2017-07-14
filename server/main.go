package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/core"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/message", core.MessageHandler) // listen message

	log.Fatal(http.ListenAndServe(":3002", r))
}
