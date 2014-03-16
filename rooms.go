package main

import (
//"fmt"
)

type Exit struct {
	Direction   string
	Destination string
}

type Room struct {
	Exits           []Exit
	Description, ID string
}

func (r *Room) CheckExit(way string) (room string) {
	for _, val := range r.Exits {
		if val.Direction == way {
			room = val.Destination
		}
	}
	return
}

func (r *Room) GetExits() (s string) {
	for _, val := range r.Exits {
		s += " " + val.Direction
	}
	if s == "" {
		s += "none"
	} else {
		s = "Exits: " + s
	}
	return
}

//MapRooms takes a loaded json list of rooms and maps the ID's to World keys.
func MapRooms(l []*Room) {
	for key, val := range l {
		id := l[key].ID
		if World.Rooms != nil {
			World.Rooms[id] = *val
		}
	}
}
