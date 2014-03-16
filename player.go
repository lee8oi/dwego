package main

import (
	"fmt"
)

type Player struct {
	ID             int
	Name, Location string
}

func (c *connection) Move(d string) {
	room := rooms[fmt.Sprintf("%s", c.player.Location)]
	if dr, ro := room.CheckExit(d); dr != "" {
		c.player.Location = ro
		c.Send("going " + dr)
		c.Send(rooms[ro].Description)
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
		c.player.Location = "0"
		//room :=
		c.Send("nickname has been set to: " + c.player.Name)
		r := rooms["0"]
		c.Send(r.Description)
	}
}
