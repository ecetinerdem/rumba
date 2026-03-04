# Rumba House Cleaner Simulation

A simple grid-based robot vacuum cleaner simulation written in **Go**.  
Supports single-room cleaning, multi-room houses, different movement algorithms, optional cat interference, and propositional logic-based decision making.

## Features

- Multiple cleaning algorithms: random walk, spiral, snake, SLAM-style
- Single room or multi-room house mode (via JSON config)
- Optional animated terminal visualization
- Cat that can block or interfere with cleaning (🐱)
- Propositional logic mode:
  - Detects people home via objects (backpack → Jack, bicycle → Sarah, skateboard → Johnny)
  - Respects closed doors (Johnny's room)
  - Smart room priority based on who’s home + weekday rules
- Realistic progress tracking and cleaning summary

## Usage

```bash
go run main.go [flags]
```

-file string
      configuration file (default "empty.json")
-algorithm string
      cleaning algorithm: random, spiral, snake, slam (default "random")
-animate
      animate cleaning process (default true)
-cat
      add a cat to the room(s)
-house
      treat config file as multi-room house (default: single room)
-logic
      use propositional logic for room priority and decisions

# Basic single room + animation + spiral
go run main.go -file room.json -algorithm spiral

# With cat + random walk
go run main.go -file room.json -cat

# Multi-room house with logic-based priority
go run main.go -house -file house.json -logic -algorithm snake

# Quiet cleaning (no animation)
go run main.go -file room.json -animate=false

Configuration Files

Single room: room.json
Example:JSON{
  "width": 400,
  "height": 300,
  "furniture": [
    {"x": 100, "y": 80, "width": 120, "height": 180, "name": "bed", "type": "bed"}
  ]
}
Multi-room house: array of room objects (when using -house)

Logic Mode Rules (when -logic is used)

Backpack → Jack is home
Bicycle → Sarah is home
Skateboard → Johnny is home
Johnny's door closed → skip Johnny's room
Sarah home → skip Living Room
Jack home + weekday → clean Jack's room last
Nobody home → clean everything (kitchen first)

Requirements

Go 1.18+

Todo / Possible Improvements

Better JSON schema validation
More furniture/object types for logic
Battery simulation
Dirt amount variation
Save path as image or animation

Enjoy your virtual vacuuming! 🧹🤖🐱


      
