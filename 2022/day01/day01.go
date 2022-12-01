package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "01"
	dayTitle = "Calorie Counting"
)

type Elf struct {
	Calories []int
	Total    int
}

func NewElf(calories []int) Elf {
	sum := 0
	for _, c := range calories {
		sum += c
	}
	return Elf{calories, sum}
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	input, _ := util.ReadInput("input.txt")
	var elves []Elf
	var calories []int
	for _, line := range input {
		if line == "" {
			elves = append(elves, NewElf(calories))
			calories = nil
			continue
		}
		a, _ := strconv.Atoi(line)
		calories = append(calories, a)
	}
	elves = append(elves, NewElf(calories))
	calories = nil

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].Total < elves[j].Total
	})
	l := len(elves) - 1
	max := elves[l].Total
	maxThree := elves[l].Total + elves[l-1].Total + elves[l-2].Total

	fmt.Printf("Part 1: %v\n", max)
	fmt.Printf("Part 2: %v\n", maxThree)
}
