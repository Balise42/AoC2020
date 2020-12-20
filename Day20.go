package main

import (
	"fmt"
	"regexp"
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
	id                   int
	possibleOrientations [][]int
	orientation          []int
	possibleSides        map[int]bool
	content []string
	orientedContent[]string
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
		sides := getSides(tileLines[1:])

		tile.possibleOrientations = make([][]int, 8, 8)
		tile.possibleOrientations[0] = createDay20TileOrientation(sides, []int{0, 1, 2, 3}, []bool{false, false, false, false})
		tile.possibleOrientations[1] = createDay20TileOrientation(sides, []int{3, 0, 1, 2}, []bool{true, false, true, false})
		tile.possibleOrientations[2] = createDay20TileOrientation(sides, []int{2, 3, 0, 1}, []bool{true, true, true, true})
		tile.possibleOrientations[3] = createDay20TileOrientation(sides, []int{1, 2, 3, 0}, []bool{false, true, false, true})
		tile.possibleOrientations[4] = createDay20TileOrientation(sides, []int{0, 3, 2, 1}, []bool{true, false, true, false})
		tile.possibleOrientations[5] = createDay20TileOrientation(sides, []int{1, 0, 3, 2}, []bool{true, true, true, true})
		tile.possibleOrientations[6] = createDay20TileOrientation(sides, []int{2, 1, 0, 3}, []bool{false, true, false, true})
		tile.possibleOrientations[7] = createDay20TileOrientation(sides, []int{3, 2, 1, 0}, []bool{false, false, false, false})

		allSides := createDay20TileOrientation(sides, []int{0, 1, 2, 3}, []bool{true, true, true, true})
		allSides = append(allSides, createDay20TileOrientation(sides, []int{0, 1, 2, 3}, []bool{false, false, false, false})...)
		tile.possibleSides = make(map[int]bool, 8)
		for _, side := range allSides {
			tile.possibleSides[side] = true
		}
		tile.content = tileLines[1:]
		tiles[id] = &tile
	}

	if !checkTileAssertions(tiles) {
		panic("Tile border unicity assertion is NOT true!")
	} else {
		fmt.Println("Tile border unicity assertion is correct.")
	}
	return tiles
}

func getSides(tileLines []string) []string {
	sides := make([]string, 4)

	// we get the sides of the tiles as "up / right / down / left"
	sides[0] = tileLines[0]
	sides[2] = tileLines[len(tileLines)-1]
	sides[1] = ""
	sides[3] = ""
	for _, line := range tileLines {
		sides[1] += string(line[len(line)-1])
		sides[3] += string(line[0])
	}
	for i, _ := range sides {
		sides[i] = strings.ReplaceAll(sides[i], ".", "0")
		sides[i] = strings.ReplaceAll(sides[i], "#", "1")
	}
	return sides
}


func createDay20TileOrientation(sides []string, isides []int, reverse []bool) []int {
	res := make([]int, 4)
	for i := 0; i < 4; i++ {
		toConvert := sides[isides[i]]
		if reverse[i] {
			toConvert = reverseString(toConvert)
		}

		longres, err := strconv.ParseInt(toConvert, 2, 32)
		if err != nil {
			panic("Could not convert side [" + toConvert + "] to number")
		}
		res[i] = int(longres)
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

func checkTileAssertions(tiles map[int]*Day20Tile) bool {
	allSides := make(map[int]bool, 0)
	for _, tile := range tiles {
		for k := range tile.possibleSides {
			allSides[k] = true
		}
	}
	for k := range allSides {
		num := 0
		for _, tile := range tiles {
			_, ok := tile.possibleSides[k]
			if ok {
				num++
			}
		}
		if num > 2 {
			return false
		}
	}
	return true
}

func getCorners(tiles *map[int]*Day20Tile) []*Day20Tile {
	corners := make([]*Day20Tile, 0)
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
			_, ok := neigh.possibleSides[side]
			if ok {
				num++
				break
			}
		}
	}
	return num
}

