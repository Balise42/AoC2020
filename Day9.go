package main

import (
	"fmt"
	"math"
)

func (day *Day) Day9a() {
	fmt.Printf("Part1: %d\n", ComputeDay9a(day.input, 25))
}

func (day *Day) Day9b() {
	fmt.Printf("Part2: %d\n", ComputeDay9b(day.input, 25))
}

func ComputeDay9a(input string, window int) int {
	intList := createListAsInts(input)

	for i := window; i<len(intList); i++ {
		found := false
		for j := i - window; j < i; j++ {
			for k := j+1; k<i; k++ {
				if intList[j] + intList[k] == intList[i] {
					found = true
					break
				}
			}
		}
		if !found {
			return intList[i]
		}
	}
	panic("Could not find solution!")
}

func ComputeDay9b(input string, window int) int {
	sum := ComputeDay9a(input, window)

	intList := createListAsInts(input)

	for i := 0; i < len(intList) - 1; i++ {
		end, found := findEnd(intList[i:], sum)
		if found {
			min, max := getMinMax(intList[i:i+end+1])
			return min + max
		}
	}

	panic("Could not find solution!")
}

func getMinMax(ints []int) (int, int) {
	min := math.MaxInt64
	max := 0
	for _, v := range ints {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func findEnd(seq []int, sum int) (int, bool) {
	partialSum := seq[0]
	for i := 1; i < len(seq); i++ {
		partialSum += seq[i]
		if partialSum == sum {
			return i, true
		}
	}
	return -1, false
}