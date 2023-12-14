package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "13"
	dayTitle = "Point of Incidence"
)

type Points map[util.Point]struct{}

type Grid struct {
	Points     Points
	Xmax, Ymax int
}

func parseGrids(filename string) []Grid {
	input, _ := util.ReadInput(filename)
	var grids []Grid
	points := make(Points)
	y, xmax, ymax := 0, 0, 0
	for _, line := range input {
		if line == "" {
			grids = append(grids, Grid{points, xmax, ymax})
			points = make(Points)
			y, xmax, ymax = 0, 0, 0
			continue
		}
		for x, r := range line {
			if r == '#' {
				points[util.NewPoint(x, y)] = struct{}{}
				if x > xmax {
					xmax = x
				}
				if y > ymax {
					ymax = y
				}
			}
		}
		y++
	}
	grids = append(grids, Grid{points, xmax, ymax})
	return grids
}

// return row (y) symmetry point, col (x) symmetry point
func symmetry(grid Grid) (int, int) {
	col, row := -1, -1

vertical:
	for x := 0; x < grid.Xmax; x++ {
		d := util.Min(x+1, grid.Xmax-x)
		for p := range grid.Points {
			// left of line
			if p.X > x-d && p.X <= x {
				delta := x - p.X + 1
				reflect := util.NewPoint(x+delta, p.Y)
				if _, ok := grid.Points[reflect]; !ok {
					continue vertical
				}
			}
			// right of line
			if p.X >= x+1 && p.X <= x+d {
				delta := p.X - x - 1
				reflect := util.NewPoint(x-delta, p.Y)
				if _, ok := grid.Points[reflect]; !ok {
					continue vertical
				}
			}
		}
		col = x
		break
	}

horizontal:
	for y := 0; y < grid.Ymax; y++ {
		d := util.Min(y+1, grid.Ymax-y)
		for p := range grid.Points {
			// above the line
			if p.Y > y-d && p.Y <= y {
				delta := y - p.Y + 1
				reflect := util.NewPoint(p.X, y+delta)
				if _, ok := grid.Points[reflect]; !ok {
					continue horizontal
				}
			}
			// below the line
			if p.Y >= y+1 && p.Y <= y+d {
				delta := p.Y - y - 1
				reflect := util.NewPoint(p.X, y-delta)
				if _, ok := grid.Points[reflect]; !ok {
					continue horizontal
				}
			}
		}
		row = y
		break
	}

	return row + 1, col + 1
}

func part1() int {
	grids := parseGrids("input.txt")
	sum := 0
	for _, grid := range grids {
		row, col := symmetry(grid)
		sum += col + 100*row
	}
	return sum
}

func part2() int {
	grids := parseGrids("test.txt")
	sum := 0
gridScan:
	for _, grid := range grids {
		for y := 0; y <= grid.Ymax; y++ {
			for x := 0; x <= grid.Xmax; x++ {
				g := make(Points)
				for p := range grid.Points {
					g[p] = struct{}{}
				}

				flip := util.NewPoint(x, y)
				if _, ok := g[flip]; ok {
					delete(g, flip)
				} else {
					g[flip] = struct{}{}
				}

				row, col := symmetry(Grid{g, grid.Xmax, grid.Ymax})
				if row != 0 || col != 0 {
					fmt.Printf("flip %v row %v col: %v\n", flip, row, col)
					sum += col + 100*row
					continue gridScan
				}
			}
		}
	}
	return sum
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}
