package main

import (
	"fmt"
)

type Player struct {
	ID       int
	Name     string
	Location int
}

func (c *connection) Move(d string) {
	if rooms[c.player.Location].Exit[d] > 0 {
		c.player.Location = rooms[c.player.Location].Exit[d]
		c.Send("going " + d)
		c.Send(rooms[c.player.Location].Description)
	} else {
		c.Send("there is nothing in that direction")
	}
	return
}

func (c *connection) SetNick(n string) {
	old := c.player.Name
	c.player.Name = n
	if len(old) > 0 {
		h.Broadcast(fmt.Sprintf("%s has changed nickname to %s", old, c.player.Name))
	} else {
		c.player.Location = 1
		c.Send("nickname has been set to: " + c.player.Name)
		c.Send(rooms[1].Description)
	}
}
