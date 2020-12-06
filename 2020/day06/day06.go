package main

import (
	"fmt"
	"strings"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "06"
	dayTitle = "Custom Customs"
)

func answerCount(lines []string) int {
	sum := 0
	answerSet := make(map[rune]struct{})

	for _, line := range lines {
		// add it to the current set
		if len(strings.TrimSpace(line)) != 0 {
			for _, r := range line {
				answerSet[r] = struct{}{}
			}
		} else { // find the count, reset and continue
			sum = sum + len(answerSet)
			answerSet = make(map[rune]struct{})
		}
	}

	sum = sum + len(answerSet)
	return sum
}

func allAnswerCount(lines []string) int {
	sum := 0

	size := 0
	answerSet := make(map[rune]int)

	for _, line := range lines {
		// add it to the current set
		if len(strings.TrimSpace(line)) != 0 {
			for _, r := range line {
				answerSet[r] = answerSet[r] + 1
			}
			size++
		} else { // find the count, reset and continue
			total := 0
			//			fmt.Printf("size: %v len(answerSet): %v answerSet: %v\n", size, len(answerSet), answerSet)
			for _, c := range answerSet {
				if size == c {
					//fmt.Println("adding it")
					total++
				}
			}
			sum = sum + total

			answerSet = make(map[rune]int)
			size = 0
		}
	}

	total := 0
	for _, c := range answerSet {
		if size == c {
			total++
		}
	}
	sum = sum + total

	return sum
}

func part1() {
	lines, _ := util.ReadInput("input.txt")
	fmt.Printf("Part 1: %v\n", answerCount(lines))
}

func part2() {
	lines, _ := util.ReadInput("input.txt")
	fmt.Printf("Part 2: %v\n", allAnswerCount(lines))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
