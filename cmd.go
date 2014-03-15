package main

import (
	"fmt"
	"strings"
)

func Parse(b []byte) (p []string) {
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
	case "l", "look":
		s = "look"
	}
	return
}

func (c *connection) Interpret(m []byte) {
	split := Parse(m)
	if cmd := Command(split[0]); cmd != "" {
		switch cmd {
		case "north", "south", "east", "west":
			Move(c, cmd)
		case "nick":
			if len(split) == 1 {
				c.send <- []byte("usage: nick <nickname>")
				return
			}
			SetNick(c, split[1])
		case "look":
			fmt.Println(rooms[c.player.Location].Description)
			c.send <- []byte(rooms[c.player.Location].Description)
		case "testing":
			c.send <- []byte("testing command received")
		default:
			h.broadcast <- m
		}
	}
}
