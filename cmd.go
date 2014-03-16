package main

import (
	"fmt"
	"strings"
)

//Command is a simple function for mapping shortcuts to full command names.
func Command(c string) (s string) {
	switch strings.ToLower(c) {
	case "s", "south":
		s = "south"
	case "n", "north":
		s = "north"
	case "e", "east":
		s = "east"
	case "w", "west":
		s = "west"
	case "ni", "nick":
		s = "nick"
	case "ex", "exit", "exits":
		s = "exits"
	case "l", "look":
		s = "look"
	}
	return
}

//Parse parses a message and handles command if one exists.
func (c *connection) Parse(m []byte) {
	s := fmt.Sprintf("%s", m)
	p := strings.Fields(s)
	if cmd := Command(p[0]); cmd != "" {
		switch cmd {
		case "north", "south", "east", "west":
			c.Move(cmd)
		case "nick":
			if len(p) == 1 {
				c.send <- []byte("usage: nick <nickname>")
				return
			}
			c.SetNick(p[1])
		case "look":
			r := World.Rooms[c.player.Location]
			c.Send(r.Description)
		case "exits":
			r := World.Rooms[c.player.Location]
			c.Send(r.GetExits())
		case "testing":
			c.Send("testing command received")
		}
	}
}
