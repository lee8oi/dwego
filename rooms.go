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

//CheckExit returns the room index if an exit is available going 'way'. Empty string means no exit that way.
func (r *Room) CheckExit(way string) (room string) {
	for _, val := range r.Exits {
		if val.Direction == way {
			room = val.Destination
		}
	}
	return
}

//GetExits returns a string containing the current rooms available exits.
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

func (w *world) LoadRooms(path string) {
	w.Rooms = make(map[string]Room)
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
		if err := writeJSON(path, l); err != nil {
			fmt.Println("error writing json: ", err)
		} else {
			fmt.Println("New default rooms json file written to " + path + ".")
		}
		MapRooms(l)
	} else {
		MapRooms(l)
		fmt.Println("Rooms in " + path + " are now loaded.")
	}
}

//MapRooms takes a loaded json list of rooms and maps the ID's to World Rooms keys.
func MapRooms(l []*Room) {
	for key, val := range l {
		id := l[key].ID
		if World.Rooms != nil {
			World.Rooms[id] = *val
		}
	}
}
