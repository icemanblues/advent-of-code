package main

import (
	"fmt"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "02"
	dayTitle = "Dive!"
)

func part1() {
	fmt.Println("Part 1")
	input, _ := util.ReadInput("input.txt")

	horizontal, depth := 0, 0
	for _, step := range input {
		splits := strings.Split(step, " ")
		switch splits[0] {
		case "forward":
			horizontal += util.MustAtoi(splits[1])
		case "down":
			depth += util.MustAtoi(splits[1])
		case "up":
			depth -= util.MustAtoi(splits[1])
		default:
			fmt.Printf("unknown direction: %v\n", input)
		}
	}

	answer := horizontal * depth
	fmt.Println(answer)
}

func part2() {
	fmt.Println("Part 2")
	input, _ := util.ReadInput("input.txt")

	horizontal, depth, aim := 0, 0, 0
	for _, step := range input {
		splits := strings.Split(step, " ")
		switch splits[0] {
		case "forward":
			horizontal += util.MustAtoi(splits[1])
			depth += aim * util.MustAtoi(splits[1])
		case "down":
			aim += util.MustAtoi(splits[1])
		case "up":
			aim -= util.MustAtoi(splits[1])
		default:
			fmt.Printf("unknown direction: %v\n", input)
		}
	}

	answer := horizontal * depth
	fmt.Println(answer)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
