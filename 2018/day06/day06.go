package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	ID int
	X  int
	Y  int
}

func readInput(filename string) []Coord {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return nil
	}
	defer file.Close()

	num := 0
	var points []Coord
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num++
		s := scanner.Text()
		t := strings.Split(s, ", ")
		x, _ := strconv.Atoi(t[0])
		y, _ := strconv.Atoi(t[1])
		points = append(points, Coord{
			ID: num,
			X:  x,
			Y:  y,
		})
	}
	return points
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dist(c Coord, x, y int) int {
	return abs(c.X-x) + abs(c.Y-y)
}

// If you are on the perimeter, you are infinite
func isInfinite(grid [][]int) map[int]bool {
	infinite := make(map[int]bool)
	for i := range grid[0] {
		infinite[grid[0][i]] = true
		infinite[grid[len(grid)-1][i]] = true
	}
	for j := range grid {
		infinite[grid[j][0]] = true
		infinite[grid[j][len(grid[0])-1]] = true
	}

	return infinite
}

func main() {
	fmt.Println("Day 06: Chronal Coordinates")

	part1()
	part2()
}

func part1() {
	fmt.Println("Part 1")

	points := readInput("input06.txt")

	xmax, ymax := -1, -1
	for _, p := range points {
		if p.X > xmax {
			xmax = p.X
		}
		if p.Y > ymax {
			ymax = p.Y
		}
	}

	pnum := make(map[int]int)
	grid := make([][]int, xmax+1)
	for i, _ := range grid {
		grid[i] = make([]int, ymax+1)
		for j, _ := range grid[i] {
			dmin := 2 * (ymax + xmax)
			idmin := 0
			for _, p := range points {
				d := dist(p, i, j)

				if d == dmin {
					idmin = 0
				} else if d < dmin {
					dmin = d
					idmin = p.ID
				}
			}

			grid[i][j] = idmin
			pnum[idmin] = pnum[idmin] + 1
		}
	}

	// transpose for printing
	// for i := range grid[0] {
	// 	for j := range grid {
	// 		fmt.Print(grid[j][i])
	// 	}
	// 	fmt.Print("\n")
	// }
	// fmt.Println(pnum)

	// remove the infinites
	infinite := isInfinite(grid)
	for i := range infinite {
		delete(pnum, i)
	}

	// get the highest from the map
	var id, max int
	for k, v := range pnum {
		if max < v {
			max = v
			id = k
		}
	}

	fmt.Printf("largest area %v for point %v\n", max, id)
}

func part2() {
	fmt.Println("Part 2")

	points := readInput("input06.txt")
	const safe = 10000

	xmax, ymax := -1, -1
	for _, p := range points {
		if p.X > xmax {
			xmax = p.X
		}
		if p.Y > ymax {
			ymax = p.Y
		}
	}

	region := 0
	for i := 0; i < xmax+1; i++ {
		for j := 0; j < ymax+1; j++ {
			total := 0
			for _, p := range points {
				d := dist(p, i, j)
				total += d
			}

			if total < safe {
				region++
			}
		}
	}

	fmt.Println(region)
}
