package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "12"
	dayTitle = "Hill Climbing Algorithm"
)

var heights []rune = []rune("abcdefghijklmnopqrstuvwxyz")
var hMap map[rune]int = make(map[rune]int)

func makeHMap() map[rune]int {
	for i, r := range heights {
		hMap[r] = i
	}
	hMap['S'] = hMap['a']
	hMap['E'] = hMap['z']
	return hMap
}

func printGrid(grid [][]rune) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Printf("%c", grid[y][x])
		}
		fmt.Println()
	}
}

func adj(grid [][]rune, p util.Point, h int) (n []util.Point) {
	a := []util.Point{
		util.Point{p.X, p.Y - 1},
		util.Point{p.X - 1, p.Y},
		util.Point{p.X, p.Y + 1},
		util.Point{p.X + 1, p.Y},
	}

	for _, b := range a {
		if b.X >= 0 && b.X < len(grid[p.Y]) && b.Y >= 0 && b.Y < len(grid) {
			bh := hMap[grid[b.Y][b.X]]
			if bh <= h+1 {
				n = append(n, b)
			}
		}
	}
	return
}

type Path struct {
	P util.Point
	D int
}

func hike(grid [][]rune, start, end util.Point) int {
	q := []Path{Path{start, 0}}
	visited := make(map[util.Point]struct{})
	for len(q) != 0 {
		curr := q[0]
		q = q[1:]

		if _, ok := visited[curr.P]; ok {
			continue
		}

		visited[curr.P] = struct{}{}
		h := hMap[grid[curr.P.Y][curr.P.X]]

		if curr.P == end {
			return curr.D
		}

		for _, a := range adj(grid, curr.P, h) {
			q = append(q, Path{a, curr.D + 1})
		}
	}

	return 0
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	hMap = makeHMap()
	grid, _ := util.ReadRuneput("input.txt")

	var start, end util.Point
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 'S' {
				start = util.Point{x, y}
			}
			if grid[y][x] == 'E' {
				end = util.Point{x, y}
			}
		}
	}
	min := hike(grid, start, end)
	fmt.Printf("Part 1: %v\n", min)

	var aPoints []util.Point
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 'a' {
				aPoints = append(aPoints, util.Point{x, y})
			}
		}
	}
	for _, p := range aPoints {
		if dist := hike(grid, p, end); dist < min && dist != 0 {
			min = dist
		}
	}
	fmt.Printf("Part 2: %v\n", min)
}
