package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "17"
	dayTitle = "Conway Cubes"
)

type Cube struct {
	//x, y, z int //part1
	x, y, z, w int
}

func (c Cube) diff(o Cube) bool {
	dx := abs(c.x - o.x)
	dy := abs(c.y - o.y)
	dz := abs(c.z - o.z)
	dw := abs(c.w - o.w)
	return dx <= 1 && dy <= 1 && dz <= 1 && dw <= 1
}

type Grid map[Cube]struct{}

func parseGrid(lines []string) Grid {
	grid := make(Grid)
	for j, line := range lines {
		for i, r := range line {
			if r == '#' {
				grid[Cube{i, j, 0, 0}] = struct{}{}
			}
		}
	}
	return grid
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getAdj(cube Cube) []Cube {
	var adj []Cube
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}
					neighbor := Cube{cube.x + dx, cube.y + dy, cube.z + dz, cube.w + dw}
					if cube.diff(neighbor) {
						adj = append(adj, neighbor)
					}
				}
			}
		}
	}

	return adj
}

func Cycles(grid Grid, numCycles int) int {
	fmt.Printf("grid: %v\n", len(grid))
	for i := 0; i < numCycles; i++ {
		nextGrid := make(Grid)

		var cubes []Cube
		for cube := range grid {
			cubes = append(cubes, cube)
			cubes = append(cubes, getAdj(cube)...)
		}

		for _, cube := range cubes {
			activeAdjCount := 0
			adjSlice := getAdj(cube)
			//fmt.Printf("num adj: %v\n", len(adjSlice))
			for _, adj := range adjSlice {
				if _, ok := grid[adj]; ok {
					activeAdjCount++
				}
			}

			// apply rules
			if _, ok := grid[cube]; ok { //active
				if activeAdjCount == 2 || activeAdjCount == 3 {
					nextGrid[cube] = struct{}{}
				}
			} else { // inactive
				if activeAdjCount == 3 {
					nextGrid[cube] = struct{}{}
				}
			}

		}
		grid = nextGrid
	}

	return len(grid)
}

func part1() {
	lines, _ := util.ReadInput("input.txt")
	grid := parseGrid(lines)
	fmt.Printf("Part 1: %v\n", Cycles(grid, 6))
}

func part2() {
	fmt.Printf("Part 2: %v\n", 2)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
