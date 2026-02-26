package main

import (
	"time"
)

func CleanRoom(room *Room, rumba *Robot) {

	// Set start time and moveCount
	startTime := time.Now()
	moveCount := 0

	// Initialize rumba's internal map
	rumbaMap := initializeRobotMap(room.Width, room.Height)

	// Initialize visired cells (tracking)
	visited := make(map[Point]bool)

	// Initialize a frontier
	frontier := make(map[Point]bool)

	// Mark starting position as visited and update internal map for the first time
	visited[rumba.Position] = true
	updateRumbaMap(rumba.Position, rumbaMap, room)
	// Clean the current position
	Clean(rumba, room)

	// Add neighbours to the frontier
	addNeighboursToFrontier(rumba.Position, rumbaMap, frontier, visited, room)

	// Display initial state
	if room.Animate {
		room.Display(rumba, false)
		time.Sleep(moveDelay)
	}

	// for - if the frontier not empty and the room is not  %100 clean

	//	// Get closest frontier point
	//	// If not valid break
	//	// REmove target from frontier
	//	// Find path to target use A*
	//	// If not path found, go to next frontier point(continue)
	//	// Move along the path
	//		// Update map(internal)
	//	// every 10 moves, do a more thorough frontier check
	//	// Check if we have sufficient coverage - break

	// Final cleanup phase
}

func initializeRobotMap(width, height int) [][]int {
	// 0 = unknown
	// 1 = free
	// 2 = obstacle
	// 3 = cleaned

	robotMap := make([][]int, width)

	for i := range robotMap {
		robotMap[i] = make([]int, height)
	}

	return robotMap
}

func updateRumbaMap(position Point, rumbaMap [][]int, room *Room) {
	if room.Grid[position.X][position.Y].Cleaned {
		rumbaMap[position.X][position.Y] = 3
	} else {
		rumbaMap[position.X][position.Y] = 1
	}

	// Scan surroundings

	for _, dir := range directions {
		newX, newY := position.X+dir[0], position.Y+dir[1]
		// Check position in bounds
		if newX >= 0 && newX < len(rumbaMap) && newY >= 0 && newY < len(rumbaMap[0]) {
			if room.Grid[newX][newY].Obstacle {
				rumbaMap[newX][newY] = 2
			} else if rumbaMap[newX][newY] == 0 {
				rumbaMap[newX][newY] = 1
			} else if room.Grid[newX][newY].Cleaned {
				rumbaMap[newX][newY] = 3
			}
		}
	}
}

func addNeighboursToFrontier(position Point, rumbaMap [][]int, frontier map[Point]bool, visitedmap map[Point]bool, room *Room) {
	// Check adjacent cells

	for _, dir := range directions {
		newX, newY := position.X+dir[0], position.Y+dir[1]
		newPoint := Point{X: newX, Y: newY}

		// Check if position valid, not visited and not an obstacle and not in frontier
		if newX >= 0 && newX <= len(rumbaMap) && newY >= 0 && newY < len(rumbaMap[0]) &&
			!visitedmap[newPoint] && !frontier[newPoint] && room.IsValid(newX, newY) {
			// Add to frontier
			frontier[newPoint] = true
		}

	}
}
