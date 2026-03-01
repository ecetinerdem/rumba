package main

import "math/rand/v2"

type Cat struct {
	Position   Point
	Active     bool
	StopTimer  int
	Path       []Point
	DirectionX int
	DirectionY int
}

func NewCat(room *Room) *Cat {

	var startX, startY int

	for {
		startX = rand.IntN(room.Width-4) + 2
		startY = rand.IntN(room.Height-4) + 2

		if !room.Grid[startX][startY].Obstacle {
			break
		}
	}

	return &Cat{
		Position:   Point{X: startX, Y: startY},
		Active:     true,
		StopTimer:  0,
		Path:       []Point{{X: startX, Y: startY}},
		DirectionX: rand.IntN(3) - 1,
		DirectionY: rand.IntN(3) - 1,
	}
}

func MoveCat(cat *Cat, room *Room) {
	if cat != nil {
		// Check to see if the cat is currently stopped
		if cat.StopTimer > 0 {
			cat.StopTimer--

		}

		// Chance for cat to stop
		if rand.Float64() < catStopProbability {
			cat.StopTimer = catStopDuration
			return
		}

		// Change direction randomly
		if rand.Float64() < 0.2 {
			cat.DirectionX = rand.IntN(3) - 1
			cat.DirectionY = rand.IntN(3) - 1
		}

		// Make sure cat doesn't just stay still when moving
		if cat.DirectionX == 0 && cat.DirectionY == 0 {
			cat.DirectionX = rand.IntN(3) - 1
			if cat.DirectionX == 0 {
				cat.DirectionY = rand.IntN(2)*2 - 1
			} else {
				cat.DirectionY = rand.IntN(3) - 1
			}
		}
		// Calculate the new position

		newX := cat.Position.X + cat.DirectionX
		newY := cat.Position.Y + cat.DirectionY

		if room.IsValid(newX, newY) {
			cat.Position = Point{X: newX, Y: newY}
			cat.Path = append(cat.Path, cat.Position)
		}

		if room.Grid[newX][newY].Cleaned {
			room.Grid[newX][newY].Cleaned = false
			room.Grid[newX][newY].Type = "dirty"
			room.CleanedCellCount--

		}
	}
}

func IsAdjacentToCat(rumba *Robot, cat *Cat) bool {
	if cat != nil {
		dx := Abs(rumba.Position.X - cat.Position.X)
		dy := Abs(rumba.Position.Y - cat.Position.Y)

		return dx <= 1 && dy <= 1
	}

	return false
}
