package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "04"
	dayTitle = "Camp Cleanup"
)

type Range struct {
	X, Y int
}

func (r Range) FullyContains(s Range) bool {
	return r.X <= s.X && r.Y >= s.Y
}

func (r Range) Overlaps(s Range) bool {
	if s.X >= r.X && s.X <= r.Y {
		return true
	}
	if s.Y >= r.X && s.Y <= r.Y {
		return true
	}
	return false
}

func parse(filename string) [][]Range {
	input, _ := util.ReadInput(filename)
	ranges := make([][]Range, 0, len(input))
	for _, line := range input {
		pair := strings.Split(line, ",")
		r := make([]Range, 0, len(pair))
		for _, p := range pair {
			a := strings.Split(p, "-")
			a1, _ := strconv.Atoi(a[0])
			a2, _ := strconv.Atoi(a[1])
			r = append(r, Range{a1, a2})
		}
		ranges = append(ranges, r)
	}
	return ranges
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	ranges := parse("input.txt")

	fullyContainsCount := 0
	overlapsCount := 0
	for _, r := range ranges {
		if r[0].FullyContains(r[1]) || r[1].FullyContains(r[0]) {
			fullyContainsCount++
		}
		if r[0].Overlaps(r[1]) || r[1].Overlaps(r[0]) {
			overlapsCount++
		}
	}
	fmt.Printf("Part 1: %v\n", fullyContainsCount)
	fmt.Printf("Part 2: %v\n", overlapsCount)
}
