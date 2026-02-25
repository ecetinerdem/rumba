package main

var directions = [][]int{
	{0, -1}, // North
	{1, 0},  // East
	{0, 1},  // South
	{-1, 0}, // West
}

type Robot struct {
	Position             Point
	Path                 []Point
	CleanRoom            func(*Room, *Robot)
	Direction            float64
	ObstaclesEncountered map[string]bool
}

func NewRobot(startX, startY int) *Robot {
	return &Robot{
		Position: Point{
			X: startX,
			Y: startY,
		},
		Path:                 []Point{{X: startX, Y: startY}},
		ObstaclesEncountered: make(map[string]bool),
	}
}

func Clean(rumba *Robot, room *Room) {
	x, y := rumba.Position.X, rumba.Position.Y

	if !room.Grid[x][y].Cleaned && !room.Grid[x][y].Obstacle {
		room.Grid[x][y].Cleaned = true
		room.Grid[x][y].Type = "clean"
		room.CleanedCellCount++
	}

	CheckAdjacentObsticle(rumba, room)
}

func CheckAdjacentObsticle(rumba *Robot, room *Room) {
	x, y := rumba.Position.X, rumba.Position.Y
	for _, direction := range directions {
		newX := x + direction[0]
		newY := y + direction[1]
		RecordObstacle(rumba, room, newX, newY)
	}
}

func RecordObstacle(rumba *Robot, room *Room, x, y int) {
	if x >= 0 && x < room.Width && y >= 0 && y <= room.Height && room.Grid[x][y].Obstacle {
		if room.Grid[x][y].Type == "furniture" && room.Grid[x][y].ObstacleName != "" {
			rumba.ObstaclesEncountered[room.Grid[x][y].ObstacleName] = true
		}
	}
}
