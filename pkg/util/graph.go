package util

import "fmt"

type Point struct{ X, Y int }

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

type Point3D struct{ X, Y, Z int }

func NewPoint3D(x, y, z int) Point3D {
	return Point3D{X: x, Y: y, Z: z}
}

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

func ManhattanDist3D(a, b Point3D) int {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y) + Abs(a.Z-b.Z)
}

func Manhattan(a Point) int {
	return a.X + a.Y
}

func Manhattan3D(a Point3D) int {
	return a.X + a.Y + a.Z
}
