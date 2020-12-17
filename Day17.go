package main

import (
	"fmt"
	"strings"
)

func (day *Day) Day17a() {
	fmt.Printf("Part 1: %d\n", ComputeDay17a(day.input))
}

func (day *Day) Day17b() {
	fmt.Printf("Part 2: %d\n", ComputeDay17b(day.input))
}

func ComputeDay17a(input string) int {
	return simulatePart1Cycles(input, 6)
}

func simulatePart1Cycles(input string, rounds int) int {
	lines := strings.Split(input, "\n")

	//initialize the cube
	dim := len(lines) + 2*rounds
	cubes := make([][][]bool, dim)
	for x := 0; x < dim; x++ {
		cubes[x] = make([][]bool, dim)
		for y := 0; y < dim; y++ {
			cubes[x][y] = make([]bool, dim)
		}
	}

	// at most we'll expand of "round" cells in all directions, so we offset the input by that value
	offset := rounds
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				cubes[x+offset][y+offset][0+offset] = true
			}
		}
	}

	for i := 0; i < rounds; i++ {
		cubes = simulatePart1Step(cubes)
	}

	count := 0

	for x := 0; x < dim; x++ {
		for y := 0; y < dim; y++ {
			for z := 0; z < dim; z++ {
				if cubes[x][y][z] {
					count++
				}
			}
		}
	}

	return count
}

func simulatePart1Step(cubes [][][]bool) [][][]bool {
	dim := len(cubes)
	newCubes := make([][][]bool, dim)
	for x := 0; x < dim; x++ {
		newCubes[x] = make([][]bool, dim)
		for y := 0; y < dim; y++ {
			newCubes[x][y] = make([]bool, dim)
		}
	}

	for x := 0; x < dim; x++ {
		for y := 0; y < dim; y++ {
			for z := 0; z < dim; z++ {
				activeCells := 0
				// ! we count the current cell in the active cells
				for i := x - 1; i <= x+1; i++ {
					for j := y - 1; j <= y+1; j++ {
						for k := z - 1; k <= z+1; k++ {
							if isValidCoordinates(i, j, k, dim) && cubes[i][j][k] {
								activeCells++
							}
						}
					}
				}
				// if the cell is active, we have one more active cell in the neighborhood
				if cubes[x][y][z] {
					if activeCells == 3 || activeCells == 4 {
						newCubes[x][y][z] = true
					} else {
						newCubes[x][y][z] = false
					}
				} else {
					if activeCells == 3 {
						newCubes[x][y][z] = true
					} else {
						newCubes[x][y][z] = false
					}
				}
			}
		}
	}
	return newCubes
}

func isValidCoordinates(i int, j int, k int, dim int) bool {
	if i < 0 || j < 0 || k < 0 || i >= dim || j >= dim || k >= dim {
		return false
	}
	return true
}

func ComputeDay17b(input string) int {
	return simulatePart2Cycles(input, 6)
}

func simulatePart2Cycles(input string, rounds int) int {
	lines := strings.Split(input, "\n")

	//initialize the cube
	dim := len(lines) + 2*rounds
	cubes := make([][][][]bool, dim)
	for x := 0; x < dim; x++ {
		cubes[x] = make([][][]bool, dim)
		for y := 0; y < dim; y++ {
			cubes[x][y] = make([][]bool, dim)
			for z := 0; z < dim; z++ {
				cubes[x][y][z] = make([]bool, dim)
			}
		}
	}

	// at most we'll expand of "round" cells in all directions, so we offset the input by that value
	offset := rounds
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				cubes[x+offset][y+offset][0+offset][0+offset] = true
			}
		}
	}

	for i := 0; i < rounds; i++ {
		cubes = simulatePart2Step(cubes)
	}

	count := 0

	for x := 0; x < dim; x++ {
		for y := 0; y < dim; y++ {
			for z := 0; z < dim; z++ {
				for w := 0; w < dim; w++ {
					if cubes[x][y][z][w] {
						count++
					}
				}
			}
		}
	}

	return count
}

func printCubes(cubes [][][][]bool) {
	dim := len(cubes)
	for w := 0; w < dim; w++ {
		for z := 0; z < dim; z++ {
			hasData := false
			for y := 0; y < dim; y++ {
				if hasData {
					break
				}
				for x := 0; x < dim; x++ {
					if cubes[x][y][z][w] {
						hasData = true
						break
					}
				}
			}
			if hasData {
				fmt.Printf("z=%d, w=%d\n", z - 6, w - 6)
				for y := 0; y < dim; y++ {
					for x := 0; x < dim; x++ {
						if cubes[x][y][z][w] {
							fmt.Print("#")
						} else {
							fmt.Print(".")
						}
					}
					fmt.Println("")
				}
				fmt.Println("")
			}
		}
	}
}

func simulatePart2Step(cubes [][][][]bool) [][][][]bool {
	dim := len(cubes)
	newCubes := make([][][][]bool, dim)
	for x := 0; x < dim; x++ {
		newCubes[x] = make([][][]bool, dim)
		for y := 0; y < dim; y++ {
			newCubes[x][y] = make([][]bool, dim)
			for z := 0; z < dim; z++ {
				newCubes[x][y][z] = make([]bool, dim)
			}
		}
	}

	for x := 0; x < dim; x++ {
		for y := 0; y < dim; y++ {
			for z := 0; z < dim; z++ {
				for w := 0; w < dim; w++ {
					activeCells := 0
					// ! we count the current cell in the active cells
					for i := x - 1; i <= x+1; i++ {
						for j := y - 1; j <= y+1; j++ {
							for k := z - 1; k <= z+1; k++ {
								for l := w - 1; l <= w+1; l++ {
									if isValidCoordinates4d(i, j, k, l, dim) && cubes[i][j][k][l] {
										activeCells++
									}
								}
							}
						}
					}
					if cubes[x][y][z][w] {
						// if the cell is active, we have one more active cell in the neighborhood
						if activeCells == 3 || activeCells == 4 {
							newCubes[x][y][z][w] = true
						} else {
							newCubes[x][y][z][w] = false
						}
					} else {
						if activeCells == 3 {
							newCubes[x][y][z][w] = true
						} else {
							newCubes[x][y][z][w] = false
						}
					}
				}
			}
		}
	}
	return newCubes
}

func isValidCoordinates4d(i int, j int, k int, l int, dim int) bool {
	if i < 0 || j < 0 || k < 0 || l < 0 || i >= dim || j >= dim || k >= dim || l >= dim {
		return false
	}
	return true
}
