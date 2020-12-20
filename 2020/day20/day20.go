package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "20"
	dayTitle = "Jurassic Jigsaw"
)

type Tile struct {
	ID   int
	Grid [][]rune
}

// Print the title to standard out
func (t Tile) Print() {
	fmt.Printf("ID: %v\n", t.ID)
	for j := 0; j < len(t.Grid); j++ {
		for i := 0; i < len(t.Grid[j]); i++ {
			fmt.Printf("%c", t.Grid[j][i])
		}
		fmt.Printf("\n")
	}
}

// ParseTile parsing a string slice into tile structs
func ParseTile(lines []string) []Tile {
	var tiles []Tile

	var tile Tile
	for _, line := range lines {
		if strings.HasPrefix(line, "Tile") {
			line = strings.ReplaceAll(line, "Tile ", "")
			line = strings.ReplaceAll(line, ":", "")
			//logDebug("tileID to parse: %v\n", line)
			id, _ := strconv.Atoi(line)
			tile.ID = id
		} else if len(line) == 0 {
			tiles = append(tiles, tile)
			tile = Tile{}
		} else {
			runes := []rune(line)
			tile.Grid = append(tile.Grid, runes)
		}
	}
	return tiles
}

// newGrid create a new 2d rune slice that is square
func newGrid(n int) [][]rune {
	newGrid := make([][]rune, n, n)
	for i := 0; i < n; i++ {
		newGrid[i] = make([]rune, n, n)
	}
	return newGrid
}

// rotate90 For the given tile, a new tile is returned that is rotated 90 degrees counter clockwise
func rotate90(tile Tile) Tile {
	n := len(tile.Grid)
	rot := newGrid(n)
	for j := 0; j < len(tile.Grid); j++ {
		for i := 0; i < len(tile.Grid[j]); i++ {
			rot[i][j] = tile.Grid[n-j-1][i]
		}
	}

	return Tile{tile.ID, rot}
}

// allRotations this will perform 3 rotations on the given title, resulting in 4 new tiles
func allRotations(tile Tile) []Tile {
	var rotations []Tile
	rot := tile
	for i := 0; i < 4; i++ {
		rotations = append(rotations, rot)
		rot = rotate90(rot)
	}
	return rotations
}

// flipHorz For the given tile, a new tile is returned that is flipped horizontally
func flipHorz(tile Tile) Tile {
	n := len(tile.Grid)
	flip := newGrid(n)
	for j := 0; j < len(tile.Grid); j++ {
		for i := 0; i < len(tile.Grid[j]); i++ {
			flip[i][j] = tile.Grid[i][n-j-1]
		}
	}
	return Tile{tile.ID, flip}
}

// flipVert For the given tile, a new tile is returned that is flipped vertically
func flipVert(tile Tile) Tile {
	n := len(tile.Grid)
	flip := newGrid(n)
	for j := 0; j < len(tile.Grid); j++ {
		for i := 0; i < len(tile.Grid[j]); i++ {
			flip[i][j] = tile.Grid[n-1-i][j]
		}
	}
	return Tile{tile.ID, flip}
}

// orientations for the given tile,
func orientations(tile Tile) []Tile {
	var orientations []Tile
	orientations = append(orientations, allRotations(tile)...)
	orientations = append(orientations, allRotations(flipHorz(tile))...)
	orientations = append(orientations, allRotations(flipVert(tile))...)
	// orientations = append(orientations, flipHorz(tile))
	// orientations = append(orientations, flipVert(tile))
	//logDebug("total orientations: %v\n", len(orientations))
	return orientations
}

// hasTileID checks to see if tile's ID exists amongst the tiles in tiles
func hasTileID(tiles []Tile, tile Tile) bool {
	for _, t := range tiles {
		if t.ID == tile.ID {
			return true
		}
	}
	return false
}

func checkTopToBot(botTile Tile, topTile Tile) bool {
	otherRow := len(topTile.Grid) - 1
	tileRow := 0
	for j := 0; j < len(topTile.Grid); j++ {
		o := topTile.Grid[otherRow][j]
		t := botTile.Grid[tileRow][j]
		//logDebug("other: %c tile: %c bool: %v \n", o, t, o == t)
		if o != t {
			return false
		}
	}
	return true
}

func checkLeftToRight(rightTile Tile, leftTile Tile) bool {
	otherCol := len(leftTile.Grid) - 1
	tileCol := 0
	for i := 0; i < len(leftTile.Grid); i++ {
		t := rightTile.Grid[i][tileCol]
		o := leftTile.Grid[i][otherCol]
		//logDebug("other: %c tile: %c bool: %v \n", o, t, o == t)
		if o != t {
			return false
		}
	}
	return true
}

