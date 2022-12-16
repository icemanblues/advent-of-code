package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "14"
	dayTitle = "Regolith Reservoir"
)

func parse(filename string) (map[util.Point]rune, int) {
	grid := make(map[util.Point]rune)
	abyss := 0
	input, _ := util.ReadInput(filename)
	for _, line := range input {
		coords := strings.Split(line, " -> ")
		for i := 0; i < len(coords)-1; i++ {
			xy1 := strings.Split(coords[i], ",")
			x1, _ := strconv.Atoi(xy1[0])
			y1, _ := strconv.Atoi(xy1[1])

			xy2 := strings.Split(coords[i+1], ",")
			x2, _ := strconv.Atoi(xy2[0])
			y2, _ := strconv.Atoi(xy2[1])

			var xDir, yDir int
			if c := x2 - x1; c > 0 {
				xDir = 1
			} else if c < 0 {
				xDir = -1
			}
			if c := y2 - y1; c > 0 {
				yDir = 1
			} else if c < 0 {
				yDir = -1
			}

			// build the grid
			cx, cy := x1, y1
			grid[util.Point{cx, cy}] = '#'
			for cx != x2 || cy != y2 {
				cx += xDir
				cy += yDir
				grid[util.Point{cx, cy}] = '#'
			}

			// compute the abyss
			if y1 > abyss {
				abyss = y1
			}
			if y2 > abyss {
				abyss = y2
			}
		}
	}
	return grid, abyss + 1
}

func part1() {
	grid, abyss := parse("input.txt")
	sandStart := util.Point{500, 0}
	curr := sandStart
	sandCount := 0
outer:
	for curr.Y <= abyss {
		possible := []util.Point{
			util.Point{curr.X, curr.Y + 1},
			util.Point{curr.X - 1, curr.Y + 1},
			util.Point{curr.X + 1, curr.Y + 1},
		}

		for _, pos := range possible {
			if _, ok := grid[pos]; !ok {
				curr = pos
				continue outer
			}
		}

		// can't move, so deposit the sand here
		grid[curr] = 'o'
		curr = sandStart
		sandCount++
	}
	fmt.Printf("Part 1: %v\n", sandCount)
}

func part2() {
	grid, abyss := parse("input.txt")
	sandStart := util.Point{500, 0}
	curr := sandStart
	sandCount := 0
outer:
	for grid[sandStart] != 'o' {
		possible := []util.Point{
			util.Point{curr.X, curr.Y + 1},
			util.Point{curr.X - 1, curr.Y + 1},
			util.Point{curr.X + 1, curr.Y + 1},
		}

		for _, pos := range possible {
			if pos.Y == abyss+1 {
				break // floor
			}
			if _, ok := grid[pos]; !ok {
				curr = pos
				continue outer
			}
		}

		// can't move, so deposit the sand here
		grid[curr] = 'o'
		curr = sandStart
		sandCount++
	}
	fmt.Printf("Part 2: %v\n", sandCount)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
