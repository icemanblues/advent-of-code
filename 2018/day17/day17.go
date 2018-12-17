package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func safeParse(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("Unable to parse int %v\n", s)
	}
	return x
}

func readInput(filename string) ([]Point, int, int, int, int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return nil, -1, -1, -1, -1
	}
	defer file.Close()

	var clay []Point
	xmin, xmax := 9999999999999, -1
	ymin, ymax := 9999999999999, -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// points that are clay
		// x=495, y=2..7
		// y=7, x=495..501
		line := scanner.Text()
		words := strings.Split(line, ", ")

		left := strings.Split(words[0], "=")
		right := strings.Split(words[1], "=")
		se := strings.Split(right[1], "..")
		start := safeParse(se[0])
		end := safeParse(se[1])

		if left[0] == "x" {
			x := safeParse(left[1])
			if x < xmin {
				xmin = x
			}
			if x > xmax {
				xmax = x
			}

			for y := start; y <= end; y++ {
				clay = append(clay, Point{x, y})

				if y < ymin {
					ymin = y
				}
				if y > ymax {
					ymax = y
				}
			}

		} else {
			y := safeParse(left[1])
			if y < ymin {
				ymin = y
			}
			if y > ymax {
				ymax = y
			}

			for x := start; x <= end; x++ {
				if x < xmin {
					xmin = x
				}
				if x > xmax {
					xmax = x
				}

				clay = append(clay, Point{x, y})
			}
		}

	}
	return clay, xmin, xmax, ymin, ymax
}

func main() {
	fmt.Println("Day 17: Reservoir Research")

	part1("test1.txt")
	part1("input17.txt")
	// part2()
}

var Spring Point = Point{X: 500, Y: 0}

func printState(clay, water map[Point]rune, xmin, xmax, ymin, ymax int) {
	for y := ymin; y <= ymax; y++ {
		for x := xmin; x <= xmax; x++ {
			p := Point{x, y}
			if r, ok := clay[p]; ok {
				fmt.Printf("%c", r)
				continue
			}
			if r, ok := water[p]; ok {
				fmt.Printf("%c", r)
				continue
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}

}

func pour(clay, water map[Point]rune, xmin, xmax, ymin, ymax int) {
	// the queue contains springs. When a container becomes full, it will pour off the side
	// and pouring off the side is effectively a new spring for the queue
	queue := []Point{Point{Spring.X, ymin}}

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]

		// continue to pour straight down until clay is reached
		x, y := curr.X, curr.Y
		for y <= ymax {
			p := Point{x, y}
			if _, ok := clay[p]; ok {
				// clay detected so start filling up
				break
			} else {
				water[p] = '|'
			}

			y++
		}

		// we're off the grid, so stop
		if y > ymax {
			continue
		}

		// not off the grid, so we must start to fill up
		l, r := true, true
		for i := 0; l || r; i++ {
			if l {
				p2 := Point{x - i, y}
				if _, ok := clay[p2]; ok {
					l = false
				} else {
					water[p2] = '~'
				}
			}
			if r {
				p1 := Point{x + i, y}
				if _, ok := clay[p1]; ok {
					r = false
				} else {
					water[p1] = '~'
				}
			}

		}
	}
}

func part1(fn string) {
	fmt.Println("Part 1")

	points, xmin, xmax, ymin, ymax := readInput(fn)

	// make a map showing the clay veins
	clay := make(map[Point]rune)
	for _, p := range points {
		clay[p] = '#'
	}
	// another map to store the water points
	water := make(map[Point]rune)

	printState(clay, water, xmin, xmax, ymin, ymax)

}

func part2() {
	fmt.Println("Part 2")
}
