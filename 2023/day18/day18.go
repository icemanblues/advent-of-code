package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/graph"
	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "18"
	dayTitle = "Lavaduct Lagoon"
)

type DigRecord struct {
	Dir    string
	Meters int
	Color  string
}

func Parse(filename string) []DigRecord {
	input, _ := util.ReadInput(filename)
	digs := make([]DigRecord, 0, len(input))
	for _, line := range input {
		fields := strings.Fields(line)
		digs = append(digs, DigRecord{fields[0], util.MustAtoi(fields[1]), fields[2][2:8]})
	}
	return digs
}

func move(p graph.Point2D, dir string, amount int) graph.Point2D {
	switch dir {
	case "R":
		return graph.NewPoint2D(p.X+amount, p.Y)
	case "U":
		return graph.NewPoint2D(p.X, p.Y-amount)
	case "D":
		return graph.NewPoint2D(p.X, p.Y+amount)
	case "L":
		return graph.NewPoint2D(p.X-amount, p.Y)
	}
	return p
}

func shoelace(vertices []graph.Point2D) int {
	s1, s2 := 0, 0
	for i := 0; i < len(vertices)-1; i++ {
		s1 += vertices[i].X * vertices[i+1].Y
		s2 += vertices[i].Y * vertices[i+1].X
	}
	return util.Abs(s1-s2) / 2
}

func polygonArea(digs []DigRecord) int {
	curr := graph.NewPoint2D(0, 0)
	trench := 0
	vertices := make([]graph.Point2D, 0, len(digs)+1)
	vertices = append(vertices, curr)
	for _, d := range digs {
		curr = move(curr, d.Dir, d.Meters)
		vertices = append(vertices, curr)
		trench += d.Meters
	}
	return shoelace(vertices) + trench/2 + 1
}

func part1() {
	digs := Parse("input.txt")
	fmt.Printf("Part 1: %v\n", polygonArea(digs))
}

func part2() {
	digs := Parse("input.txt")
	newDigs := make([]DigRecord, 0, len(digs))
	for _, dig := range digs {
		dist := dig.Color[:5]
		dir := dig.Color[5:]
		var d string
		switch dir {
		case "0":
			d = "R"
		case "1":
			d = "D"
		case "2":
			d = "L"
		case "3":
			d = "U"
		}
		meters, _ := strconv.ParseInt(dist, 16, 64)
		newDigs = append(newDigs, DigRecord{d, int(meters), dig.Color})
	}
	fmt.Printf("Part 2: %v\n", polygonArea(newDigs))
}

// https://www.themathdoctors.org/polygon-coordinates-and-areas/
func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