func ComputeDay20b(input string) int {
	tilesStr := strings.Split(input, "\n\n")
	tiles := parseDay20Tiles(tilesStr)

	jigsaw := buildJigsaw(&tiles)
	orientTiles(&jigsaw)

	tileSize := len(jigsaw[0][0].orientedContent)

	imageChar := make([][]uint8, len(jigsaw) * tileSize)
	for x := 0; x < len(imageChar); x++ {
		imageChar[x] = make([]uint8, len(jigsaw) * tileSize)
		for y := 0; y < len(imageChar[x]); y++ {
			tilex := x / tileSize
			tiley := y / tileSize
			imageChar[x][y] = jigsaw[tiley][tilex].orientedContent[x % tileSize][y % tileSize]
		}
	}

	image := make([]string, 0)
	for _, line := range imageChar {
		image = append(image, string(line))
	}

	sneks := 0
	for i := 0; i < 4; i++ {
		sneks = countSneks(image)
		if sneks != 0 {
			break
		}
		image = rotateContent(image)
	}
	if sneks == 0 {
		image = flipContent(image)
		for i := 0; i < 4; i++ {
			sneks = countSneks(image)
			if sneks != 0 {
				break
			}
			image = rotateContent(image)
		}
	}

	if sneks == 0 {
		panic("Could not find sneks :(")
	}

	roughness := 0
	for _, line := range image {
		for _, c := range line {
			if c == '#' {
				roughness++
			}
		}
	}

	return roughness - sneks * 15
}

func countSneks(image []string) int {
	numSneks := 0
	bodyRegex := regexp.MustCompile(`#....##....##....###`)
	for i, line := range image[:len(image) - 1] {
		if i == 0 {
			continue
		}
		matches := bodyRegex.FindAllStringIndex(line, -1)
		if matches == nil {
			continue
		}
		for _, match := range matches {
			if image[i-1][match[1] - 2] != '#' {
				continue
			}
			foundSnek := true
			for x := match[0] + 1; x < match[1] - 1; x += 3 {
				if image[i+1][x] != '#' {
					foundSnek = false
					break
				}
			}
			if foundSnek {
				numSneks++
			}
		}
	}
	return numSneks
}

func buildJigsaw(tiles *map[int]*Day20Tile) [][]*Day20Tile {
	corners := getCorners(tiles)
	if len(corners) != 4 {
		panic("Didn't find the right amount of corners!")
	}

	firstCorner := corners[0]
	for _, orientation := range firstCorner.possibleOrientations {
		if numTilesWithEdge(tiles, orientation[0]) == 1 && numTilesWithEdge(tiles, orientation[3]) == 1 {
			firstCorner.orientation = orientation
			break
		}
	}


	jigsawSize := 0
	for {
		if jigsawSize*jigsawSize == len(*tiles) {
			break
		}
		jigsawSize++
	}

	jigsaw := make([][]*Day20Tile, jigsawSize)
	for i := range jigsaw {
		jigsaw[i] = make([]*Day20Tile, jigsawSize)
	}

	freeTiles := make([]*Day20Tile, 0, len(*tiles))
	for _, tile := range *tiles {
		if tile.id != firstCorner.id {
			freeTiles = append(freeTiles, tile)
		}
	}

	jigsaw[0][0] = firstCorner

	addNextPieces(&jigsaw, tiles, freeTiles)
	return jigsaw
}

func addNextPieces(jigsaw *[][]*Day20Tile, tiles *map[int]*Day20Tile, freeTiles []*Day20Tile) {
	jigsawSize := len(*jigsaw)

	for x := 0; x < jigsawSize; x++ {
		for y := 0; y < jigsawSize; y++ {
			if x == 0 && y == 0 {
				continue
			}

			topConst := 0
			rightConst := 0
			if x > 0 {
				rightConst = (*jigsaw)[x-1][y].orientation[1]
			}
			if y > 0 {
				topConst = (*jigsaw)[x][y-1].orientation[2]
			}

			var index int
			(*jigsaw)[x][y], index = getOrientedTile(topConst, rightConst, freeTiles, tiles)
			freeTiles = append(freeTiles[:index], freeTiles[index+1:]...)

		}
	}
}

