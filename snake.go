package main

import "time"

func CleanRoomSnake(room *Room, rumba *Robot) {
	// Initialize start time and start count
	startTime := time.Now()
	moveCount := 0

	// Generate the snaking pattern
	coveragePoints := generateSneakingPattern(room)

	// Clean cell
	Clean(rumba, room)

	if room.Animate {
		room.Display(rumba, false)
		time.Sleep(moveDelay)
	}

	// Visit each point in coverage pattern(for)
	for _, point := range coveragePoints {
		// Skip cell if already cleaned
		if room.Grid[point.X][point.Y].Cleaned {
			continue
		}
		// Find path to next point
		path := Astar(room, rumba.Position, point)

		// If no path found try next point
		if len(path) == 0 {
			continue
		}

		// Move along the path(for)
		for i := 1; i < len(path); i++ {
			// Update robot position
			rumba.Position = path[i]
			rumba.Path = append(rumba.Path, path[i])

			// Clean
			Clean(rumba, room)
			// Display room
			if room.Animate {
				room.Display(rumba, false)
				time.Sleep(moveDelay)
			}
			// Increment movecount
			moveCount++
		}

	}

	// Do final sweep
	finalCleanUp(room, rumba, &moveCount)

	// Calculate cleaning time
	cleaningTime := time.Since(startTime)
	// Display final result
	displaySummary(room, rumba, moveCount, cleaningTime)

}

func generateSneakingPattern(room *Room) []Point {

	var points []Point

	var directionX = 1

	for y := 1; y < room.Height-1; y++ {
		if directionX == 1 {
			// Move left to right
			for x := 1; x < room.Width-1; x++ {
				if !room.Grid[x][y].Obstacle {
					points = append(points, Point{X: x, Y: y})
				}
			}
		} else {
			// Moving right to left

			for x := room.Width - 2; x >= 1; x-- {
				if !room.Grid[x][y].Obstacle {
					points = append(points, Point{X: x, Y: y})
				}
			}
		}
		directionX *= -1
	}

	return points
}
