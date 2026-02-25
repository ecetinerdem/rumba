package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

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
	// Load from JSON config
	roomConfig, err := LoadroomConfig(configFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Convert dimensions to grid cells

	gridWith := roomConfig.Width / cellSize
	gridHeight := roomConfig.Height / cellSize

	// Create grid
	grid := make([][]Cell, gridWith)
	for i := range grid {
		grid[i] = make([]Cell, gridHeight)
		for j := range grid[i] {
			grid[i][j] = Cell{Type: "dirty", Cleaned: false, Obstacle: false}
		}
	}
	// Add walls

	for i := range gridWith {
		grid[i][0] = Cell{Type: "wall", Cleaned: true, Obstacle: true, ObstacleName: "wall"}
		grid[i][gridHeight-1] = Cell{Type: "wall", Cleaned: false, Obstacle: true, ObstacleName: "wall"}

	}

	for j := range gridHeight {
		grid[0][j] = Cell{Type: "wall", Cleaned: false, Obstacle: true, ObstacleName: "wall"}
		grid[gridWith-1][j] = Cell{Type: "wall", Cleaned: false, Obstacle: true, ObstacleName: "wall"}
	}
	// Add furniture
	// To do

	// Count cleanable cells
	cleanableCellCount := 0

	for i := range gridWith {
		for j := range gridHeight {
			if !grid[i][j].Obstacle {
				cleanableCellCount++
			}
		}
	}
	return &Room{
		Width:              gridWith,
		Height:             gridHeight,
		Grid:               grid,
		CleanableCellCount: cleanableCellCount,
		CleanedCellCount:   0,
		Animate:            animate,
	}
}

func LoadroomConfig(fileName string) (*RoomConfig, error) {
	jsonData, err := os.ReadFile(fileName)

	if err != nil {
		return nil, fmt.Errorf("error reading json file: %v", err)
	}

	var config RoomConfig

	if err := json.Unmarshal(jsonData, &config); err != nil {
		return nil, fmt.Errorf("error parsing json file: %v", err)
	}

	return &config, nil
}
