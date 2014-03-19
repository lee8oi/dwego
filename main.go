/*
dwego can be thought of as muD - WEbsockets - Go (making it up as I go). It is,
as it might sound, a web-based MUD (Multi-User Dungeon) server written in Go (golang.org).
The server utilizes websockets to enable users to join & play via a capable web browser
(Try Firefox or Chrome).
*/
package main

import (
	"flag"
	//"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

var addr = flag.String("addr", ":8080", "http service address")
var homeTempl *template.Template
var scriptTmpl *template.Template

var players = make(map[int]Player)
var World world

func homeHandler(c http.ResponseWriter, req *http.Request) {
	homeTempl.Execute(c, req.Host)
}

func scriptHandler(c http.ResponseWriter, req *http.Request) {
	scriptTmpl.Execute(c, req.Host)
}

type world struct {
	Rooms map[string]Room
}

func init() {
	os.Mkdir("./world", 0744)
	os.Mkdir("./conf/", 0700)
	World.LoadRooms("world/rooms.json")
}

func main() {
	flag.Parse()
	go h.run()
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/scripts", scriptHandler)
	http.HandleFunc("/ws", wsHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./pub"))))
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
