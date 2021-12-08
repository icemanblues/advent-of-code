package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "07"
	dayTitle = "The Treachery of Whales"
)

type DistFormula func(int, int) int

func dist(a, b int) int {
	return util.Abs(a - b)
}

func sumNum(a, b int) int {
	a = util.Abs(a - b)
	return (a * (a + 1)) / 2
}

func search(input []int, f DistFormula) (int, int) {
	// establish our search boundaries (inclusive)
	min, max := input[0], input[0]
	for i := 1; i < len(input); i++ {
		if input[i] < min {
			min = input[i]
		}
		if input[i] > max {
			max = input[i]
		}
	}

	// find the lowest fuel via the distance function
	minFuel := -1
	minIndex := -1
	for i := min; i <= max; i++ {
		fuel := 0
		for _, e := range input {
			fuel += f(e, i)
		}
		if minFuel == -1 || minFuel > fuel {
			minFuel = fuel
			minIndex = i
		}
	}

	return minFuel, minIndex
}

func part1() {
	input, _ := util.ReadIntLine("input.txt", ",")
	minFuel, minIndex := search(input, dist)
	fmt.Printf("Part 1: fuel: %v index: %v\n", minFuel, minIndex)
}

func part2() {
	input, _ := util.ReadIntLine("input.txt", ",")
	minFuel, minIndex := search(input, sumNum)
	fmt.Printf("Part 2: fuel: %v index: %v\n", minFuel, minIndex)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
