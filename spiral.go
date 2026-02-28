package main

import (
	"time"
)

func CleanRoomSpiral(room *Room, rumba *Robot) {
	startTime := time.Now()
	moveCount := 0

	// Find the center of the room
	centerX := room.Width / 2
	centerY := room.Height / 2

	// Find a valid point near the center
	centerPoint := findNearestCleanablePoint(room, Point{X: centerX, Y: centerY})

	// Find path to center (A*)
	pathToCenter := Astar(room, rumba.Position, centerPoint)
	// Move to the center
	if len(pathToCenter) > 1 {
		for i := 1; i < len(pathToCenter); i++ {
			rumba.Position = pathToCenter[i]
			rumba.Path = append(rumba.Path, rumba.Position)
			Clean(rumba, room)
			if room.Animate {
				room.Display(rumba, false)
				time.Sleep(moveDelay)
			}
			moveCount++
		}
	}

	// Create a spiral pattern
	spiralPoints := generateSpiralPattern(room, centerPoint)

	// Follow the spiral pattern (for loop)

	//	// Skip if cell is already cleaned or obstacle
	//	// Find path to the next point
	//	// Move along the path

	// end for

	// Final Cleanup

	// Calculate cleaning time
	cleaningTime := time.Since(startTime)

	// Display final result
	displaySummary(room, rumba, moveCount, cleaningTime)

}

func findNearestCleanablePoint(room *Room, targetPoint Point) Point {

	if room.IsValid(targetPoint.X, targetPoint.Y) && !room.Grid[targetPoint.X][targetPoint.Y].Obstacle {
		return targetPoint
	}

	// Search for valid Point in expanding squares
	for radius := 1; radius < room.Width || radius < room.Height; radius++ {
		// Check all points in the current radius
		for dx := -radius; dx <= radius; dx++ {
			for dy := -radius; dy <= radius; dy++ {
				if Abs(dx) != radius && Abs(dy) != radius {
					continue
				}
				x, y := targetPoint.X+dx, targetPoint.Y+dy

				// Check if this point valid

				if room.IsValid(x, y) && !room.Grid[x][y].Obstacle {
					return Point{X: x, Y: y}
				}
			}
		}
	}
	// if no valid point return the starting point
	return Point{X: 1, Y: 1}
}

func generateSpiralPattern(room *Room, center Point) []Point {
	var points []Point

	// Maximal possible spiral size

	maxSize := max(room.Width, room.Height)

	// Set deltax and deltay
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}

	// Start at center
	x, y := center.X, center.Y
	dir := 0 // Start moving right

	// Set spiral parameters
	step := 1
	stepCount := 0
	dirChanges := 0

	// Generate Spiral pattern

	for range maxSize * maxSize {
		// Add current point if valid
		if room.IsValid(x, y) {
			points = append(points, Point{X: x, Y: y})

		}

		// Take a step
		x += dx[dir]
		y += dy[dir]
		stepCount++

		// Check to see if we need to change directions
		if stepCount == step {
			dir = (dir + 1) % 4
			stepCount = 0
			dirChanges++

			// Increase step size after every 2 direction changes
			if dirChanges == 2 {
				step++
				dirChanges = 0
			}
		}

		// Break if we are out of bounds
		if x < 0 || x >= room.Width || y < 0 || y >= room.Height {
			break
		}
	}

	return points

}
