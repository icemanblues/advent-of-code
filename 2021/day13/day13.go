package main

import (
	"fmt"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "13"
	dayTitle = "Transparent Origami"
)

type Point struct {
	x, y int
}

type PointSet map[Point]struct{}

type Fold struct {
	dir string
	num int
}

func Parse(filename string) (PointSet, []Fold) {
	input, _ := util.ReadInput(filename)
	isFold := false
	dots := make(PointSet)
	var folds []Fold
	for _, line := range input {
		if line == "" {
			isFold = true
			continue
		}

		if isFold {
			line = strings.ReplaceAll(line, "fold along ", "")
			parts := strings.Split(line, "=")
			f := Fold{parts[0], util.MustAtoi(parts[1])}
			folds = append(folds, f)
		} else {
			parts := strings.Split(line, ",")
			p := Point{util.MustAtoi(parts[0]), util.MustAtoi(parts[1])}
			dots[p] = struct{}{}
		}
	}
	return dots, folds
}

func transform(dots PointSet, f Fold) PointSet {
	t := make(PointSet)
	for p := range dots {
		if f.dir == "y" && p.y > f.num {
			n := Point{p.x, f.num - (p.y - f.num)}
			t[n] = struct{}{}
		} else if f.dir == "x" && p.x > f.num {
			n := Point{f.num - (p.x - f.num), p.y}
			t[n] = struct{}{}
		} else {
			t[p] = struct{}{}
		}
	}
	return t
}

func part1() {
	dots, folds := Parse("input.txt")
	dots = transform(dots, folds[0])
	fmt.Printf("Part 1: %v\n", len(dots))
}

func printDots(points PointSet) {
	maxX, maxY := -1, -1
	for p := range points {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			n := Point{x, y}
			if _, ok := points[n]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func part2() {
	dots, folds := Parse("input.txt")
	for _, f := range folds {
		dots = transform(dots, f)
	}
	printDots(dots)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
