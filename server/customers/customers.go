package customers

import "github.com/gorilla/websocket"

type Client struct {
	Conn *websocket.Conn
}

var Clients = map[string]Client{}
