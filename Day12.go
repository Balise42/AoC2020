package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day12 struct {
	posX int
	posY int
	dir uint8
}

type Day12b struct {
	posX int
	posY int
	dirX int
	dirY int
}

func (day *Day) Day12a() {
	fmt.Printf("Part 1: %d\n", ComputeDay12a(day.input))
}

func (day *Day) Day12b() {
	fmt.Printf("Part 2: %d\n", ComputeDay12b(day.input))
}

func ComputeDay12a(input string) int {
	state := &Day12{0, 0, 'E'}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		moveShip(state, line)
	}
	return int(math.Abs(float64(state.posX)) + math.Abs(float64(state.posY)))
}

func moveShip(state *Day12, line string) {
	val, err := strconv.Atoi(line[1:])
	if err != nil {
		panic("Could not interpret value for " + line)
	}
	if line[0] == 'F' {
		advanceShip(state, state.dir, val)
	} else if line[0] == 'N' || line[0] == 'S' || line[0] == 'E' || line[0] == 'W' {
		advanceShip(state, line[0], val)
	} else if line[0] == 'L' {
		rotateShip(state, -val)
	} else if line[0] == 'R' {
		rotateShip(state, val)
	} else {
		panic("Could not interpret instruction for " + line)
	}
}

func advanceShip(state *Day12, dir uint8, val int) {
	if dir == 'N' {
		state.posY -= val
	} else if dir == 'S' {
		state.posY += val
	} else if dir == 'E' {
		state.posX += val
	} else if dir == 'W' {
		state.posX -= val
	}
}

func getRotation(val int) (int, int) {
	var dir int
	var numTurns int
	if val < 0 {
		dir = -1
		numTurns = -val / 90
	} else {
		dir = 1
		numTurns = val / 90
	}
	return numTurns, dir
}

func rotateShip(state *Day12, val int) {
	cardinals := "ESWN"

	numTurns, dir := getRotation(val)

	pos := strings.Index(cardinals, string(state.dir))
	pos = pos + dir * numTurns
	for pos < 0 {
		pos += 4
	}

	state.dir = cardinals[pos % 4]
}


func ComputeDay12b(input string) int {
	state := &Day12b{0, 0, 10, -1}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		updateDay12State(state, line)
	}
	return int(math.Abs(float64(state.posX)) + math.Abs(float64(state.posY)))
}

func updateDay12State(state *Day12b, line string) {
	val, err := strconv.Atoi(line[1:])
	if err != nil {
		panic("Could not interpret value for " + line)
	}
	if line[0] == 'F' {
		state.posX += val * state.dirX
		state.posY += val * state.dirY
	}
	if line[0] == 'N' {
		state.dirY -= val
	} else if line[0] == 'S' {
		state.dirY += val
	} else if line[0] == 'E' {
		state.dirX += val
	} else if line[0] == 'W' {
		state.dirX -= val
	} else if line[0] == 'L' {
		rotateWaypoint(state, -val)
	} else if line[0] == 'R' {
		rotateWaypoint(state, val)
	}
}

func rotateWaypoint(state *Day12b, val int) {
	numTurns, dir := getRotation(val)
	for i := 0; i<numTurns; i++ {
		state.dirX, state.dirY = state.dirY, state.dirX
		if dir < 0 {
			state.dirY = -state.dirY
		} else {
			state.dirX = -state.dirX
		}
	}
}
