package main

import (
	"fmt"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "05"
	dayTitle = "Hydrothermal Venture"
)

type Point struct {
	X, Y int
}

type Segment struct {
	Start, End Point
}

func parseSegments(filename string) []Segment {
	input, _ := util.ReadInput(filename)
	segments := make([]Segment, 0, len(input))
	for _, line := range input {
		parts := strings.Split(line, " -> ")
		s := strings.Split(parts[0], ",")
		e := strings.Split(parts[1], ",")
		start := Point{util.MustAtoi(s[0]), util.MustAtoi(s[1])}
		end := Point{util.MustAtoi(e[0]), util.MustAtoi(e[1])}
		segments = append(segments, Segment{start, end})
	}
	return segments
}

func countDanger(vents map[Point]int) int {
	answer := 0
	for _, v := range vents {
		if v > 1 {
			answer++
		}
	}
	return answer
}

func printVents(vents map[Point]int, maxX, maxY int) {
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			p := Point{x, y}
			c, ok := vents[p]
			if !ok {
				fmt.Printf(".")
			} else {
				fmt.Printf("%v", c)
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	segments := parseSegments("input.txt")
	vents := make(map[Point]int)
	for _, seg := range segments {
		// iterate on the Y axis
		if seg.Start.X == seg.End.X {
			s := util.Min(seg.Start.Y, seg.End.Y)
			e := util.Max(seg.Start.Y, seg.End.Y)
			for i := s; i <= e; i++ {
				p := Point{seg.Start.X, i}
				vents[p]++
			}
			// iterate on the X axis
		} else if seg.Start.Y == seg.End.Y {
			s := util.Min(seg.Start.X, seg.End.X)
			e := util.Max(seg.Start.X, seg.End.X)
			for i := s; i <= e; i++ {
				p := Point{i, seg.Start.Y}
				vents[p]++
			}
		}
	}
	fmt.Printf("Part 1: %v\n", countDanger(vents))

	for _, seg := range segments {
		// diagonals
		if seg.Start.X != seg.End.X && seg.Start.Y != seg.End.Y {
			xi := 1
			if seg.Start.X > seg.End.X {
				xi = -1
			}
			yi := 1
			if seg.Start.Y > seg.End.Y {
				yi = -1
			}

			for i := 0; i <= util.Abs(seg.End.X-seg.Start.X); i++ {
				p := Point{seg.Start.X + i*xi, seg.Start.Y + i*yi}
				vents[p]++
			}
		}
	}
	fmt.Printf("Part 2: %v\n", countDanger(vents))
}
