package main

import (
	"flag"
	"fmt"
	"time"
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
	// Add cats to room in necessary
	if cat {
		for _, room := range house.Rooms {
			room.Cat = NewCat(room)
		}
	}
	roomCount := 0

	if useLogic {
		// Use propositional logic for cleaning
		fmt.Println("Using propositional logic for cleaning decisions")
		rumba := NewRobotWithLogic(1, 1)

		// Assign a cleaning algorithm
		setUpAlgorithm(algorithm, rumba.Robot)
		// Scan the house
		roomNameToIndex := rumba.ScanHouseWithLogic(house)

		fmt.Println("\nLogical state after scanning:")
		fmt.Printf("Today is %s (Weekday: %t)\n", time.Now().Weekday(), rumba.World.IsWeekDay)
		fmt.Printf("Jack is home: %t\n", rumba.World.Jack.IsHome)
		fmt.Printf("Jack is home: %t\n", rumba.World.Sarah.IsHome)
		fmt.Printf("Jack is home: %t\n", rumba.World.Johnny.IsHome)
		fmt.Printf("Johnny's door closed: %t\n", rumba.World.Johnny.DoorClosed)

		if rumba.World.Johnny.DoorClosed {
			fmt.Println("Logic: Will not vacum Johnny's room")
		} else {
			fmt.Println("Logic: Will vacum Johnny's room")
		}

		// Determine cleaning priority based on logical rules
		cleaningPriority := rumba.World.DetermineCleaningPriority()
		fmt.Println("\nDetermined cleaning priority based on propositional logic:")
		for i, roomName := range cleaningPriority {
			fmt.Printf("%d. %s\n", i+1, roomName)
		}

		for k, v := range roomNameToIndex {
			fmt.Println(k, "->", v)
		}
		fmt.Println("\n Press enter to start cleaning")

		fmt.Scanln()

		// Clean the rooms in priority order
		for _, roomName := range cleaningPriority {
			// Check to see if room exist
			roomIndex, exist := roomNameToIndex[roomName]
			if !exist {
				fmt.Printf("Room %s not found in the house, skipping\n", roomName)
			}

			// Get the room from house.Rooms by index
			room := house.Rooms[roomIndex]
			// reset rumba position 1,1
			rumba.Position = Point{X: 1, Y: 1}
			rumba.Path = []Point{{X: 1, Y: 1}}

			// Clean the room
			rumba.Robot.CleanRoom(room, rumba.Robot)
			roomCount++

		}

	} else {
		// Use the original cleaning approach without propositional approach and for multiple logic
		for _, room := range house.Rooms {

			// Get a rumba
			rumba := NewRobot(1, 1)

			// Assign a cleaning algorithm
			setUpAlgorithm(algorithm, rumba)

			// Clean the room
			rumba.CleanRoom(room, rumba)
			roomCount++
		}
	}
	fmt.Printf("All done. Cleaned a total of %d room(s)\n", roomCount)

}

func setUpAlgorithm(algorithm string, rumba *Robot) {
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
}
