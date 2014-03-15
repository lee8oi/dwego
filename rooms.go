package main

type Room struct {
	ID          int
	Exit        map[string]int
	Description string
}

func init() {
	rooms[1] = Room{
		ID:          1,
		Description: "The first room! There's an exit to the north.",
		Exit:        make(map[string]int),
	}
	rooms[1].Exit["north"] = 2
	rooms[2] = Room{
		ID:          2,
		Description: "The second room! There's an exit to the south.",
		Exit:        make(map[string]int),
	}
	rooms[2].Exit["south"] = 1
}
