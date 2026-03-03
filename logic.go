package main

import (
	"fmt"
	"time"
)

type PersonStatus struct {
	Name       string
	IsHome     bool
	Room       string // person's room name
	DoorClosed bool
}

type LogicalWord struct {
	Jack      PersonStatus
	Sarah     PersonStatus
	Johnny    PersonStatus
	IsWeekDay bool
	Objects   map[string]bool
}

func newLogicalWorld() *LogicalWord {
	// Get current day to determine if it is a weekday

	today := time.Now()
	weekDay := today.Weekday()
	isWeekDay := weekDay >= time.Monday && weekDay <= time.Friday

	return &LogicalWord{
		Jack: PersonStatus{
			Name:   "Jack",
			IsHome: false,
			Room:   "Jack's room",
		},
		Sarah: PersonStatus{
			Name:   "Sarah",
			IsHome: false,
			Room:   "Sarah's Room",
		},
		Johnny: PersonStatus{
			Name:       "Johnny",
			IsHome:     false,
			Room:       "Johnny's Room",
			DoorClosed: false,
		},
		IsWeekDay: isWeekDay,
		Objects:   make(map[string]bool),
	}
}

func (world *LogicalWord) UpdateObjectFound(objectName string) {
	world.Objects[objectName] = true

	// Apply object to person identification rules

	// Rule 1: If backpack found Jack is home
	if objectName == "backpack" {
		world.Jack.IsHome = true
		fmt.Println("Logic: Backpack found deducing Jack is home")
	}

	// Rule 2: If bicycle found Sarah is home
	if objectName == "bicycle" {
		world.Sarah.IsHome = true
		fmt.Println("Logic: Bicycle found deducing Sarah is home")
	}

	// Rule 3: If skateboard found Johnny is home
	if objectName == "skateboard" {
		world.Johnny.IsHome = true
		fmt.Println("Logic: Skateboard found deducing Johnny is home")
	}
}

// Update door status updates whether Johnny's door open or closed
func (world *LogicalWord) UpdateDoorStatus(doorName string, isClosed bool) {

	if doorName == "Johnny's Door" {
		world.Johnny.DoorClosed = isClosed
		fmt.Printf("Logic: Johnny's door is now %s\n", map[bool]string{true: "closed", false: "open"}[isClosed])
	}
}

// Determine the rules cleaning priority

func (world *LogicalWord) DetermineCleaningPriority() []string {
	availableRooms := []string{
		"Kitchen",
		"Living Room",
		"Jack's Room",
		"Sarah's Room",
		"Johnny's Room",
	}
	// Rule: if no one is home them vacum all rooms starting from kitchen

	if !world.Jack.IsHome && !world.Sarah.IsHome && !world.Johnny.IsHome {
		fmt.Println("Logic: No one is home vacuuming all rooms starting with the kitchen")
		return availableRooms
	}

	// Initialize a priority list with all available rooms
	priorityList := make([]string, 0)
	skipRooms := make(map[string]bool)

	// Rule: if Sarah's home then don't vacuum living room
	if world.Sarah.IsHome {
		fmt.Println("Logic: Sarah is home skipping the living room")
		skipRooms["Living Room"] = true
	}

	// Rule: if Johnny is home and his door is closed skip his room

	if world.Johnny.IsHome && world.Johnny.DoorClosed {
		fmt.Println("Logic: Johnny is home and his door is closed, skipping his room ")
		skipRooms["Johny's Room"] = true
	}

	// Rule: if Jack is home and it's a weekday do his room last
	jackRoomLast := world.Jack.IsHome && world.IsWeekDay

	// Build our Priority list. add kitchen first and filter out skip rooms

	priorityList = append(priorityList, "Kitchen")

	// Add all other rooms except Jack's if it needs to be last and skipped rooms
	for _, room := range availableRooms {
		if room == "Kitchen" {
			continue
		}

		if room == "Jack's Room" && jackRoomLast {
			continue // will be added last
		}

		if skipRooms[room] {
			continue
		}
		priorityList = append(priorityList, room)
	}

	// Add Jack's room if it should be last
	if jackRoomLast && !skipRooms["Jack's Room"] {
		priorityList = append(priorityList, "Jack's Room")
	}
	return priorityList
}

// Type for a robot for logic

// Factory methot for a robot with Logic
