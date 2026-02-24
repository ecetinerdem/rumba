package main

import "time"

const (
	// Display characters
	charRobot          = "🔴"
	charWall           = "🟦"
	charFurniture      = "🪑"
	charClean          = "🧼"
	charDirty          = "🟫"
	charPath           = "🟢"
	charCat            = "🐱" // Display character for cat
	catStopProbability = 0.1 // Probability of cat stopping
	catStopDuration    = 5   // Duration cat stays still (in animation frames)
	moveDelay          = 50 * time.Millisecond
	cellSize           = 10
)

type Point struct {
	x, y int
}

type Cell struct {
	Type         string // wall, furniture, clean, dirty, bike
	Cleaned      bool
	Obstacle     bool
	ObstacleName string
}

type Furniture struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Name   string `json:"name"`
	Type   string `json:"type"`
}

type Room struct {
	Grid               [][]Cell
	Width              int
	Height             int
	CleanableCellCount int
	CleanedCellCount   int
	Animate            bool
}

type RoomConfig struct {
	Width     int         `json:"width"`
	Height    int         `json:"height"`
	Furniture []Furniture `json:"furniture"`
}

func NewRoom(configFile string, animate bool) *Room {
	return &Room{}
}
