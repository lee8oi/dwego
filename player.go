package main

import (
	"fmt"
)

type Player struct {
	ID             int
	Name, Location string
}

//Move changes players location if there is an exit in the specified direction.
func (c *connection) Move(d string) {
	room := World.Rooms[fmt.Sprintf("%s", c.player.Location)]
	if r := room.CheckExit(d); r != "" {
		c.player.Location = r
		c.Send("going " + d)
		c.Send(World.Rooms[r].Description)
	} else {
		c.Send("there is nothing in that direction")
	}
	return
}

//SetNick handles setting the player nickname. Initializes player object if necessary.
func (c *connection) SetNick(n string) {
	old := c.player.Name
	c.player.Name = n
	if len(old) > 0 {
		h.Broadcast(fmt.Sprintf("%s has changed nickname to %s", old, c.player.Name))
	} else {
		c.player.Location = "0"
		c.Send("nickname has been set to: " + c.player.Name)
		c.Send(World.Rooms["0"].Description)
	}
}
