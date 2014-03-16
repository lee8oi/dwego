package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type connection struct {
	ws     *websocket.Conn
	player Player
	send   chan []byte
}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		c.Interpret(message)
		time.Sleep(time.Second)
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.send {
		err := c.ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

func (c *connection) Send(s string) {
	if len(s) > 0 {
		c.send <- []byte(s)
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		return
	}
	c := &connection{send: make(chan []byte, 256), ws: ws}
	h.register <- c
	defer func() { fmt.Println(c.player.Name, " has disconnected"); h.unregister <- c }()
	go c.writer()
	c.send <- []byte("set your nickname with 'nick <nickname>'")
	c.reader()
}
