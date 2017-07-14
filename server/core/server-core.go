package core

import (
	"log"
	"strconv"
	"net/http"
	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	conn *websocket.Conn
	message chan []byte
}

var clients = map[string]Client{}

func MessageHandler(w http.ResponseWriter, r *http.Request) {

	conn := addNewClient(w, r)

	log.Println("Add new connection", conn)

	go func() {

		for {
			_, text, _ := conn.ReadMessage()

			receiver_id := somefunc() // This function from Valeryi return me info from header and body return to me json

			//I check this json and get receiver_id

			//send json to Valery's function which return me []byte text

			writeMsg(text, receiver_id) //I send this text []byte to receiver
		}
	}()
}

func somefunc() string{return ""} //(THIS IS THE MOCK FUNC) remove this empty func when Valery write your own func //TODO remove this

func addNewClient(w http.ResponseWriter, r *http.Request) *websocket.Conn{

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	clients[strconv.Itoa(len(clients)+1)] = Client{conn:conn}

	return conn
}


func writeMsg(text []byte, receiver_id string) {

	clients[receiver_id].conn.WriteMessage(websocket.TextMessage, text)
}
