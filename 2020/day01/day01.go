package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "01"
	dayTitle = "Report Repair"
)

func find2sum(inputs []int, target int) (int, bool) {
	nums := make(map[int]struct{})
	for _, n := range inputs {
		nums[n] = struct{}{}

		m := target - n
		if _, ok := nums[m]; ok {
			return n * m, true
		}
	}

	return 0, false
}

func find3sum(inputs []int, target int) (int, bool) {
	for i, n := range inputs {
		in := append(inputs[:i], inputs[i+1:]...)
		m, ok := find2sum(in, target-n)
		if ok {
			return n * m, true
		}
	}
	return 0, false
}

func part1() {
	inputs, _ := util.ReadIntput("input1.txt")
	a, _ := find2sum(inputs, 2020)
	fmt.Printf("Part 1: %v\n", a)
}

func part2() {
	inputs, _ := util.ReadIntput("input1.txt")
	a, _ := find3sum(inputs, 2020)
	fmt.Printf("Part 2: %v\n", a)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
