package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "01"
	dayTitle = "Report Repair"
)

func part1() {
	fmt.Println("Part 1")
	inputs, err := util.ReadIntput("input1.txt")
	if err != nil {
		panic(err)
	}

	nums := make(map[int]struct{})
	for _, n := range inputs {
		nums[n] = struct{}{}

		m := 2020 - n
		if _, ok := nums[m]; ok {
			mult := n * m
			fmt.Println(mult)
			break
		}
	}

}

func part2() {
	fmt.Println("Part 2")
	inputs, _ := util.ReadIntput("input1.txt")
	nums := make(map[int]struct{})

outer:
	for i, n := range inputs {
		nums[n] = struct{}{}

		t := 2020 - n
		for j, nn := range inputs {
			if i == j {
				continue
			}

			m := t - nn
			if _, ok := nums[m]; ok {
				mult := n * m * nn
				fmt.Println(mult)
				break outer
			}
		}
	}
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
