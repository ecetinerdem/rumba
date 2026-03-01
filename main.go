package main

import (
	"flag"
)

func main() {
	var configFile, algorithm string
	var animate bool
	var cat bool

	flag.StringVar(&configFile, "file", "empty.json", "configuration file")
	flag.StringVar(&algorithm, "algorithm", "random", "cleaning algorithm")
	flag.BoolVar(&animate, "animate", true, "animate while cleaning")
	flag.BoolVar(&cat, "cat", false, "add a cat to the room")

	flag.Parse()

	room := NewRoom(configFile, animate)

	if cat {
		room.Cat = NewCat(room)
	}

	// Get a rumba
	rumba := NewRobot(1, 1)

	// Assign a cleaning algorithm
	switch algorithm {
	case "random":
		rumba.CleanRoom = CleanRandomWalk
	case "slam":
		rumba.CleanRoom = CleanRoomSlam
	case "spiral":
		rumba.CleanRoom = CleanRoomSpiral
	case "snake":
		rumba.CleanRoom = CleanRoomSnake
	default:
		// Do nothing
	}

	// Clean the room
	rumba.CleanRoom(room, rumba)

}
