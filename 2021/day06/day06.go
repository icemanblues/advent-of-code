package main

import (
	"fmt"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "06"
	dayTitle = "Lanternfish"
)

func Parse(filename string) []int {
	input, _ := util.ReadInput(filename)
	fishes := strings.Split(input[0], ",")
	var laternfish []int
	for _, f := range fishes {
		laternfish = append(laternfish, util.MustAtoi(f))
	}
	return laternfish
}

func spawn(laternfish []int, days int) []int {
	for i := 0; i < days; i++ {
		newCount := 0
		var newFish []int
		for _, f := range laternfish {
			if f == 0 {
				newCount++
				newFish = append(newFish, 6)
			} else {
				newFish = append(newFish, f-1)
			}
		}
		for j := 0; j < newCount; j++ {
			newFish = append(newFish, 8)
		}
		laternfish = newFish
	}

	return laternfish
}

func part1() {
	laternfish := Parse("input.txt")
	days := 80
	laternfish = spawn(laternfish, days)
	fmt.Printf("Part 1: %v\n", len(laternfish))
}

func part2() {
	input := Parse("input.txt")
	laternfish := make([]int64, 9, 9)
	for _, fish := range input {
		laternfish[fish]++
	}

	days := 256
	for i := 0; i < days; i++ {
		births := laternfish[0]
		laternfish = laternfish[1:]
		laternfish[6] += births
		laternfish = append(laternfish, births)
	}

	var sum int64 = 0
	for _, v := range laternfish {
		sum += v
	}
	fmt.Printf("Part 2: %v\n", sum)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
