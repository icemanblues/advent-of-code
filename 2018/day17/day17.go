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
	visited := make(map[Point]struct{})
	// tick := 0

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]

		if _, ok := visited[curr]; ok {
			// fmt.Println("[WARN WARN WARN] already processed this spout, so skip it", curr)
			continue
		}
		visited[curr] = struct{}{}

		x, y := curr.X, curr.Y

		// continue to pour straight down until clay is reached
		for y <= ymax {
			p := Point{x, y}

			// fill down until clay or resting water
			// if r, restBot := water[p]; restBot && r == '~' {
			// 	y++
			// 	break
			// }

			// fill down until clay
			if _, clayBot := clay[p]; clayBot {
				break
			} else {
				water[p] = '|'
			}

			y++
		}

		// tick++
		// fmt.Println(tick, "+++ we're on spout number +++", tick, curr)
		// fmt.Printf("Currently at depth %d out of %d. Only %d left to go\n", y, ymax, ymax-y)
		// fmt.Println("number of water tiles", len(water))
		// fmt.Println("number of spouts remaining", len(queue))

		// we're off the grid, so stop
		if y > ymax {
			continue
		}

		// not off the grid, so we must start to fill up (left to right)
		fill := true
		for fill {
			y--

			// should never make it here
			if y < ymin {
				fmt.Println("!!! [ERROR] when filling up, we are passed ymin. How did we do this?")
				break
			}

			l, r := true, true
			maxr, maxl := x, x
			for i := 0; l || r; i++ {
				if l {
					pl := Point{x - i, y}
					bot := Point{x - i, y + 1}
					_, clayLeft := clay[pl]
					_, clayBot := clay[bot]
					r, restBot := water[bot]
					clayBot = clayBot || (restBot && r == '~')

					if clayLeft && clayBot {
						l = false
						maxl = pl.X
					} else if clayBot {
						water[pl] = '~'
					} else if !clayBot {
						// this point becomes a spout
						water[pl] = '|'
						l = false
						maxl = pl.X
						fill = false
						queue = append(queue, pl)
					}
				}
				if r {
					pr := Point{x + i, y}
					bot := Point{x + i, y + 1}
					_, clayRight := clay[pr]
					_, clayBot := clay[bot]
					ru, restBot := water[bot]
					clayBot = clayBot || (restBot && ru == '~')

					if clayRight && clayBot {
						r = false
						maxr = pr.X
					} else if clayBot {
						water[pr] = '~'
					} else if !clayBot {
						// this point becomes a spout
						water[pr] = '|'
						fill = false
						maxr = pr.X
						r = false
						queue = append(queue, pr)
					}
				}

				// fmt.Printf("fill %v left %v right %v at depth %v\n", fill, l, r, y)
				// printState(clay, water, xmin, xmax, ymin, ymax)
			}

			if !fill {
				// this means that we created a spout
				for i := maxl + 1; i <= maxr; i++ {
					p := Point{i, y}
					water[p] = '|'
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

	fmt.Println(xmin, xmax, ymin, ymax)
	pour(clay, water, xmin, xmax, ymin, ymax)

	fmt.Printf("%d water tiles\n", len(water))

	sum := 0
	for _, r := range water {
		if r == '~' {
			sum++
		}
	}
	fmt.Printf("resting water count %d\n", sum)

	// printState(clay, water, xmin, xmax, ymin, ymax)
}
