package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Star struct {
	X  int
	Y  int
	Dx int
	Dy int
}

func readInput(filename string) []*Star {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return nil
	}
	defer file.Close()

	var stars []*Star
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// position=<-43373,  43655> velocity=< 4, -4>
		line := scanner.Text()

		vel := strings.Split(line, "velocity=")
		velocity := strings.Trim(vel[1], "<>")
		// fmt.Printf("velocity: %v\n", velocity)
		dx, dy := parse(velocity)

		pos := strings.Split(vel[0], "position=")
		position := strings.Trim(pos[1], "<> ")
		// fmt.Printf("position: %v\n", position)
		x, y := parse(position)

		stars = append(stars, &Star{
			X:  x,
			Y:  y,
			Dx: dx,
			Dy: dy,
		})
	}
	return stars
}

func parse(s string) (int, int) {
	digits := strings.Split(s, ",")
	x, _ := strconv.Atoi(strings.Trim(digits[0], " "))
	y, _ := strconv.Atoi(strings.Trim(digits[1], " "))
	return x, y
}

func main() {
	fmt.Println("Day 10: The Stars Align")

	part1()
	part2()
}

func part1() {
	fmt.Println("Part 1")

	stars := readInput("input10.txt")

	fmt.Printf("number of stars: %v\n", len(stars))

	t := 0
	p := 0
	for p < 10 {
		// print
		fmt.Printf("--- Time %v ---\n", t)
		if printStars(stars) {
			p++
		}

		// increment the stars
		for _, star := range stars {

			star.X = star.X + star.Dx
			star.Y = star.Y + star.Dy
		}
		// increment time
		t++
	}
}

type Point struct {
	X int
	Y int
}

func printStars(stars []*Star) bool {
	grid := make(map[Point]bool)
	minX := stars[0].X
	maxX := stars[0].X
	minY := stars[0].Y
	maxY := stars[0].Y
	for _, star := range stars {
		grid[Point{star.X, star.Y}] = true
		if star.X < minX {
			minX = star.X
		}
		if star.X > maxX {
			maxX = star.X
		}
		if star.Y < minY {
			minY = star.Y
		}
		if star.Y > maxY {
			maxY = star.Y
		}
	}

	// only print if it fits on the terminal
	const limit = 100
	result := (maxX-minX) <= limit && (maxY-minY) <= limit
	if result {
		for j := minY; j <= maxY; j++ {
			for i := minX; i <= maxX; i++ {
				if _, ok := grid[Point{i, j}]; ok {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			}

			fmt.Println()
		}
	}

	return result
}

func part2() {
	fmt.Println("Part 2")
}
