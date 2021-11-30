package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "03"
	dayTitle = "Toboggan Trajectory"
)

type route struct {
	r, d int
}

func treeCount(grid [][]rune, r route) int {
	x, y := 0, 0

	treeCount := 0
	for y < len(grid) {
		wrap := x % len(grid[y])
		spot := grid[y][wrap]
		if spot == '#' {
			treeCount++
		}

		x += r.r
		y += r.d
	}
	return treeCount
}

func part1() {
	grid, _ := util.ReadRuneput("input.txt")
	r := route{3, 1}
	fmt.Printf("Part 1: %v\n", treeCount(grid, r))
}

func part2() {
	grid, _ := util.ReadRuneput("input.txt")
	routes := []route{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	product := 1
	for _, r := range routes {
		tc := treeCount(grid, r)
		product *= tc
	}
	fmt.Printf("Part 2: %v\n", product)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
