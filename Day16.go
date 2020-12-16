package main

import (
	"fmt"
	"strings"
)

type Day16 struct {
	rules         map[string][]int
	myticket      []int
	nearbytickets [][]int
}

func (day *Day) Day16a() {
	fmt.Printf("Part 1: %d\n", ComputeDay16a(day.input))
}

func (day *Day) Day16b() {
	fmt.Printf("Part 2: %d\n", ComputeDay16b(day.input))
}

func parseDay16(input string) Day16 {
	parts := strings.Split(input, "\n\n")
	rulesStr := strings.Split(parts[0], "\n")
	myticketStr := strings.Split(strings.Split(parts[1], "\n")[1], ",")
	nearbyticketsStr := strings.Split(parts[2], "\n")[1:]

	var data Day16
	data.rules = make(map[string][]int)
	data.myticket = convertStringListToInts(myticketStr)
	data.nearbytickets = make([][]int, len(nearbyticketsStr))
	for i, nearbyticketStr := range nearbyticketsStr {
		data.nearbytickets[i] = convertStringListToInts(strings.Split(nearbyticketStr, ","))
	}
	for _, rulestr := range rulesStr {
		ruleparts := strings.Split(rulestr, ": ")
		rulename := ruleparts[0]
		intervals := strings.Split(ruleparts[1], " or ")
		bounds := append(strings.Split(intervals[0], "-"), strings.Split(intervals[1], "-")...)
		data.rules[rulename] = convertStringListToInts(bounds)
	}
	return data
}

func ComputeDay16a(input string) int {
	data := parseDay16(input)
	errorRate := 0
	for i, _ := range data.nearbytickets {
		errorRate += ticketErrorRate(data, i)
	}
	return errorRate
}

func ComputeDay16b(input string) int {
	data := parseDay16(input)
	validTickets := make([][]int, 0)
	for i, ticket := range data.nearbytickets {
		if ticketErrorRate(data, i) == 0 {
			validTickets = append(validTickets, ticket)
		}
	}

	consistentIndices := make(map[string][]int)
	for rulename, rule := range data.rules {
		consistentIndices[rulename] = getConsistentFieldIndices(validTickets, rule)
	}

	established := make(map[string]int)
	for len(consistentIndices) != 0 {
		for name, indices := range consistentIndices {
			if len(indices) == 1 {
				established[name] = indices[0]
				delete(consistentIndices, name)
				removeFromValidIndices(consistentIndices, indices[0])
				break
			}
			// weird that it's necessary, may have a bug somewhere :^)
			if len(indices) == 0 {
				delete(consistentIndices, name)
			}
		}
	}

	res := 1
	for k, v := range established {
		if strings.HasPrefix(k, "departure") {
			res = res * data.myticket[v]
		}
	}
	return res
}

func removeFromValidIndices(validIndices map[string][]int, i int) {
	for name, indices := range validIndices {
		for k, v := range indices {
			if v == i {
				validIndices[name] = append(indices[:k], indices[k+1:]...)
				break
			}
		}
	}
}

func getConsistentFieldIndices(tickets [][]int, rule []int) []int {
	res := make([]int, 0)
	// all lines have the same number of fields, we pick the first one
	numFields := len(tickets[0])

	for i := 0; i < numFields; i++ {
		consistent := true
		for _, ticket := range tickets {
			field := ticket[i]
			if !((field >= rule[0] && field <= rule[1]) || (field >= rule[2] && field <= rule[3])) {
				consistent = false
				break
			}
		}
		if consistent {
			res = append(res, i)
		}
	}
	return res
}

func ticketErrorRate(data Day16, ticket int) int {
	res := 0
	for _, field := range data.nearbytickets[ticket] {
		matches := false
		for _, bounds := range data.rules {
			if (field >= bounds[0] && field <= bounds[1]) || (field >= bounds[2] && field <= bounds[3]) {
				matches = true
				break
			}
		}
		if !matches {
			res += field
		}
	}
	return res
}
