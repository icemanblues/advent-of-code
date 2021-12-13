package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "11"
	dayTitle = "Dumbo Octopus"
)

func Parse(filename string) (grid [][]int) {
	input, _ := util.ReadRuneput(filename)
	for _, runes := range input {
		row := make([]int, 0, len(runes))
		for _, r := range runes {
			row = append(row, util.MustAtoi(string(r)))
		}
		grid = append(grid, row)
	}
	return

}

type Point struct{ x, y int }

func step(grid [][]int) int {
	// increase all by 1
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			grid[y][x]++
		}
	}

	// search for flashes
	var flashing []Point
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			p := Point{x, y}
			if grid[y][x] > 9 {
				flashing = append(flashing, p)
			}
		}
	}

	// visit the ones that are flashing.
	visited := make(map[Point]struct{})
	for len(flashing) != 0 {
		p := flashing[0]
		flashing = flashing[1:]

		if _, ok := visited[p]; ok {
			continue
		}

		visited[p] = struct{}{}

		// find adj
		for yi := -1; yi <= 1; yi++ {
			for xi := -1; xi <= 1; xi++ {
				if xi == 0 && yi == 0 {
					continue
				}
				yy := p.y + yi
				if yy < 0 || yy >= len(grid) {
					continue
				}
				xx := p.x + xi
				if xx < 0 || xx >= len(grid[yy]) {
					continue
				}

				grid[yy][xx]++
				if grid[yy][xx] > 9 {
					flashing = append(flashing, Point{xx, yy})
				}
			}
		}
	}

	// energy reset to 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] > 9 {
				grid[y][x] = 0
			}
		}
	}

	return len(visited)
}

func part1() {
	grid := Parse("input.txt")
	score := 0
	for steps := 0; steps < 100; steps++ {
		score += step(grid)
	}
	fmt.Printf("Part 1: %v\n", score)
}

func part2() {
	grid := Parse("input.txt")
	target := len(grid) * len(grid[0])
	score, steps := 0, 0
	for ; score != target; steps++ {
		score = step(grid)
	}
	fmt.Printf("Part 2: %v\n", steps)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
