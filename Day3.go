package main

import (
	"fmt"
	"strings"
)

func (day *Day) Day3a() {
	fmt.Printf("Part 1: %d\n", ComputeDay3a(day.input))
}

func (day *Day) Day3b() {
	fmt.Printf("Part 2: %d\n", ComputeDay3b(day.input))
}

func ComputeDay3a(input string) uint64 {
	lines := strings.Split(input, "\n")
	return NumTrees(lines, 3, 1)
}

func NumTrees(lines []string, xinc int, yinc int) uint64 {
	numtrees := 0
	for i := 0; i < len(lines) / yinc; i++ {
		if lines[i*yinc][i*xinc % len(lines[i*yinc])] == '#' {
			numtrees++
		}
	}
	// if the number of lines is not divisible by the number of steps we're off by one, so we re-add it
	if len(lines) / yinc * yinc < len(lines) {
		i := len(lines) / yinc
		if lines[i*yinc][i*xinc % len(lines[i*yinc])] == '#' {
			numtrees++
		}
	}
	println(numtrees)
	return uint64(numtrees)
}

func ComputeDay3b(input string) uint64 {
	lines := strings.Split(input, "\n")
	return NumTrees(lines, 1, 1) * NumTrees(lines, 3, 1) * NumTrees(lines, 5, 1) * NumTrees(lines, 7, 1) * NumTrees(lines, 1, 2)
}



