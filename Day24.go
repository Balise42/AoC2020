package main

import (
	"fmt"
	"strings"
)

func (day *Day) Day24a() {
	fmt.Printf("Part 1: %d\n", ComputeDay24a(day.input))
}

func (day *Day) Day24b() {
	fmt.Printf("Part 2: %d\n", ComputeDay24b(day.input))
}

type Day24Coords struct {
	x int
	y int
}

func ComputeDay24a(input string) int {
	lines := strings.Split(input, "\n")

	grid := initDay24Grid(lines)
	return computeDay24BlackTiles(grid)
}

func computeDay24BlackTiles(grid map[Day24Coords]bool) int {
	numBlackTiles := 0
	for _, v := range grid {
		if v {
			numBlackTiles++
		}
	}
	return numBlackTiles
}

func initDay24Grid(lines []string) map[Day24Coords]bool {
	grid := make(map[Day24Coords]bool)

	for _, line := range lines {
		posx, posy := 0, 0
		for i := 0; i < len(line); i++ {
			if line[i] == 'e' {
				posx = posx + 2
			} else if line[i] == 'w' {
				posx = posx - 2
			} else if line[i] == 's' {
				if line[i+1] == 'e' {
					posx = posx + 1
				} else {
					posx = posx - 1
				}
				posy = posy - 1
				i++
			} else if line[i] == 'n' {
				if line[i+1] == 'e' {
					posx = posx + 1
				} else {
					posx = posx - 1
				}
				posy = posy + 1
				i++
			}
		}
		tile, ok := grid[Day24Coords{posx, posy}]
		if !ok {
			grid[Day24Coords{posx, posy}] = true
		} else {
			grid[Day24Coords{posx, posy}] = !tile
		}
	}

	return grid
}

func ComputeDay24b(input string) int {
	lines := strings.Split(input, "\n")

	grid := initDay24Grid(lines)

	for i := 0; i < 100; i++ {
		grid = runDay24Step(grid)
	}
	return computeDay24BlackTiles(grid)

}

func runDay24Step(grid map[Day24Coords]bool) map[Day24Coords]bool {
	// first we fill in neighbors of black cells so that we're sure to have them when we run through the step
	for k, v := range grid {
		if v {
			neighbors := getNeighbors(k)
			for _, n := range neighbors {
				_, ok := grid[n]
				if !ok {
					grid[n] = false
				}
			}
		}
	}

	newGrid := make(map[Day24Coords]bool)

	for k, v := range grid {
		num := getNumBlackNeighbors(k, grid)
		if v && (num == 0 || num > 2) {
			newGrid[k] = false
		} else if !v && num == 2 {
			newGrid[k] = true
		} else {
			newGrid[k] = grid[k]
		}
	}
	return newGrid
}

func getNumBlackNeighbors(k Day24Coords, grid map[Day24Coords]bool) int {
	neighbors := getNeighbors(k)
	num := 0
	for _, n := range neighbors {
		black, ok := grid[n]
		if ok && black {
			num++
		}
	}
	return num
}

func getNeighbors(k Day24Coords) []Day24Coords {
	return []Day24Coords {
		{k.x - 2, k.y},
		{k.x + 2, k.y},
		{k.x - 1, k.y - 1},
		{k.x - 1, k.y + 1},
		{k.x + 1, k.y - 1},
		{k.x + 1, k.y + 1},
	}
}