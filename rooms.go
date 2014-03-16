package main

import (
	"fmt"
)

type Exit struct {
	Direction   string
	Destination string
}

type Room struct {
	Exits           []Exit
	Description, ID string
}

func (r *Room) CheckExit(way string) (direction, room string) {
	for _, val := range r.Exits {
		if val.Direction == way {
			direction = way
			room = val.Destination
		}
	}
	return
}

func (r *Room) GetExits() (s string) {
	s = "Exits: "
	c := 0
	for _, val := range r.Exits {
		c++
		s += " " + val.Direction
	}
	if c == 0 {
		s += "none"
	}
	return
}

func MapRooms(l []*Room) {
	for key, val := range l {
		id := l[key].ID
		if rooms != nil {
			rooms[id] = *val
		}
	}
}

func LoadRooms(path string) {
	rooms = make(map[string]Room)
	var l []*Room
	if err := loadJSON(path, &l); err != nil {
		l = append(l, &Room{
			ID:          "0",
			Description: "The first room!",
			Exits: []Exit{
				Exit{Direction: "north", Destination: "1"},
			},
		})
		l = append(l, &Room{
			ID:          "1",
			Description: "The second room!",
			Exits: []Exit{
				Exit{Direction: "south", Destination: "0"},
				Exit{Direction: "east", Destination: "2"},
			},
		})
		l = append(l, &Room{
			ID:          "2",
			Description: "The third room!",
			Exits: []Exit{
				Exit{Direction: "west", Destination: "1"},
			},
		})
		if err := writeJSON("rooms.json", l); err != nil {
			fmt.Println("error writing json: ", err)
		} else {
			fmt.Println("new default json written to 'rooms.json'")
		}
		MapRooms(l)
	} else {
		MapRooms(l)
		fmt.Println(path + " loaded into map")
	}
}
