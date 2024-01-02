package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/graph"
	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "16"
	dayTitle = "The Floor Will Be Lava"
)

type Laser struct {
	Point     graph.Point2D
	Direction graph.Direction
}

func (l Laser) move() Laser {
	return Laser{graph.Move(l.Point, l.Direction), l.Direction}
}

func split(l Laser, tile rune) []Laser {
	switch tile {
	case '.':
		return []Laser{l}
	case '|':
		switch l.Direction {
		case graph.North:
			return []Laser{l}
		case graph.East:
			return []Laser{{l.Point, graph.North}, {l.Point, graph.South}}
		case graph.South:
			return []Laser{l}
		case graph.West:
			return []Laser{{l.Point, graph.North}, {l.Point, graph.South}}
		}
	case '-':
		switch l.Direction {
		case graph.North:
			return []Laser{{l.Point, graph.East}, {l.Point, graph.West}}
		case graph.East:
			return []Laser{l}
		case graph.South:
			return []Laser{{l.Point, graph.East}, {l.Point, graph.West}}
		case graph.West:
			return []Laser{l}
		}
	case '/':
		switch l.Direction {
		case graph.North:
			return []Laser{{l.Point, graph.East}}
		case graph.East:
			return []Laser{{l.Point, graph.North}}
		case graph.South:
			return []Laser{{l.Point, graph.West}}
		case graph.West:
			return []Laser{{l.Point, graph.South}}
		}
	case '\\':
		switch l.Direction {
		case graph.North:
			return []Laser{{l.Point, graph.West}}
		case graph.East:
			return []Laser{{l.Point, graph.South}}
		case graph.South:
			return []Laser{{l.Point, graph.East}}
		case graph.West:
			return []Laser{{l.Point, graph.North}}
		}
	default:
		fmt.Printf("Unknown spot on grid: %v\n", tile)
		return nil
	}
	fmt.Printf("WTF Unknown spot on grid: %v\n", tile)
	return nil
}

func energized(grid [][]rune, start Laser) int {
	visited := make(map[Laser]struct{})
	energized := make(map[graph.Point2D]struct{})
	lasers := []Laser{start}
	for len(lasers) != 0 {
		laser := lasers[0]
		lasers = lasers[1:]

		if _, ok := visited[laser]; ok {
			continue // laser loop detected
		}

		// check if in bounds
		if laser.Point.X < 0 || laser.Point.Y < 0 || laser.Point.Y >= len(grid) || laser.Point.X >= len(grid[laser.Point.Y]) {
			continue
		}

		energized[laser.Point] = struct{}{}
		visited[laser] = struct{}{}

		// based on the tile, add the next lasers to the queue
		for _, ll := range split(laser, grid[laser.Point.Y][laser.Point.X]) {
			lasers = append(lasers, ll.move())
		}
	}
	return len(energized)
}

func part1() {
	grid, _ := util.ReadRuneput("input.txt")
	start := Laser{graph.NewPoint2D(0, 0), graph.East}
	fmt.Printf("Part 1: %v\n", energized(grid, start))
}

func part2() {
	grid, _ := util.ReadRuneput("input.txt")
	max := 0

	for x := 0; x < len(grid[0]); x++ {
		// top row y=0, south
		if m := energized(grid, Laser{graph.NewPoint2D(x, 0), graph.South}); m > max {
			max = m
		}
		// bot row y=len-1, north
		if m := energized(grid, Laser{graph.NewPoint2D(x, len(grid)-1), graph.North}); m > max {
			max = m
		}
	}

	for y := 0; y < len(grid); y++ {
		// left col x=0, east
		if m := energized(grid, Laser{graph.NewPoint2D(0, y), graph.East}); m > max {
			max = m
		}
		// right col x=len-1, west
		if m := energized(grid, Laser{graph.NewPoint2D(len(grid[y])-1, y), graph.East}); m > max {
			max = m
		}
	}
	fmt.Printf("Part 2: %v\n", max)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
