package main

import (
	"fmt"
	"strings"
)

func (day *Day) Day6a() {
	fmt.Printf("Part 1: %d\n", ComputeDay6a(day.input))
}

func (day *Day) Day6b() {
	fmt.Printf("Part 2: %d\n", ComputeDay6b(day.input))
}

func ComputeDay6a(input string) int {
	groups := strings.Split(input, "\n\n")

	res := 0

	for _, group := range groups {
		answers := make(map[int32]bool)
		persons := strings.Split(group, "\n")
		for _, person := range persons {
			for _, letter := range person {
				answers[letter] = true
			}
		}
		res = res + len(answers)
	}
	return res
}

func ComputeDay6b(input string) int {
	groups := strings.Split(input, "\n\n")

	res := 0

	for _, group := range groups {
		answers := make(map[int32]int)
		persons := strings.Split(group, "\n")
		for _, person := range persons {
			for _, letter := range person {
				if v, ok := answers[letter]; ok {
					answers[letter] = v + 1
				} else {
					answers[letter] = 1
				}
			}
		}
		for _, v := range answers {
			if v == len(persons) {
				res++
			}
		}
	}
	return res
}
