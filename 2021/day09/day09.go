package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "09"
	dayTitle = "Smoke Basin"
)

func Parse(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0, len(line))
		for _, r := range line {
			row = append(row, util.MustAtoi(string(r)))
		}
		grid = append(grid, row)
	}

	return grid, nil
}

type Point struct{ x, y int }

func expand(grid [][]int, p Point) int {
	var queue []Point
	queue = append(queue, p)
	visited := make(map[Point]struct{})
	size := 0

	for len(queue) != 0 {
		q := queue[0]
		queue = queue[1:]

		if _, ok := visited[q]; ok {
			continue
		}
		if grid[q.y][q.x] == 9 {
			continue
		}

		size++
		visited[q] = struct{}{}

		if q.x > 0 {
			queue = append(queue, Point{q.x - 1, q.y})
		}
		if q.x < len(grid[q.y])-1 {
			queue = append(queue, Point{q.x + 1, q.y})
		}
		if q.y > 0 {
			queue = append(queue, Point{q.x, q.y - 1})
		}
		if q.y < len(grid)-1 {
			queue = append(queue, Point{q.x, q.y + 1})
		}
	}

	return size
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	grid, _ := Parse("input.txt")

	risk := 0
	var lowPoints []Point
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			top, left, right, bot := true, true, true, true
			if x > 0 {
				left = grid[y][x-1] > grid[y][x]
			}
			if x < len(grid[y])-1 {
				right = grid[y][x+1] > grid[y][x]
			}
			if y > 0 {
				top = grid[y-1][x] > grid[y][x]
			}
			if y < len(grid)-1 {
				bot = grid[y+1][x] > grid[y][x]
			}

			if top && bot && left && right {
				risk += grid[y][x] + 1
				lowPoints = append(lowPoints, Point{x, y})
			}
		}
	}
	fmt.Printf("Part 1: %v\n", risk)

	sizes := make([]int, 0, len(lowPoints))
	for _, point := range lowPoints {
		sizes = append(sizes, expand(grid, point))
	}
	sort.Ints(sizes)
	l := len(sizes)
	answer := sizes[l-1] * sizes[l-2] * sizes[l-3]
	fmt.Printf("Part 2: %v\n", answer)
}
