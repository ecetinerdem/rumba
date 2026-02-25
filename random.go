package main

import "time"

func CleanRandomWalk(room *Room, rumba *Robot) {
	startTime := time.Now()
	moveCount := 0

	//Calculate cleaning time

	cleaningTime := time.Since(startTime)
	displaySummary(room, rumba, moveCount, cleaningTime)
}
