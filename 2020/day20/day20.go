package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "20"
	dayTitle = "Jurassic Jigsaw"
)

type Tile struct {
	ID   int
	Grid [][]rune
}

func ParseTile(lines []string) []Tile {
	var tiles []Tile
	var tile Tile
	for _, line := range lines {
		if strings.HasPrefix(line, "Tile") {
			line = strings.ReplaceAll(line, "Tile ", "")
			line = strings.ReplaceAll(line, ":", "")
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

func newGrid(n int) [][]rune {
	newGrid := make([][]rune, n, n)
	for i := 0; i < n; i++ {
		newGrid[i] = make([]rune, n, n)
	}
	return newGrid
}

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

func allRotations(tile Tile) []Tile {
	var rotations []Tile
	rot := tile
	for i := 0; i < 4; i++ {
		rotations = append(rotations, rot)
		rot = rotate90(rot)
	}
	return rotations
}

func flip(tile Tile) Tile {
	n := len(tile.Grid)
	flip := newGrid(n)
	for j := 0; j < len(tile.Grid); j++ {
		for i := 0; i < len(tile.Grid[j]); i++ {
			flip[i][j] = tile.Grid[i][n-j-1]
		}
	}
	return Tile{tile.ID, flip}
}

func orientations(tile Tile) []Tile {
	var orientations []Tile
	orientations = append(orientations, allRotations(tile)...)
	orientations = append(orientations, allRotations(flip(tile))...)
	return orientations
}

func checkTopToBot(botTile Tile, topTile Tile) bool {
	col := len(topTile.Grid) - 1
	for i := 0; i < len(topTile.Grid[col]); i++ {
		if topTile.Grid[col][i] != botTile.Grid[0][i] {
			return false
		}
	}
	return true
}

func checkLeftToRight(rightTile Tile, leftTile Tile) bool {
	row := len(leftTile.Grid) - 1
	for j := 0; j < len(leftTile.Grid); j++ {
		if leftTile.Grid[j][row] != rightTile.Grid[j][0] {
			return false
		}
	}
	return true
}

func checkJigsaw(sol []Tile, tile Tile, n int) bool {
	l := len(sol)
	if l == 0 {
		return true
	}
	if l%n != 0 {
		other := sol[l-1]
		if !checkLeftToRight(tile, other) {
			return false
		}
	}
	if l >= n {
		other := sol[l-n]
		if !checkTopToBot(tile, other) {
			return false
		}
	}
	return true
}

func SolveCorners(filename string) int {
	lines, _ := util.ReadInput(filename)
	tiles := ParseTile(lines)
	corners := FindConnections(tiles)[2]

	prod := 1
	for _, c := range corners {
		prod *= c.ID
	}
	return prod
}

func hasTile(tiles []Tile, tile Tile) bool {
	for _, t := range tiles {
		if t.ID == tile.ID {
			return true
		}
	}
	return false
}

func search(tiles []Tile, sol []Tile, n int, connCountMap map[int][]Tile) []Tile {
	l := len(sol)
	if l == len(tiles) {
		return sol
	}

	// Given where we are in the solution, we can greatly reduce the search space
	var searchTiles []Tile
	if l == 0 || l == n-1 || l == len(tiles)-1 || l == len(tiles)-n { // corners
		searchTiles = connCountMap[2]
	} else if l%n == 0 || l%n == n-1 { // sides
		searchTiles = connCountMap[3]
	} else if l < n || l >= n*(n-1) { //top and bot
		searchTiles = connCountMap[3]
	} else { // must be in the middle
		searchTiles = connCountMap[4]
	}

	for _, tile := range searchTiles {
		// skip if the tile is already in the solution
		if hasTile(sol, tile) {
			continue
		}
		for _, ori := range orientations(tile) {
			if !checkJigsaw(sol, ori, n) {
				continue
			}

			sol = append(sol, ori)
			ans := search(tiles, sol, n, connCountMap)
			if ans != nil {
				return ans
			}
			sol = sol[:len(sol)-1]
		}
	}

	return nil
}

func FindConnections(tiles []Tile) map[int][]Tile {
	connCountsToTiles := make(map[int][]Tile)
	for _, tile := range tiles {
		count := 0
		for _, rot := range allRotations(tile) {
		tileSearch:
			for _, other := range tiles {
				if tile.ID == other.ID { // no need to compare to yourself
					continue
				}

				for _, otherOri := range orientations(other) {
					if checkLeftToRight(rot, otherOri) {
						count++
						continue tileSearch
					}
				}
			}
		}
		connCountsToTiles[count] = append(connCountsToTiles[count], tile)
	}
	return connCountsToTiles
}

func countHash(tile Tile) int {
	count := 0
	for j := 0; j < len(tile.Grid); j++ {
		for i := 0; i < len(tile.Grid[j]); i++ {
			if '#' == tile.Grid[j][i] {
				count++
			}
		}
	}
	return count
}

func removeBorder(tile Tile) Tile {
	newGrid := tile.Grid[1 : len(tile.Grid)-1]
	for j := 0; j < len(newGrid); j++ {
		newGrid[j] = newGrid[j][1 : len(newGrid[j])-1]
	}
	return Tile{tile.ID, newGrid}
}

func stitchImage(tiles []Tile, n int) Tile {
	var newGrid [][]rune
	for k := 0; k < len(tiles); k += n { // k is the start in tiles
		for j := 0; j < len(tiles[k].Grid); j++ {
			var newRow []rune
			for l := 0; l < n; l++ { // l is the offset within i..i+n
				tile := tiles[k+l]
				newRow = append(newRow, tile.Grid[j]...)
			}
			newGrid = append(newGrid, newRow)
		}
	}

	return Tile{20, newGrid}
}

type Coord struct {
	x, y int
}

/*
                  #
#    ##    ##    ###
 #  #  #  #  #  #
*/
var dragon = []Coord{
	{18, 0},
	{0, 1},
	{5, 1},
	{6, 1},
	{11, 1},
	{12, 1},
	{17, 1},
	{18, 1},
	{19, 1},
	{1, 2},
	{4, 2},
	{7, 2},
	{10, 2},
	{13, 2},
	{16, 2},
}

func isMonster(grid [][]rune, x, y int) bool {
	if x+19 >= len(grid[y]) {
		return false
	}
	if y+2 >= len(grid) {
		return false
	}

	for _, c := range dragon {
		if '#' != grid[y+c.y][x+c.x] {
			return false
		}
	}
	return true
}

func SolveMonsters(filename string) int {
	lines, _ := util.ReadInput(filename)
	tiles := ParseTile(lines)

	n := int(math.Sqrt(float64(len(tiles))))
	connCount := FindConnections(tiles)
	sol := make([]Tile, 0, len(tiles))
	sol = search(tiles, sol, n, connCount)
	return countAllDragons(sol, n)
}

func countAllDragons(sol []Tile, n int) int {
	solNoBorder := make([]Tile, 0, len(sol))
	for _, tile := range sol {
		solNoBorder = append(solNoBorder, removeBorder(tile))
	}
	single := stitchImage(solNoBorder, n)

	countDragon := 0
	for _, ori := range orientations(single) {
		for j := 0; j < len(single.Grid); j++ {
			for i := 0; i < len(single.Grid[j]); i++ {
				if isMonster(ori.Grid, i, j) {
					countDragon++
				}
			}
		}
		if countDragon != 0 {
			break
		}
	}

	dragonHash := countDragon * len(dragon)
	hashCount := countHash(single)
	return hashCount - dragonHash
}

func part1() {
	fmt.Printf("Part 1: %v\n", SolveCorners("input.txt"))
}

func part2() {
	fmt.Printf("Part 2: %v\n", SolveMonsters("input.txt"))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
