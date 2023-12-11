package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "11"
	dayTitle = "Cosmic Expansion"
)

func Expand(galaxies map[util.Point]struct{}, expandSize int) map[util.Point]struct{} {
	rows := make(map[int]struct{})
	cols := make(map[int]struct{})
	maxX, maxY := 0, 0
	for p := range galaxies {
		rows[p.X] = struct{}{}
		cols[p.Y] = struct{}{}
		if maxX < p.X {
			maxX = p.X
		}
		if maxY < p.Y {
			maxY = p.Y
		}
	}

	xCount := 0
	xVoids := make([]int, 0, maxX+1)
	for x := 0; x <= maxX; x++ {
		xVoids = append(xVoids, xCount)
		if _, ok := rows[x]; !ok {
			xCount += expandSize - 1
		}
	}
	yCount := 0
	yVoids := make([]int, 0, maxY+1)
	for y := 0; y <= maxY; y++ {
		yVoids = append(yVoids, yCount)
		if _, ok := cols[y]; !ok {
			yCount += expandSize - 1
		}
	}

	expanded := make(map[util.Point]struct{})
	for g := range galaxies {
		dx := xVoids[g.X]
		dy := yVoids[g.Y]
		expanded[util.Point{X: g.X + dx, Y: g.Y + dy}] = struct{}{}
	}
	return expanded
}

func shortPathSum(galaxy map[util.Point]struct{}) int {
	sum := 0
	points := make([]util.Point, 0, len(galaxy))
	for g := range galaxy {
		points = append(points, g)
	}
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			sum += util.ManhattanDist(points[i], points[j])
		}
	}
	return sum
}

func part1() {
	galaxies, _ := util.ReadSparseGrid("input.txt", '#')
	expanded := Expand(galaxies, 2)
	fmt.Printf("Part 1: %v\n", shortPathSum(expanded))
}

func part2() {
	galaxies, _ := util.ReadSparseGrid("input.txt", '#')
	expanded := Expand(galaxies, 1000000)
	fmt.Printf("Part 2: %v\n", shortPathSum(expanded))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
