package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (day *Day) Day20a() {
	fmt.Printf("Part 1: %d\n", ComputeDay20a(day.input))
}

func (day *Day) Day20b() {
	fmt.Printf("Part 2: %d\n", ComputeDay20b(day.input))
}

type Day20Tile struct {
	id int
	possibleOrientations [][]int64
	orientation []int64
	possibleSides map[int64]bool
	up *Day20Tile
	right *Day20Tile
	bottom *Day20Tile
	left *Day20Tile
}

func ComputeDay20a(input string) int64 {
	tilesStr := strings.Split(input, "\n\n")
	tiles := parseDay20Tiles(tilesStr)
	corners := getCorners(&tiles)
	if len(corners) != 4 {
		panic("Didn't find the right amount of corners!")
	}

	res := int64(1)
	for _, corner := range corners {
		res *= int64(corner.id)
	}
	return res
}

func getCorners(tiles *map[int]*Day20Tile) []*Day20Tile{
	corners := make([]*Day20Tile,0)
	for id, tile := range *tiles {
		numMatchingEdges := findNumMatchingEdges(tiles, id)
		if numMatchingEdges == 2 {
			corners = append(corners, tile)
		}
	}
	return corners
}

func findNumMatchingEdges(tiles *map[int]*Day20Tile, id int) int {
	tile := (*tiles)[id]
	// we can pick an arbitrary orientation: we're interested in existence, not complete building (yet)
	orientation := tile.possibleOrientations[0]
	num := 0
	for _, side := range orientation {
		for neighId, neigh := range *tiles {
			if neighId == id {
				continue
			}
			_, ok := neigh.possibleSides[side]; if ok {
				num++
				break
			}
		}
	}
	return num
}

func parseDay20Tiles(str []string) map[int]*Day20Tile {
	tiles := make(map[int]*Day20Tile, len(str))

	for _, tilestr := range str {
		tile := Day20Tile{}
		tileLines := strings.Split(tilestr, "\n")
		idStr := tileLines[0][5:9]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic("Could not parse id [" + idStr + "]")
		}

		tile.id = id

		sides := make([]string, 4)

		// we get the sides of the tiles as "up / right / down / left"
		sides[0] = tileLines[1]
		sides[2] = tileLines[len(tileLines) - 1]
		sides[1] = ""
		sides[3] = ""
		for _, line := range tileLines[1:] {
			sides[1] += string(line[len(line) - 1])
			sides[3] += string(line[0])
		}
		for i, _ := range sides {
			sides[i] = strings.ReplaceAll(sides[i], ".", "0")
			sides[i] = strings.ReplaceAll(sides[i], "#", "1")
		}
		tile.possibleOrientations = make([][]int64, 12, 12)
		tile.possibleOrientations[0] = createDay20TileOrientation(sides, []int{0, 1, 2, 3}, []bool {false, false, false, false})
		tile.possibleOrientations[1] = createDay20TileOrientation(sides, []int{3, 0, 1, 2}, []bool {true, false, true, false})
		tile.possibleOrientations[2] = createDay20TileOrientation(sides, []int{2, 3, 0, 1}, []bool{true, true, true, true})
		tile.possibleOrientations[3] = createDay20TileOrientation(sides, []int{1, 2, 3, 0}, []bool{false, true, false, true})
		tile.possibleOrientations[4] = createDay20TileOrientation(sides, []int{0, 3, 2, 1}, []bool{true, false, true, false})
		tile.possibleOrientations[5] = createDay20TileOrientation(sides, []int{1, 0, 3, 2}, []bool{true, true, true, true})
		tile.possibleOrientations[6] = createDay20TileOrientation(sides, []int{2, 1, 0, 3}, []bool{false, true, false, true})
		tile.possibleOrientations[7] = createDay20TileOrientation(sides, []int{3, 2, 1, 0}, []bool{false, false, false, false})
		tile.possibleOrientations[8] = createDay20TileOrientation(sides, []int{2, 3, 0, 1}, []bool{false, true, false, true})
		tile.possibleOrientations[9] = createDay20TileOrientation(sides, []int{1, 2, 3, 0}, []bool{false, false, false, false})
		tile.possibleOrientations[10] = createDay20TileOrientation(sides, []int{0, 1, 2, 3}, []bool{true, false, true, false})
		tile.possibleOrientations[11] = createDay20TileOrientation(sides, []int{3, 0, 1, 2}, []bool{true, true, true, true})

		allSides := createDay20TileOrientation(sides, []int{0,1,2,3}, []bool{true, true, true, true})
		allSides = append(allSides, createDay20TileOrientation(sides, []int{0,1,2,3}, []bool{false, false, false, false})...)
		tile.possibleSides = make(map[int64]bool, 8)
		for _, side := range allSides {
			tile.possibleSides[side] = true
		}
		tiles[id] = &tile
	}

	return tiles
}

func createDay20TileOrientation(sides []string, isides []int, reverse []bool) []int64 {
	res := make([]int64, 4)
	for i := 0; i<4; i++ {
		toConvert := sides[isides[i]]
		if reverse[i] {
			toConvert = reverseString(toConvert)
		}
		var err error
		res[i], err = strconv.ParseInt(toConvert, 2, 32)
		if err != nil {
			panic("Could not convert side [" + toConvert + "] to number")
		}
	}
	return res
}

func reverseString(str string) string {
	res := ""
	for i := len(str) - 1; i >= 0; i-- {
		res = res + string(str[i])
	}
	return res
}

func ComputeDay20b(input string) int64 {
	return 0
}
