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
	x, y, z, w int
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

func getAdj3D(cube Cube) []Cube {
	var adj []Cube
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if dx == 0 && dy == 0 && dz == 0 {
					continue
				}
				neighbor := Cube{cube.x + dx, cube.y + dy, cube.z + dz, 0}
				adj = append(adj, neighbor)
			}
		}
	}

	return adj
}

func getAdj4D(cube Cube) []Cube {
	var adj []Cube
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}
					neighbor := Cube{cube.x + dx, cube.y + dy, cube.z + dz, cube.w + dw}
					adj = append(adj, neighbor)
				}
			}
		}
	}

	return adj
}

func Cycles(grid Grid, numCycles int, getAdj func(Cube) []Cube) int {
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
	fmt.Printf("Part 1: %v\n", Cycles(grid, 6, getAdj3D))
}

func part2() {
	lines, _ := util.ReadInput("input.txt")
	grid := parseGrid(lines)
	fmt.Printf("Part 2: %v\n", Cycles(grid, 6, getAdj4D))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
