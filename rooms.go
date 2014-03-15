package main

type Room struct {
	ID          int
	Exit        map[string]int
	Description string
}

func (r *Room) Exits() (s string) {
	s = "Exits: "
	c := 0
	for key, _ := range r.Exit {
		c++
		s += " " + key
	}
	if c == 0 {
		s += "none"
	}
	return
}

func init() {
	rooms[1] = Room{
		ID:          1,
		Description: "The first room!",
		Exit:        make(map[string]int),
	}
	rooms[1].Exit["north"] = 2
	rooms[2] = Room{
		ID:          2,
		Description: "The second room!",
		Exit:        make(map[string]int),
	}
	rooms[2].Exit["south"] = 1
	rooms[2].Exit["east"] = 3
	rooms[3] = Room{
		ID:          2,
		Description: "The third room!",
		Exit:        make(map[string]int),
	}
	rooms[3].Exit["west"] = 2
}
