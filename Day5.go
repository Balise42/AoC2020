package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (day *Day) Day5a() {
	fmt.Printf("Part 1: %d\n", ComputeDay5a(day.input))
}

func (day *Day) Day5b() {
	fmt.Printf("Part 2: %d\n", ComputeDay5b(day.input))
}

func ComputeDay5a(input string) int {
	lines := strings.Split(input, "\n")
	maxId := 0
	for _, line := range lines {
		newId := ComputeSeatId(line)
		if newId > maxId {
			maxId = newId
		}
	}
	return maxId
}

func ComputeDay5b(input string) int {
	maxId := ComputeDay5a(input)
	seating := make([]bool, maxId+1)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		seating[ComputeSeatId(line)] = true
	}

	started := false

	for i, b := range seating {
		if b {
			started = true
		} else {
			if started {
				return i
			}
		}
	}
	panic ("Could not find seat!")
}

// the seat id is just binary with FBLR replacing 0101
func ComputeSeatId(line string) int {
	conv := strings.ReplaceAll(line, "F", "0")
	conv = strings.ReplaceAll(conv, "B", "1")
	conv = strings.ReplaceAll(conv, "R", "1")
	conv = strings.ReplaceAll(conv, "L", "0")
	if ret, err := strconv.ParseInt(conv, 2, 32); err == nil {
		return int(ret)
	}
	panic("Could not convert ID")
}
