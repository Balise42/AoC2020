package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (day *Day) Day8a() {
	fmt.Printf("Part1: %d\n", ComputeDay8a(day.input))
}

func (day *Day) Day8b() {
	fmt.Printf("Part2: %d\n", ComputeDay8b(day.input))
}

func ComputeDay8a(input string) int {
	loc := strings.Split(input, "\n")
	ret, _ := runBootCode(loc)
	return ret
}

func ComputeDay8b(input string) int {
	loc := strings.Split(input, "\n")
	for i, line := range loc {
		if line[0:3] == "nop" {
			loc[i] = strings.Replace(loc[i], "nop", "jmp", 1)
			ret, terminates := runBootCode(loc)
			if terminates {
				return ret
			}
			loc[i] = strings.Replace(loc[i], "jmp", "nop", 1)
		} else if line[0:3] == "jmp" {
			loc[i] = strings.Replace(loc[i], "jmp", "nop", 1)
			ret, terminates := runBootCode(loc)
			if terminates {
				return ret
			}
			loc[i] = strings.Replace(loc[i], "nop", "jmp", 1)
		}
	}
	panic("Could not find solution")
}

func runBootCode(loc []string) (int, bool) {
	visited := make([]bool, len(loc))
	ptr := 0
	acc := 0
	for ptr < len(loc) {
		if visited[ptr] {
			return acc, false
		}
		visited[ptr] = true
		line := loc[ptr]
		if line[0:3] == "nop" {
			ptr++
		} else if line[0:3] == "acc" {
			ptr++
			val, err := strconv.Atoi(line[4:])
			if err != nil {
				panic("Cannot parse new accumulator value")
			}
			acc += val
		} else if line [0:3] == "jmp" {
			val, err := strconv.Atoi(line[4:])
			if err != nil {
				panic("Cannot parse jmp value")
			}
			ptr += val
		}
	}
	return acc, true
}