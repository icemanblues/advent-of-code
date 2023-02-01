package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "XX"
	dayTitle = "Title"
)

func part1() {
	input, _ := util.ReadInput("input.txt")
	fmt.Printf("Part 1: %v\n", len(input))
}

func part2() {
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
