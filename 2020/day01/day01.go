package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "XX"
	dayTitle = "Title"
)

func part1() {
	fmt.Println("Part 1")
	inputs, err := util.ReadInput("input1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(inputs)
}

func part2() {
	fmt.Println("Part 2")
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
