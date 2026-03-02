package main

import (
	"flag"
	"fmt"
)

func main() {
	var configFile, algorithm string
	var animate bool
	var cat bool
	var isHouse bool
	var useLogic bool

	flag.StringVar(&configFile, "file", "empty.json", "configuration file")
	flag.StringVar(&algorithm, "algorithm", "random", "cleaning algorithm")
	flag.BoolVar(&animate, "animate", true, "animate while cleaning")
	flag.BoolVar(&cat, "cat", false, "add a cat to the room")
	flag.BoolVar(&isHouse, "house", false, "config file has multiple rooms")
	flag.BoolVar(&useLogic, "logic", false, "use propositional logic")

	flag.Parse()

	var house *House

	if !isHouse {
		// If not a house, we just have a room. Create a house and assign one room to it
		// This way, we can use the same loop for houses and for individual rooms

		var rooms []*Room
		// Get a room from json
		room := NewRoom(configFile, animate)

		rooms = append(rooms, room)
		var h House
		h.Rooms = rooms
		house = &h

	} else {
		// We are doing a comlete house. Just get a house from json

		house = NewHouse(configFile, animate)
	}

	roomCount := 0

	if useLogic {
		// Use propositional logic for cleaning
		// rumba := newRobotWithLogic(1, 1)
	} else {
		// Use the original cleaning approach without propositional approach and for multiple logic
		for _, room := range house.Rooms {
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
			roomCount++
		}
	}
	fmt.Printf("All done. Cleaned a total of %d room(s)\n", roomCount)

}
