package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (day *Day) Day11a() {
	fmt.Printf("Part 1: %d\n", ComputeDay11a(day.input))
}

func (day *Day) Day11b() {
	fmt.Printf("Part 2: %d\n", ComputeDay11b(day.input))
}

func ComputeDay11a(input string) int {
	return runSeatingSimulation(input, false)
}

func runSeatingSimulation(input string, part2 bool) int {
	grid := strings.Split(input, "\n")
	newGrid := computeNextSeating(grid, part2)
	for isDifferentSeating(grid, newGrid) {
		grid = newGrid
		newGrid = computeNextSeating(grid, part2)
	}
	return countOccupiedSeats(grid)
}

func countOccupiedSeats(grid []string) int {
	res := 0
	for _, line := range grid {
		res += strings.Count(line, "#")
	}
	return res
}

func computeNextSeating(grid []string, part2 bool) []string {
	newGrid := make([]string, len(grid))
	for i, line := range grid {
		newGrid[i] = ""
		for j := 0; j < len(line); j++ {
			newGrid[i] += computeNewValue(grid, i, j, part2)
		}
	}
	return newGrid
}

func computeNewValue(grid []string, i int, j int, part2 bool) string {
	if grid[i][j] == '.' {
		return "."
	}
	if grid[i][j] == 'L' {
		if shouldFillSeat(grid, i, j, part2) {
			return "#"
		} else {
			return "L"
		}
	}
	if grid[i][j] == '#' {
		if shouldEmptySeat(grid, i, j, part2) {
			return "L"
		} else {
			return "#"
		}
	}
	panic("Could not determine state of seat " + strconv.Itoa(i) + " " + strconv.Itoa(j))
}

func shouldFillSeat(grid []string, i int, j int, part2 bool) bool {
	if !part2 {
		return numOccupiedAdjacentSeats(grid, i, j) == 0
	}
	return NumberOccupiedVisibleSeats(grid, i, j) == 0
}

func shouldEmptySeat(grid []string, i int, j int, part2 bool) bool {
	if !part2 {
		return numOccupiedAdjacentSeats(grid, i, j) >= 4
	}
	return NumberOccupiedVisibleSeats(grid, i, j) >= 5
}

func NumberOccupiedVisibleSeats(grid []string, i int, j int) int {
	res := 0
	for x, y := i-1, j-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
		if grid[x][y] == '#' {
			res++
			break
		}
		if grid[x][y] == 'L' {
			break
		}
	}
	for x, y := i+1, j-1; y >= 0 && x < len(grid); x, y = x+1, y-1 {
		if grid[x][y] == '#' {
			res++
			break
		}
		if grid[x][y] == 'L' {
			break
		}
	}
	for x, y := i+1, j+1; x < len(grid) && y < len(grid[x]); x, y = x+1, y+1 {
		if grid[x][y] == '#' {
			res++
			break
		}
		if grid[x][y] == 'L' {
			break
		}
	}
	for x, y := i-1, j+1; x >= 0 && y < len(grid[x]); x, y = x-1, y+1 {
		if grid[x][y] == '#' {
			res++
			break
		}
		if grid[x][y] == 'L' {
			break
		}
	}

	for y := j-1; y>=0; y-- {
		if grid[i][y] == '#' {
			res++
			break
		}
		if grid[i][y] == 'L' {
			break
		}
	}
	for y := j+1; y < len(grid[i]); y++ {
		if grid[i][y] == '#' {
			res++
			break
		}
		if grid[i][y] == 'L' {
			break
		}
	}
	for x := i-1; x >= 0; x-- {
		if grid[x][j] == '#' {
			res++
			break
		}
		if grid[x][j] == 'L' {
			break
		}
	}
	for x := i+1; x < len(grid); x++ {
		if grid[x][j] == '#' {
			res++
			break
		}
		if grid[x][j] == 'L' {
			break
		}
	}
	return res
}

func numOccupiedAdjacentSeats(grid []string, i int, j int) int {
	res := 0
	if i > 0 && j > 0 && grid[i-1][j-1] == '#' {
		res++
	}
	if i > 0 && grid[i-1][j] == '#' {
		res++
	}
	if i > 0 && j < len(grid[i-1])-1 && grid[i-1][j+1] == '#' {
		res++
	}
	if j > 0 && grid[i][j-1] == '#' {
		res++
	}
	if j < len(grid[i])-1 && grid[i][j+1] == '#' {
		res++
	}
	if i < len(grid)-1 && j > 0 && grid[i+1][j-1] == '#' {
		res++
	}
	if i < len(grid)-1 && grid[i+1][j] == '#' {
		res++
	}
	if i < len(grid)-1 && j < len(grid[i+1])-1 && grid[i+1][j+1] == '#' {
		res++
	}
	return res
}

func isDifferentSeating(grid []string, grid2 []string) bool {
	if len(grid) != len(grid2) {
		return true
	}
	for i := 0; i < len(grid); i++ {
		if grid[i] != grid2[i] {
			return true
		}
	}
	return false
}

func ComputeDay11b(input string) int {
	return runSeatingSimulation(input, true)
}
