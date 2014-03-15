/*
dwego can be thought of as muD - WEbsockets - Go (making it up as I go). It is,
as it might sound, a web-based MUD (Multi-User Dungeon) server written in Go (golang.org).
The server utilizes websockets to enable users to join & play via a capable web browser
(Try Firefox or Chrome).
*/
package main

import (
	"flag"
	"log"
	"net/http"
	"text/template"
)

func homeHandler(c http.ResponseWriter, req *http.Request) {
	homeTempl.Execute(c, req.Host)
}

var addr = flag.String("addr", ":8080", "http service address")
var homeTempl = template.Must(template.ParseFiles("home.html"))
var players = make(map[int]Player)
var rooms = make(map[int]Room)

func main() {
	flag.Parse()
	go h.run()
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", wsHandler)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
