package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func (day *Day) Day14a() {
	fmt.Printf("Part 1: %d\n", ComputeDay14a(day.input))
}

func (day *Day) Day14b() {
	fmt.Printf("Part 2: %d\n", ComputeDay14b(day.input))
}

func ComputeDay14a(input string) int64 {
	return writeMemory(input, false)
}

func ComputeDay14b(input string) int64 {
	return writeMemory(input, true)
}

func writeMemory(input string, part2 bool) int64 {
	lines := strings.Split(input, "\n")
	memory := make(map[int64]int64)
	mask := strings.Repeat("0", 36)

	var lineParser = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)

	for _, line := range lines {
		if line[0:4] == "mask" {
			mask = line[7:]
		} else {
			toks := lineParser.FindStringSubmatch(line)
			if len(toks) < 3 {
				panic("Could not parse memory line " + line)
			}
			key, err := strconv.ParseInt(toks[1], 10, 64)
			if err != nil {
				panic("Could not parse memory key " + line)
			}
			val, err := strconv.ParseInt(toks[2], 10, 64)
			if err != nil {
				panic("Could not parse memory value " + line)
			}
			if !part2 {
				val = applyMaskPart1(val, mask)
				memory[key] = val
			} else {
				applyMaskPart2(val, mask, &memory, key)
			}
		}
	}

	sum := int64(0)

	for _, v := range memory {
		sum += v
	}

	return sum
}

func applyMaskPart1(val int64, mask string) int64 {
	bitvalue := strconv.FormatInt(val, 2)
	bitvalue = strings.Repeat("0", 36-len(bitvalue)) + bitvalue
	for i, c := range mask {
		if c != 'X' {
			bitvalue = bitvalue[:i] + string(c) + bitvalue[i+1:]
		}
	}
	ret, err := strconv.ParseInt(bitvalue, 2, 64)
	if err != nil {
		panic("Could not parse masked value " + bitvalue)
	}
	return ret
}

func applyMaskPart2(val int64, mask string, memory *map[int64]int64, key int64) {
	bitvalue := strconv.FormatInt(key, 2)
	bitvalue = strings.Repeat("0", 36-len(bitvalue)) + bitvalue

	xindices := make([]int, 0)

	// let's first set to 1 what needs to be set to 1, and gather all the Xs
	for i, c := range mask {
		if c == '1' {
			bitvalue = bitvalue[:i] + string(c) + bitvalue[i+1:]
		} else if c == 'X' {
			xindices = append(xindices, i)
		}
	}

	base, err := strconv.ParseInt(bitvalue, 2, 64)
	if err != nil {
		panic("Could not compute base " + bitvalue)
	}

	for i := 0; float64(i) < math.Pow(2.0, float64(len(xindices))); i++ {
		rawValue := strconv.FormatInt(int64(i), 2)
		if len(rawValue) < len(xindices) {
			rawValue = strings.Repeat("0", len(xindices)-len(rawValue)) + rawValue
		}

		mask := strings.Repeat("X", 36)
		for j, index := range xindices {
			mask = mask[:index] + string(rawValue[j]) + mask[index + 1:]
		}
		(*memory)[applyMaskPart1(base, mask)] = val
	}
}
