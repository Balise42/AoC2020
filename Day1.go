package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (day *Day) Day1a() {
	fmt.Printf("Part 1: %d\n", computeDay1a(day.input))
}

func (day *Day) Day1b() {
	fmt.Printf("Part 2: %d\n", computeDay2a(day.input))
}

func computeDay1a(input string) int {
	listAsInts := createListAsInts(input)
	for i, a := range listAsInts {
		for _, b := range listAsInts[i+1:] {
			if a+b == 2020 {
				return a * b
			}
		}
	}
	panic("Could not find result")
}

func computeDay2a(input string) int {
	listAsInts := createListAsInts(input)
	for i, a := range listAsInts {
		for j, b := range listAsInts[i+1:] {
			for _, c := range listAsInts[j+1:] {
				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}
	panic("Could not find result")
}

func createListAsInts(input string) []int {
	intsAsString := strings.Split(input, "\r\n")
	res := make([]int, len(intsAsString))
	var err error
	for i, v := range intsAsString {
		res[i], err = strconv.Atoi(v)
		if err != nil {
			panic("Could not convert input to number")
		}
	}
	return res
}
