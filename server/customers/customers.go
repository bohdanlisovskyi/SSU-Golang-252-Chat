package customers

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn  *websocket.Conn
	Token string
}

var Clients = map[string]Client{}
