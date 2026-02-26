package main

import (
	"math"
	"time"
)

func CleanRandomWalk(room *Room, rumba *Robot) {
	startTime := time.Now()
	moveCount := 0

	// Set variables
	maxMoves := room.Width * room.Height * 5
	stuckCount := 0
	maxStuckCount := 5 // Max consecutive failed moves before changing strategy

	// Clean Current Position
	Clean(rumba, room)

	if room.Animate {
		room.Display(rumba, false)
		time.Sleep(moveDelay)
	}

	for moveCount < maxMoves && room.CleanedCellCount < room.CleanableCellCount {
		// Generate a random angle in radians

		// Calculate a direction vector based on angle

		// Use Bresenham's line algorithm to move in that direction until hitting an obsticle

		// If we did not move very much, increment stuck counter and possibly change strategy

		// If stuck too many times, use a* to find path to neares dirty cell

		// Add some adaptive behaviour. Scan for dirty cells every once in a while

		// end for

		// Final sweep to complete coverage
	}

	//Calculate cleaning time

	cleaningTime := time.Since(startTime)
	displaySummary(room, rumba, moveCount, cleaningTime)
}

func bresenhamLine(x0, y0, x1, y1 int) []Point {
	// Initialize a slice to store all points on the line
	var points []Point
	// Calculate the absolute difference between the endpoints
	dx := Abs(x1 - x0)
	dy := Abs(y1 - y0)
	// Determin the direction of the movement along each axis
	sx := -1
	if x0 < x1 {
		sx = 1
	}
	sy := -1
	if y0 < y1 {
		sy = 1
	}
	// Calculate initial error value
	err := dx - dy
	// for loop until we reach endpoint
	for {
		// Add the current point to our result
		points = append(points, Point{X: x0, Y: y0})

		// Check to see if reached the endpoint
		if x0 == x1 && y0 == y1 {
			break
		}

		// Calculate the error for the next step
		e2 := 2 * err

		// If moving in the x direction would be keep us closer to the ideal line
		if e2 > -dy {
			// If we reached the endpoint
			if x0 == x1 {
				break
			}
			//Update the error and move in the x-direction
			err -= dy
			x0 += sx
		}

		// If moving in the y direction would be keep us closer to the ideal line
		if e2 < -dx {
			if y0 == y1 {
				break
			}
			//Update the error and move in the y-direction
			err -= dx
			y0 += sy
		}

	}

	// return points
	return points
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func findNearestDirtyCell(room *Room, position Point) Point {
	var nearestCell Point = Point{X: -1, Y: -1}
	minDistance := math.MaxFloat64

	for i := 1; i < room.Width-1; i++ {
		for j := 1; j < room.Height-1; j++ {
			distance := heuristic(position, Point{X: i, Y: j})
			if distance < minDistance {
				minDistance = distance
				nearestCell = Point{X: i, Y: j}
			}
		}
	}
	return nearestCell
}
