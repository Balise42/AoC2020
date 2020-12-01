package main

import (
	"strconv"
	"strings"
)

func (day *Day) Day1a() {
	println(computeDay1a(day.input))
}

func computeDay1a(input string) int {
	listAsInts := createListAsInts(input)
	for i, a := range listAsInts {
		for _, b := range listAsInts[i+1:] {
			if a + b == 2020 {
				return a * b
			}
		}
	}
	panic ("Could not find result")
}

func (day *Day) Day1b() {
	println(computeDay2a(day.input))
}

func computeDay2a(input string) int {
	listAsInts := createListAsInts(input)
	for i, a := range listAsInts {
		for j, b := range listAsInts[i+1:] {
			for _, c := range listAsInts[j+1:] {
				if a + b + c == 2020 {
					return a * b * c
				}
			}
		}
	}
	panic ("Could not find result")
}

func createListAsInts(input string) []int {
	intsAsString := strings.Split(input, "\r\n")
	res := make([]int, len(intsAsString))
	for i, v := range intsAsString {
		conv, err := strconv.Atoi(v)
		if err != nil {
			panic("Could not convert input to number")
		}
		res[i] = conv
	}
	return res
}