package main

import (
	"fmt"
	"sort"
	"strconv"
)

func (day *Day) Day10a() {
	fmt.Printf("Part 1: %d\n", ComputeDay10a(day.input))
}

func (day *Day) Day10b() {
	fmt.Printf("Part 2: %d\n", ComputeDay10b(day.input))
}

func ComputeDay10a(input string) int {
	list := createListAsInts(input)
	list = append(list, 0)
	sort.Ints(list)

	diff1 := 0
	// we have one diff3 with the last adapter in the bag
	diff3 := 1

	for i := 1; i<len(list); i++ {
		if list[i] == list[i-1] + 1 {
			diff1++
		} else if list[i] == list[i-1] + 3 {
			diff3++
		}
	}

	return diff1 * diff3
}

var memo map[string]uint64

func ComputeDay10b(input string) uint64 {
	memo = make(map[string]uint64, 0)

	list := createListAsInts(input)
	list = append(list, 0)
	sort.Ints(list)
	list = append(list, list[len(list) - 1] + 3)

	return computeAdapterArrangements(list)
}



func computeAdapterArrangements(list []int) uint64 {
	key := listToString(list)
	v, ok := memo[key]; if ok {
		return v
	}

	if len(list) <= 2 {
		return 1
	}

	var ret uint64

	if len(list) >= 4 && list[3] - list[0] <= 3 {
		ret = computeAdapterArrangements(list[1:]) + computeAdapterArrangements(list[2:]) + computeAdapterArrangements(list[3:])
	} else if list[2] - list[0] <= 3 {
		ret = computeAdapterArrangements(list[1:]) + computeAdapterArrangements(list[2:])
	} else if list[1]- list[0] <= 3 {
		ret = computeAdapterArrangements(list[1:])
	} else {
		ret = 0
	}
	memo[key] = ret
	return ret
}

func listToString(list []int) string {
	s := ""
	for _, v := range list {
		s = s + strconv.Itoa(v)
	}
	return s
}