// checkBorder true if title can be added to solution. In that its border matches the tiles before it (left and up)
func checkBorder(sol []Tile, tile Tile, n int) bool {
	l := len(sol)
	if l == 0 {
		return true
	}

	if l%n != 0 {
		//logDebug("comparing left to right: %v to %v\n", l, l-1)
		other := sol[l-1]
		if !checkLeftToRight(other, tile) {
			return false
		}
	}

	if l >= n {
		//logDebug("comparing top to bottom: %v to %v\n", l, l-n)
		other := sol[l-n]
		if !checkTopToBot(tile, other) {
			return false
		}
	}

	return true
}

func search(tiles []Tile, sol []Tile, n int) []Tile {
	//logDebug("solution so far: %v\n", len(sol))
	if len(tiles) == 0 {
		return sol
	}

	for tileIdx, tile := range tiles {
		// need to check to see if the tile already exists in the solution
		/*
			if hasTileID(sol, tile) {
				continue
			}
		*/

		for _, ori := range orientations(tile) {
			// check if ori can be added to solutions
			if !checkBorder(sol, ori, n) {
				continue
			}

			// add it to the solution
			sol = append(sol, ori)
			// remove it from the possible choice of tiles
			var iterTiles []Tile
			iterTiles = append(iterTiles, tiles[:tileIdx]...)
			iterTiles = append(iterTiles, tiles[tileIdx+1:]...)
			// recurse to continue the search, on a smaller problem space
			ans := search(iterTiles, sol, n)
			if ans != nil {
				return ans
				//printSolution(ans, n)
			}
			sol = sol[:len(sol)-1]
		}
	}

	return nil
}

func prodCorners(tiles []Tile, n int) int {
	if len(tiles) == 0 {
		return -1
	}

	corners := []int{0, n - 1, len(tiles) - 1, len(tiles) - n}
	prod := 1
	for _, idx := range corners {
		prod *= tiles[idx].ID
	}
	return prod
}

func printSolution(tiles []Tile, n int) {
	for i, tile := range tiles {
		if i%n == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%v ", tile.ID)
	}
	fmt.Printf("\n")
}

func printTileIds(tiles []Tile) {
	for i, t := range tiles {
		logDebug("%v: %v\n", i, t.ID)
	}
}

// FindCorners will return the 4 corners
func FindCorners(tiles []Tile) []Tile {
	// Tiles at the edge of the image also have this border, but the outermost edges won't line up with any other tiles.
	var corners []Tile
	for _, tile := range tiles {
		count := 0
		for _, rot := range allRotations(tile) {
		tileSearch:
			for _, other := range tiles {
				// no need to compare to yourself
				if tile.ID == other.ID {
					continue
				}

				for _, otherOri := range orientations(other) {
					if checkLeftToRight(rot, otherOri) {
						//logDebug("tile %v matches with other %v\n", tile.ID, other.ID)
						count++
						continue tileSearch
					}
				}
			}
		}
		// we tried all other tiles against this one
		// should equal the number of matching sides
		//logDebug("tile %v count: %v\n", tile.ID, count)
		if count == 2 {
			corners = append(corners, tile)
		}
	}
	return corners
}

func SolveCorners(filename string) int {
	lines, _ := util.ReadInput(filename)
	tiles := ParseTile(lines)
	n := int(math.Sqrt(float64(len(tiles))))

	logDebug("number of tiles: %v %vx%v\n", len(tiles), n, n)
	corners := FindCorners(tiles)

	if l := len(corners); l != 4 {
		logDebug("Incorrect number of corners detected: %v\n", l)
		return -1
	}

	prod := 1
	for _, c := range corners {
		prod *= c.ID
		//logDebug("%v: tile %v prod %v\n", i, c.ID, prod)
	}
	return prod
}

func SolveBorders(filename string) int {
	lines, _ := util.ReadInput(filename)
	tiles := ParseTile(lines)
	n := int(math.Sqrt(float64(len(tiles))))

	logDebug("number of tiles: %v %vx%v\n", len(tiles), n, n)

	sol := make([]Tile, 0, len(tiles))
	sol = search(tiles, sol, n)
	return prodCorners(sol, n)
}

const debug = true

func logDebug(s string, arg ...interface{}) {
	if debug {
		fmt.Printf(s, arg...)
	}
}

func checkTest() {
	fmt.Println("check test")
	testLines, _ := util.ReadInput("check_test.txt")
	testTiles := ParseTile(testLines)
	testN := 3
	testTile := testTiles[5]
	testTiles = testTiles[:5]
	fmt.Printf("should pass: %v \n", checkBorder(testTiles, testTile, testN))
}

func test1() {
	fmt.Println("Test 1")
	fmt.Printf("Corners: %v\n", SolveCorners("test1.txt"))
	fmt.Printf("Borders: %v\n", SolveBorders("test1.txt"))
}

func part1() {
	fmt.Printf("Part 1: %v\n", SolveCorners("input.txt"))
}

func part2() {
	fmt.Printf("Part 2: %v\n", 2)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	//checkTest()
	test1()

	part1()
	part2()
}
