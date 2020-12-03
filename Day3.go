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
	for i := 0; i*yinc < len(lines); i++ {
		if lines[i*yinc][i*xinc%len(lines[i*yinc])] == '#' {
			numtrees++
		}
	}
	return uint64(numtrees)
}

func ComputeDay3b(input string) uint64 {
	lines := strings.Split(input, "\n")
	return NumTrees(lines, 1, 1) * NumTrees(lines, 3, 1) * NumTrees(lines, 5, 1) * NumTrees(lines, 7, 1) * NumTrees(lines, 1, 2)
}
