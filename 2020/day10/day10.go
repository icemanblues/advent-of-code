package main

import (
	"fmt"
	"sort"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "10"
	dayTitle = "Adapter Array"
)

func part1() {
	curr := 0
	ints, _ := util.ReadIntput("input.txt")
	sort.Ints(ints)

	diff1 := 0
	diff3 := 1
	for _, jolt := range ints {
		if jolt-curr == 1 {
			diff1++
			curr = jolt
		}
		if jolt-curr == 3 {
			diff3++
			curr = jolt
		}
	}

	fmt.Printf("Part 1, %v\n", diff1*diff3)
}

func part2() {
	ints, _ := util.ReadIntput("input.txt")
	sort.Ints(ints)
	paths := make(map[int]int)
	paths[0] = 1

	for _, jolt := range ints {
		paths[jolt] = paths[jolt-1] + paths[jolt-2] + paths[jolt-3]
	}

	fmt.Println(paths)
	fmt.Printf("Part 2: %v\n", paths[ints[len(ints)-1]])
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
