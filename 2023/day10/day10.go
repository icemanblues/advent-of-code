package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "10"
	dayTitle = "Pipe Maze"
)

func traverse(grid [][]rune, p util.Point) (a, b util.Point) {
	switch grid[p.Y][p.X] {
	case '|':
		a, b = util.Point{X: p.X, Y: p.Y - 1}, util.Point{X: p.X, Y: p.Y + 1}
	case '-':
		a, b = util.Point{X: p.X - 1, Y: p.Y}, util.Point{X: p.X + 1, Y: p.Y}
	case 'L':
		a, b = util.Point{X: p.X, Y: p.Y - 1}, util.Point{X: p.X + 1, Y: p.Y}
	case 'J':
		a, b = util.Point{X: p.X, Y: p.Y - 1}, util.Point{X: p.X - 1, Y: p.Y}
	case '7':
		a, b = util.Point{X: p.X, Y: p.Y + 1}, util.Point{X: p.X - 1, Y: p.Y}
	case 'F':
		a, b = util.Point{X: p.X, Y: p.Y + 1}, util.Point{X: p.X + 1, Y: p.Y}
	default:
		fmt.Printf("You are not on the pipeline. %v %c\n", p, grid[p.Y][p.X])
		a, b = p, p
	}
	return
}

func part1() {
	grid, _ := util.ReadRuneput("input.txt")

	// find S
	var start util.Point
startSearch:
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 'S' {
				start = util.Point{X: x, Y: y}
				break startSearch
			}
		}
	}

	// convert S to the proper type of pipe // Do I really need to do this?
	adj := []util.Point{{X: start.X, Y: start.Y + 1},
		{X: start.X, Y: start.Y - 1},
		{X: start.X + 1, Y: start.Y},
		{X: start.X - 1, Y: start.Y},
	}
	startAdj := make([]util.Point, 0, 2)
	for _, a := range adj {
		if a.Y >= 0 && a.Y < len(grid) && a.X >= 0 && a.X < len(grid[a.Y]) {
			b, c := traverse(grid, a)
			if b == start || c == start {
				startAdj = append(startAdj, a)
			}
		}
	}

	// move 2 cursors across the pipe until they meet.
	// return the step count of the cursor traversal
	steps := 1
	curr1, curr2 := startAdj[0], startAdj[1]
	visited := map[util.Point]struct{}{start: {}, curr1: {}, curr2: {}}

	for curr1 != curr2 {
		steps++
		curr1A, curr1B := traverse(grid, curr1)
		if _, ok := visited[curr1A]; ok {
			curr1 = curr1B
		} else {
			curr1 = curr1A
		}

		curr2A, curr2B := traverse(grid, curr2)
		if _, ok := visited[curr2A]; ok {
			curr2 = curr2B
		} else {
			curr2 = curr2A
		}

		visited[curr1] = struct{}{}
		visited[curr2] = struct{}{}
	}

	fmt.Printf("Part 1: %v\n", steps)
}

func part2() {
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
