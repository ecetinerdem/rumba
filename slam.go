package main

import (
	"math"
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
	for len(frontier) > 0 && room.CleanedCellCount < room.CleanableCellCount {
		// Get closest frontier point
		target := getClosestFrontierPoint(rumba.Position, frontier)

		// If not valid break

		if target.X == -1 && target.Y == -1 {
			break
		}
		// REmove target from frontier
		delete(frontier, target)
		// Find path to target use A*
		path := Astar(room, rumba.Position, target)

		// If not path found, go to next frontier point(continue)
		if len(path) <= 1 {
			continue
		}
		// Move along the path
		for i := 1; i < len(path); i++ {
			// Update rumba position
			rumba.Position = path[i]
			rumba.Path = append(rumba.Path, rumba.Position)
			// Clean Position
			Clean(rumba, room)

			// Mark as visited
			visited[rumba.Position] = true

			// Update map(internal) based upan what we can see from current position
			updateRumbaMap(rumba.Position, rumbaMap, room)

			// Update frontier with newly discovered cells
			addNeighboursToFrontier(rumba.Position, rumbaMap, frontier, visited, room)

			// Display the state
			if room.Animate {
				room.Display(rumba, false)
				time.Sleep(moveDelay)
			}
			moveCount++
		}

		// every 10 moves, do a more thorough frontier check
		if moveCount%10 == 0 {
			updateAllFrontiers(rumbaMap, frontier, visited, room)
		}
		// Check if we have sufficient coverage - break
		if float64(room.CleanedCellCount)/float64(room.CleanableCellCount) > 0.95 {
			break
		}
	}
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

func getClosestFrontierPoint(position Point, frontier map[Point]bool) Point {
	closestPoint := Point{X: -1, Y: -1}
	minDistance := math.MaxFloat64

	for point := range frontier {
		distance := heuristic(position, point)
		if distance < minDistance {
			minDistance = distance
			closestPoint = point
		}
	}
	return closestPoint
}

func updateAllFrontiers(rumbaMap [][]int, frontier map[Point]bool, visited map[Point]bool, room *Room) {
	for x := 1; x < room.Width-1; x++ {
		for y := 1; y < room.Height-1; y++ {
			// If a cell free but not visited add to frontier
			point := Point{X: x, Y: y}
			if rumbaMap[x][y] == 1 && !visited[point] && !frontier[point] && !room.Grid[x][y].Obstacle {
				// Check to see if it is accesiable. At least has one visited neighbour
				for _, dir := range directions {
					nx, ny := x+dir[0], y+dir[1]
					neighbourPoint := Point{x: nx, Y: ny}
					if nx >= 0 && nx < room.Width && ny < room.Height && ny >= 0 && visited[neighbourPoint] {
						frontier[point] = true
						break
					}
				}
			}
		}
	}
}
