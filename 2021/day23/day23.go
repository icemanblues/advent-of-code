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
	fmt.Println("Part 1")
	input, _ := util.ReadInput("input.txt")
	fmt.Println(input)
}

func part2() {
	fmt.Println("Part 2")
	runes, _ := util.ReadRuneput("input.txt")
	fmt.Println(runes)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
