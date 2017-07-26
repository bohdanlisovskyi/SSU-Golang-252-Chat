package main

import (
	"github.com/gorilla/websocket"
	"net/url"
	"log"
	"flag"
	"time"
	"os"
	"os/signal"
	"strconv"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"encoding/json"
)

var addr = flag.String("addr", "localhost:3002", "http service address")

func main() {

	go connect1()

	time.Sleep(time.Second * 2)

	connect2()

	/*u := url.URL{Scheme: "ws", Host: *addr, Path: "/message"}

	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Fatal("dial:", err)
	}

	err = c.WriteMessage(websocket.TextMessage, []byte("some string"))

	if err != nil {
		log.Println("write:", err)
		return
	}

	u2 := url.URL{Scheme: "ws", Host: *addr, Path: "/message"}

	log.Printf("connecting to %s", u2.String())

	c2, _, err := websocket.DefaultDialer.Dial(u2.String(), nil)

	if err != nil {
		log.Fatal("dial:", err)
	}

	err = c2.WriteMessage(websocket.TextMessage, []byte("some string"))

	if err != nil {
		log.Println("write:", err)
		return
	}

	select {

	}*/
}

func connect1() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/message"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Fatal("dial:", err)
	}

	defer c.Close()

	go func() {

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	User1Id := "123"

	register1 := messageService.Message{
		Header:messageService.MessageHeader{
			Type_:"register",
			UserName: User1Id,
		},
		Body:messageService.MessageBody{

			Text:"Some text with Unix nano" + strconv.Itoa(int(time.Now().UnixNano())),
		},
	}

	msgByte, _ := json.Marshal(register1)

	err = c.WriteMessage(websocket.TextMessage, msgByte)

	select {

	}
	/*for {
		msg := messageService.Message{
			Header:messageService.MessageHeader{
				Type_:"register",
				UserName: strconv.Itoa(int(time.Now().UnixNano())),
			},
			Body:messageService.MessageBody{

				Text:"Some text with Unix nano" + strconv.Itoa(int(time.Now().UnixNano())),
			},
		}

		msgByte, _ := json.Marshal(msg)

		err = c.WriteMessage(websocket.TextMessage, msgByte)

		if err != nil {
			log.Println("write:", err)
			return
		}

		time.Sleep(time.Second * 100)
	}*/
}

func connect2() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/message"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Fatal("dial:", err)
	}

	defer c.Close()

	go func() {

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	User2Id := "321"

	register2 := messageService.Message{
		Header:messageService.MessageHeader{
			Type_:"register",
			UserName: User2Id,
		},
		Body:messageService.MessageBody{

			Text:"Some text with Unix nano" + strconv.Itoa(int(time.Now().UnixNano())),
		},
	}

	time.Sleep(time.Second * 3)
	msgByte, _ := json.Marshal(register2)

	err = c.WriteMessage(websocket.TextMessage, msgByte)

	message := messageService.Message{
		Header:messageService.MessageHeader{
			Type_:"message",
			UserName: User2Id,
		},
		Body:messageService.MessageBody{
			ReceiverName:"123",
			Text:"Some text with Unix nano" + strconv.Itoa(int(time.Now().UnixNano())),
		},
	}

	messageByte, _ := json.Marshal(message)

	err = c.WriteMessage(websocket.TextMessage, messageByte)

	select {

	}
	/*for {
		msg := messageService.Message{
			Header:messageService.MessageHeader{
				Type_:"register",
				UserName: strconv.Itoa(int(time.Now().UnixNano())),
			},
			Body:messageService.MessageBody{

				Text:"Some text with Unix nano" + strconv.Itoa(int(time.Now().UnixNano())),
			},
		}

		msgByte, _ := json.Marshal(msg)

		err = c.WriteMessage(websocket.TextMessage, msgByte)

		if err != nil {
			log.Println("write:", err)
			return
		}

		time.Sleep(time.Second * 100)
	}*/
}
