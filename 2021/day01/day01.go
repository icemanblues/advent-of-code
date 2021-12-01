package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "01"
	dayTitle = "Sonar Sweep"
)

func part1() {
	fmt.Println("Part 1")
	input, _ := util.ReadIntput("input.txt")

	prev := input[0]
	incCount := 0
	for i := 1; i < len(input); i++ {
		if input[i] > prev {
			incCount++
		}
		prev = input[i]
	}
	fmt.Println(incCount)
}

func part2() {
	fmt.Println("Part 2")
	input, _ := util.ReadIntput("input.txt")

	prev := input[0] + input[1] + input[2]
	incCount := 0
	for i := 1; i < len(input)-2; i++ {
		window := input[i] + input[i+1] + input[i+2]
		if window > prev {
			incCount++
		}
		prev = window
	}
	fmt.Println(incCount)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
