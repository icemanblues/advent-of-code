package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "03"
	dayTitle = "Perfectly Spherical Houses in a Vacuum"
)

type Coord struct{ x, y int }

func part1() {
	lines, _ := util.ReadInput("input.txt")
	curr := Coord{0, 0}
	visited := map[Coord]struct{}{
		curr: {},
	}
	for _, line := range lines {

		for _, r := range line {
			switch r {
			case '<':
				curr = Coord{curr.x - 1, curr.y}
			case '>':
				curr = Coord{curr.x + 1, curr.y}
			case '^':
				curr = Coord{curr.x, curr.y + 1}
			case 'v':
				curr = Coord{curr.x, curr.y - 1}
			}
			visited[curr] = struct{}{}
		}
	}
	fmt.Printf("Part 1: %v\n", len(visited))
}

func part2() {
	lines, _ := util.ReadInput("input.txt")
	santa := Coord{0, 0}
	robot := Coord{0, 0}
	visited := map[Coord]struct{}{
		santa: {},
	}
	for _, line := range lines {
		for i, r := range line {
			curr := &santa
			if i%2 == 1 {
				curr = &robot
			}
			switch r {
			case '<':
				curr.x--
			case '>':
				curr.x++
			case '^':
				curr.y++
			case 'v':
				curr.y--
			}
			visited[*curr] = struct{}{}
		}
	}
	fmt.Printf("Part 2: %v\n", len(visited))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
