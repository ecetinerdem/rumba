package main

import (
	"fmt"
	"time"
)

type PersonStatus struct {
	Name       string
	IsHome     bool
	Room       string
	DoorClosed bool
}

type LogicalWord struct {
	Jack      PersonStatus
	Sarah     PersonStatus
	Johnny    PersonStatus
	IsWeekDay bool
	Objects   map[string]bool
}

func NewLogicalWorld() *LogicalWord {
	today := time.Now()
	weekDay := today.Weekday()
	isWeekDay := weekDay >= time.Monday && weekDay <= time.Friday

	return &LogicalWord{
		Jack: PersonStatus{
			Name:   "Jack",
			IsHome: false,
			Room:   "Jack's Room",
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
	// Avoid processing duplicates
	if world.Objects[objectName] {
		return
	}
	world.Objects[objectName] = true

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

// UpdateDoorStatus updates whether Johnny's door is open or closed
func (world *LogicalWord) UpdateDoorStatus(doorName string, isClosed bool) {
	if doorName == "johnny's door" {
		world.Johnny.DoorClosed = isClosed
		fmt.Printf("Logic: Johnny's door is now %s\n", map[bool]string{true: "closed", false: "open"}[isClosed])
	}
}

// DetermineCleaningPriority determines the cleaning order based on logical rules
func (world *LogicalWord) DetermineCleaningPriority() []string {
	availableRooms := []string{
		"Kitchen",
		"Living Room",
		"Jack's Room",
		"Sarah's Room",
		"Johnny's Room",
	}

	// Rule: if no one is home vacuum all rooms starting from kitchen
	if !world.Jack.IsHome && !world.Sarah.IsHome && !world.Johnny.IsHome {
		fmt.Println("Logic: No one is home vacuuming all rooms starting with the kitchen")
		return availableRooms
	}

	priorityList := make([]string, 0)
	skipRooms := make(map[string]bool)

	// Rule: if Sarah is home then don't vacuum living room
	if world.Sarah.IsHome {
		fmt.Println("Logic: Sarah is home skipping the living room")
		skipRooms["Living Room"] = true
	}

	// Rule: if Johnny is home and his door is closed skip his room
	if world.Johnny.IsHome && world.Johnny.DoorClosed {
		fmt.Println("Logic: Johnny is home and his door is closed, skipping his room")
		skipRooms["Johnny's Room"] = true
	}

	// Rule: if Jack is home and it's a weekday do his room last
	jackRoomLast := world.Jack.IsHome && world.IsWeekDay

	// Always start with kitchen
	priorityList = append(priorityList, "Kitchen")

	// Add all other rooms except Jack's (if it needs to be last) and skipped rooms
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

	// Add Jack's room last if needed
	if jackRoomLast && !skipRooms["Jack's Room"] {
		priorityList = append(priorityList, "Jack's Room")
	}

	return priorityList
}

// RobotWithLogic embeds Robot and adds propositional logic world
type RobotWithLogic struct {
	*Robot
	World *LogicalWord
}

// NewRobotWithLogic creates a new robot with logic
func NewRobotWithLogic(startX, startY int) *RobotWithLogic {
	return &RobotWithLogic{
		Robot: NewRobot(startX, startY),
		World: NewLogicalWorld(),
	}
}

func (robot *RobotWithLogic) ScanHouseWithLogic(house *House) map[string]int {
	roomNameToIndex := make(map[string]int)
	assignedRooms := make(map[string]bool) // FIX: track assigned bed rooms properly

	// Identify all rooms
	for i, room := range house.Rooms {
		roomName := ""

		for x := range room.Width {
			for y := range room.Height {
				if room.Grid[x][y].Type == "furniture" {
					// Identify room type from furniture
					if roomName == "" {
						switch room.Grid[x][y].ObstacleName {
						case "bed":
							// FIX: use assignedRooms map instead of broken == 0 check
							if !assignedRooms["Jack's Room"] {
								roomName = "Jack's Room"
							} else if !assignedRooms["Sarah's Room"] {
								roomName = "Sarah's Room"
							} else {
								roomName = "Johnny's Room"
							}
						case "desk":
							roomName = "Study"
						case "sofa", "tv":
							roomName = "Living Room"
						case "stove", "fridge", "sink":
							roomName = "Kitchen"
						}
					}

					// Check for Johnny's door
					if room.Grid[x][y].ObstacleName == "johnny's door" {
						robot.World.UpdateDoorStatus("johnny's door", true)
					}
				}
			}
		}

		if roomName == "" {
			roomName = fmt.Sprintf("Room %d", i)
		}

		roomNameToIndex[roomName] = i
		assignedRooms[roomName] = true // FIX: mark room as assigned
		fmt.Printf("Identified room %s (index %d)\n", roomName, i)
	}

	// Scan for objects to update logical world
	fmt.Println("Robot is scanning the house for objects...")
	for _, room := range house.Rooms {
		for x := range room.Width {
			for y := range room.Height {
				cell := room.Grid[x][y]
				if cell.Type == "furniture" && cell.ObstacleName != "" {
					robot.World.UpdateObjectFound(cell.ObstacleName)
				}
			}
		}
	}

	return roomNameToIndex
}
