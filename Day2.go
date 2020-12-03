package main

import (
	"fmt"
	"strconv"
	"strings"
)

type day2 struct {
	min      int
	max      int
	letter   uint8
	password string
}

func (day *Day) Day2a() {
	fmt.Printf("Part1: %d\n", ComputeDay2a(day.input))
}

func (day *Day) Day2b() {
	fmt.Printf("Part2: %d\n", ComputeDay2b(day.input))
}

func ComputeDay2a(input string) int {
	lines := strings.Split(input, "\n")
	valid := 0
	for _, line := range lines {
		data := parseDay2(line)
		if isValidPart1(data) {
			valid++
		}
	}
	return valid
}

func ComputeDay2b(input string) int {
	lines := strings.Split(input, "\n")
	valid := 0
	for _, line := range lines {
		data := parseDay2(line)
		if isValidPart2(data) {
			valid++
		}
	}
	return valid
}

func parseDay2(line string) day2 {
	tokens := strings.Split(line, " ")
	interval := tokens[0]
	letter := tokens[1][0]
	password := tokens[2]

	intervalBounds := strings.Split(interval, "-")
	min, err := strconv.Atoi(intervalBounds[0])
	if err != nil {
		panic("Could not parse interval")
	}
	max, err := strconv.Atoi(intervalBounds[1])
	if err != nil {
		panic("Could not parse interval")
	}
	return day2{min, max, letter, password}
}

func isValidPart1(data day2) bool {
	numOcc := strings.Count(data.password, string(data.letter))
	return numOcc >= data.min && numOcc <= data.max
}

func isValidPart2(data day2) bool {
	l1 := data.password[data.min-1]
	l2 := data.password[data.max-1]
	return (l1 == data.letter || l2 == data.letter) && !(l1 == data.letter && l2 == data.letter)
}
