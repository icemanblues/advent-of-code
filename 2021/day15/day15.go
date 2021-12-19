package main

import (
	"fmt"
	"sort"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "15"
	dayTitle = "Chiton"
)

func Parse(filename string) [][]int {
	grid, _ := util.ReadIntGrid(filename, "")
	return grid
}

type Path struct {
	Point util.Point
	Risk  int
}

type RiskFunc func(util.Point) int

func adj(p util.Point, maxX, maxY int) []util.Point {
	points := make([]util.Point, 0, 2)
	if p.X != maxX {
		points = append(points, util.Point{X: p.X + 1, Y: p.Y})
	}
	if p.Y != maxY {
		points = append(points, util.Point{X: p.X, Y: p.Y + 1})
	}
	if p.X != 0 {
		points = append(points, util.Point{X: p.X - 1, Y: p.Y})
	}
	if p.Y != 0 {
		points = append(points, util.Point{X: p.X, Y: p.Y - 1})
	}
	return points
}

func dijkstra(start, end util.Point, riskFunc RiskFunc) int {
	shortest := make(map[util.Point]int)
	queue := []Path{{start, 0}}
	for len(queue) != 0 {
		path := queue[0]
		queue = queue[1:]

		nexts := adj(path.Point, end.X, end.Y)
		add := false
		for _, n := range nexts {
			np := Path{n, riskFunc(n) + path.Risk}
			if risk := shortest[np.Point]; risk == 0 || risk > np.Risk {
				shortest[np.Point] = np.Risk
				queue = append(queue, np)
				add = true
			}
		}
		if add {
			sort.Slice(queue, func(i, j int) bool {
				return queue[i].Risk < queue[j].Risk
			})
		}
	}

	return shortest[end]
}

func part1() {
	grid := Parse("input.txt")
	start := util.Point{X: 0, Y: 0}
	end := util.Point{X: len(grid[0]) - 1, Y: len(grid) - 1}
	var riskFunc RiskFunc = func(p util.Point) int {
		return grid[p.Y][p.X]
	}
	fmt.Printf("Part 1: %v\n", dijkstra(start, end, riskFunc))
}

func part2() {
	grid := Parse("input.txt")
	start := util.Point{X: 0, Y: 0}
	end := util.Point{X: 5*len(grid[0]) - 1, Y: 5*len(grid) - 1}
	var riskFunc RiskFunc = func(p util.Point) int {
		lenY := len(grid)
		lenX := len(grid[0])

		x, xr := p.X/lenX, p.X%lenX
		y, yr := p.Y/lenY, p.Y%lenY
		risk := grid[yr][xr]
		for i := 0; i < x+y; i++ {
			risk++
			if risk == 10 {
				risk = 1
			}
		}
		return risk
	}
	fmt.Printf("Part 2: %v\n", dijkstra(start, end, riskFunc))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
