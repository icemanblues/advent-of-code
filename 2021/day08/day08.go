package main

import (
	"fmt"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "08"
	dayTitle = "Seven Segment Search"
)

type Segment struct {
	in, out []string
}

var siglenToNum map[int]int = map[int]int{
	2: 1,
	4: 4,
	3: 7,
	7: 8,
}

func Parse(filename string) []Segment {
	input, _ := util.ReadInput(filename)
	var segments []Segment
	for _, in := range input {
		parts := strings.Split(in, " | ")
		i := strings.Split(parts[0], " ")
		j := strings.Split(parts[1], " ")
		segments = append(segments, Segment{i, j})
	}
	return segments
}

func part1() {
	input := Parse("input.txt")
	count := 0
	for _, segment := range input {
		for _, out := range segment.out {
			if _, ok := siglenToNum[len(out)]; ok {
				count++
			}
		}
	}

	fmt.Printf("Part 1: %v\n", count)
}

func part2() {
	fmt.Println("Part 2")
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
