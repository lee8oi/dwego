package main

import (
	"fmt"
)

type Player struct {
	ID       int
	Name     string
	Location int
}

func Move(c *connection, d string) {
	if rooms[c.player.Location].Exit[d] > 0 {
		c.player.Location = rooms[c.player.Location].Exit[d]
		c.send <- []byte("going " + d)
		c.send <- []byte(rooms[c.player.Location].Description)
	} else {
		c.send <- []byte("there is nothing in that direction")
	}
	return
}

func SetNick(c *connection, n string) {
	old := c.player.Name
	c.player.Name = n
	if len(old) > 0 {
		h.broadcast <- []byte(fmt.Sprintf("%s has changed nickname to %s", old, c.player.Name))
	} else {
		c.player.Location = 1
		c.send <- []byte("nickname has been set to: " + c.player.Name)
		c.send <- []byte(rooms[1].Description)
	}
}
