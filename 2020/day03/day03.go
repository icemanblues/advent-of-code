package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "03"
	dayTitle = "Title"
)

func part1() {
	fmt.Println("Part 1")
	grid, _ := util.ReadInput("input.txt")
	x, y := 0, 0
	r, d := 3, 1

	treeCount := 0
	for y < len(grid) {
		wrap := x % len(grid[y])
		spot := []rune(grid[y])[wrap]
		if spot == '#' {
			treeCount++
		}

		x += r
		y += d
	}
	fmt.Println(treeCount)
}

type route struct {
	r, d int
}

func treeCount(grid []string, r route) int {
	x, y := 0, 0

	treeCount := 0
	for y < len(grid) {
		wrap := x % len(grid[y])
		spot := []rune(grid[y])[wrap]
		if spot == '#' {
			treeCount++
		}

		x += r.r
		y += r.d
	}
	return treeCount
}

func part2() {
	fmt.Println("Part 2")
	grid, _ := util.ReadInput("input.txt")
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
	fmt.Println(product)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