func getOrientedTile(topConst int, rightConst int, tiles []*Day20Tile, allTiles *map[int]*Day20Tile) (*Day20Tile, int) {
	if topConst != 0 && rightConst != 0 {
		for i, tile := range tiles {
			_, ok := tile.possibleSides[topConst]
			if !ok {
				continue
			}
			_, ok = tile.possibleSides[rightConst]
			if !ok {
				continue
			}
			for _, orientation := range tile.possibleOrientations {
				if orientation[0] == topConst && orientation[3] == rightConst {
					tile.orientation = orientation
				}
			}
			if tile.orientation == nil {
				panic("Could not find orientation!")
			}
			return tile, i
		}
		panic("Could not find tile!")
	} else {
		if topConst == 0 {
			for i, tile := range tiles {
				_, ok := tile.possibleSides[rightConst]
				if !ok {
					continue
				}
				for _, orientation := range tile.possibleOrientations {
					if orientation[3] == rightConst && numTilesWithEdge(allTiles, orientation[0]) == 1 {
						tile.orientation = orientation
					}
				}
				if tile.orientation == nil {
					panic("Could not find orientation!")
				}
				return tile, i
			}
			panic("Could not find tile!")
		} else {
			for i, tile := range tiles {
				_, ok := tile.possibleSides[topConst]
				if !ok {
					continue
				}
				for _, orientation := range tile.possibleOrientations {
					if orientation[0] == topConst && numTilesWithEdge(allTiles, orientation[3]) == 1 {
						tile.orientation = orientation
					}
				}
				if tile.orientation == nil {
					panic("Could not find orientation!")
				}
				return tile, i
			}
			panic("Could not find tile!")
		}
	}
}

func orientTiles(tiles *[][]*Day20Tile) {
	for x, col := range *tiles {
		for y, tile := range col {
			content := tile.content
			orientation := tile.orientation
			for i := 0; i < 4; i++ {
				if matchesOrientation(content, orientation) {
					(*tiles)[x][y].orientedContent = stripContent(content)
					break
				}
				content = rotateContent(content)
			}
			if (*tiles)[x][y].orientedContent != nil {
				continue
			}
			content = flipContent(content)
			for i := 0; i < 4; i++ {
				if matchesOrientation(content, orientation) {
					(*tiles)[x][y].orientedContent = stripContent(content)
					break
				}
				content = rotateContent(content)
			}
			if (*tiles)[x][y].orientedContent == nil {
				panic("Could not orient tile!")
			}
		}
	}
}

func stripContent(content []string) []string {
	res := make([]string, 0)
	for _, line := range content[1:len(content) - 1] {
		res = append(res, line[1:len(line) - 1])
	}
	return res
}

func flipContent(content []string) []string {
	res := make([]string, 0)
	for _, line := range content {
		res = append(res, reverseString(line))
	}
	return res
}

func rotateContent(content []string) []string {
	reschar := make([][]uint8, len(content))
	for x, line := range content {
		reschar[x] = make([]uint8, len(content))
		for y := 0; y < len(line); y++ {
			reschar[x][y] = content[y][x]
		}
	}
	res := make([]string, 0)
	for _, line := range reschar {
		res = append(res, reverseString(string(line)))
	}
	return res
}

func matchesOrientation(content []string, orientation []int) bool {
	sides := getSides(content)
	for i, side := range sides {
		longint, err := strconv.ParseInt(side, 2, 64)
		if err != nil {
			panic("Could not parse to int!")
		}
		if int(longint) != orientation[i] {
			return false
		}
	}
	return true
}

func numTilesWithEdge(tiles *map[int]*Day20Tile, edge int) int {
	num := 0
	for _, tile := range *tiles {
		_, ok := tile.possibleSides[edge]
		if ok {
			num++
		}
	}
	return num
}
