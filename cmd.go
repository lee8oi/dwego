package main

import (
	"fmt"
	"strings"
)

func parse(b []byte) (p []string) {
	s := fmt.Sprintf("%s", b)
	p = strings.Fields(s)
	return
}

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

func (c *connection) Interpret(m []byte) {
	split := parse(m)
	if cmd := Command(split[0]); cmd != "" {
		switch cmd {
		case "north", "south", "east", "west":
			c.Move(cmd)
		case "nick":
			if len(split) == 1 {
				c.send <- []byte("usage: nick <nickname>")
				return
			}
			c.SetNick(split[1])
		case "look":
			fmt.Println(rooms[c.player.Location].Description)
			c.Send(rooms[c.player.Location].Description)
		case "exits":
			r := rooms[c.player.Location]

			c.Send(r.Exits())
		case "testing":
			c.Send("testing command received")
		default:
			h.broadcast <- m
		}
	}
}
