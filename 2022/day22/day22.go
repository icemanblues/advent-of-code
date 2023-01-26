package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "22"
	dayTitle = "Monkey Map"
)

type Grid map[util.Point]rune

type Direction rune

const (
	U Direction = 'U'
	R Direction = 'R'
	D Direction = 'D'
	L Direction = 'L'
)

var Movement map[Direction]util.Point = map[Direction]util.Point{
	U: util.NewPoint(0, -1),
	R: util.NewPoint(1, 0),
	D: util.NewPoint(0, 1),
	L: util.NewPoint(-1, 0),
}

func Turn(d Direction, t rune) Direction {
	switch t {
	case 'R':
		switch d {
		case U:
			return R
		case R:
			return D
		case D:
			return L
		case L:
			return U
		default:
			fmt.Printf("Unknown direction for Turn: %c\n", d)
			return d
		}
	case 'L':
		switch d {
		case U:
			return L
		case R:
			return U
		case D:
			return R
		case L:
			return D
		default:
			fmt.Printf("Unknown direction for Turn: %c\n", d)
			return d
		}
	default:
		fmt.Printf("Unknown turn type for Turn: %c\n", t)
		return d
	}
}

func Move(p util.Point, d Direction) util.Point {
	offset, ok := Movement[d]
	if !ok {
		fmt.Printf("Unknown direction for Move: %c\n", d)
		return p
	}
	return util.NewPoint(p.X+offset.X, p.Y+offset.Y)
}

func Parse(filename string) int {
	parseMap := true
	grid := make(Grid)
	input, _ := util.ReadInput("input.txt")
	for y, line := range input {
		if line == "" {
			parseMap = false
			continue
		}

		if parseMap {
			// add it to the grid
			for x, r := range line {
				if r == '.' || r == '#' {
					p := util.NewPoint(x, y)
					grid[p] = r
				}
				// TODO: probably want the max X and max Y for boundary limits
			}
		} else {
			// parse the movement instructions
		}
	}

	return len(input)
}

func part1() {
	n := Parse("input.txt")
	fmt.Printf("Part 1: %v\n", n)
}

func part2() {}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
