/*
dwego can be thought of as muD - WEbsockets - Go (making it up as I go). It is,
as it might sound, a web-based MUD (Multi-User Dungeon) server written in Go (golang.org).
The server utilizes websockets to enable users to join & play via a capable web browser
(Try Firefox or Chrome).
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var addr = flag.String("addr", ":8080", "http service address")
var homeTempl = template.Must(template.ParseFiles("home.html"))
var players = make(map[int]Player)
var World world

func homeHandler(c http.ResponseWriter, req *http.Request) {
	homeTempl.Execute(c, req.Host)
}

type world struct {
	Rooms map[string]Room
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
			fmt.Println("new default json written to ", path)
		}
		MapRooms(l)
	} else {
		MapRooms(l)
		fmt.Println(path + " loaded into map")
	}
}

func init() {
	World.LoadRooms("rooms.json")
}

func main() {
	flag.Parse()
	go h.run()
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", wsHandler)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
