package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (day *Day) Day15a() {
	fmt.Printf("Part 1: %d\n", ComputeDay15a("6,4,12,1,20,0,16"))
}

func (day *Day) Day15b() {
	fmt.Printf("Part 2: %d\n", ComputeDay15b("6,4,12,1,20,0,16"))
}

func ComputeDay15a(input string) int {
	return playMemory(input, 2020)
}

func playMemory(input string, rounds int) int {
	numbers := make([]int, 0)
	numbers = append(numbers, -1)
	memo := make([][]int, rounds + 2)

	for i, tok := range strings.Split(input, ",") {
		num, err := strconv.Atoi(tok)
		if err != nil {
			panic("Could not parse input " + input)
		}
		numbers = append(numbers, num)
		memo[num] = append(memo[num], i + 1)
	}

	for len(numbers) <= rounds + 1 {
		lastSeen := memo[numbers[len(numbers) - 1]]
		if len(lastSeen) < 2 {
			numbers = append(numbers, 0)
			memo[0] = append(memo[0], len(numbers) - 1)
		} else {
			val := len(numbers) - 1 -lastSeen[len(lastSeen) - 2]
			numbers = append(numbers, val)
			memo[val] = append(memo[val], len(numbers) - 1)
		}
	}
	return numbers[rounds]
}

func ComputeDay15b(input string) int {
	return playMemory(input, 30000000)
}
