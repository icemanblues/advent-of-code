package util

import "fmt"

type Point struct{ X, Y int }

func PrintGrid(grid [][]int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Print(grid[y][x])
		}
		fmt.Println()
	}
}

func ManhattanDist(a, b Point) int {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y)
}

func Manhattan(a Point) int {
	return a.X + a.Y
}
