package main

import (
	"fmt"
	"sort"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

func joltDiff(adapters []int) int {
	sort.Ints(adapters)
	curr := 0
	diff1 := 0
	diff3 := 1
	for _, jolt := range adapters {
		if jolt-curr == 1 {
			diff1++
			curr = jolt
		}
		if jolt-curr == 3 {
			diff3++
			curr = jolt
		}
	}
	return diff1 * diff3
}

func numPaths(adapters []int) int {
	sort.Ints(adapters)
	paths := make(map[int]int)
	paths[0] = 1
	for _, jolt := range adapters {
		paths[jolt] = paths[jolt-1] + paths[jolt-2] + paths[jolt-3]
	}
	return paths[adapters[len(adapters)-1]]
}

func part1() {
	ints, _ := util.ReadIntput("input.txt")
	fmt.Printf("Part 1, %v\n", joltDiff(ints))
}

func part2() {
	ints, _ := util.ReadIntput("input.txt")
	fmt.Printf("Part 2: %v\n", numPaths(ints))
}

func main() {
	fmt.Println("Day 10: Adapter Array")
	part1()
	part2()
}
