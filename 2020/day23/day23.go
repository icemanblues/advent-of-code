package main

import (
	"fmt"
	"math"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "23"
	dayTitle = "Title"
)

const input string = "974618352"

func ParseCups(s string) []int {
	var cups []int
	for _, r := range s {
		cups = append(cups, util.MustAtoi(string(r)))
	}
	return cups
}

func PlayGame(cups []int, numMoves int) []int {
	l := len(cups)
	currIdx := 0
	min, max := cups[0], cups[0]
	for _, c := range cups {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}

	for m := 1; m <= numMoves; m++ {
		curr := cups[currIdx]
		idx1 := (currIdx + 1) % l
		idx2 := (currIdx + 2) % l
		idx3 := (currIdx + 3) % l
		cup1 := cups[idx1]
		cup2 := cups[idx2]
		cup3 := cups[idx3]

		// find the label
		label := curr - 1
		if label < min {
			label = max
		}
		for label == cup1 || label == cup2 || label == cup3 {
			label--
			if label < min {
				label = max
			}
		}

		// remove the pick ups
		for i, c := range cups {
			if c == cup1 {
				cups = append(cups[:i], cups[i+1:]...)
				break
			}
		}
		for i, c := range cups {
			if c == cup2 {
				cups = append(cups[:i], cups[i+1:]...)
				break
			}
		}
		for i, c := range cups {
			if c == cup3 {
				cups = append(cups[:i], cups[i+1:]...)
				break
			}
		}

		// find the index of the label
		labelIdx := -1
		for i, c := range cups {
			if c == label {
				labelIdx = i
				break
			}
		}

		// add them back in by the label
		newCups := make([]int, 0, l)
		insertPoint := (labelIdx + 1) % l
		newCups = append(newCups, cups[:insertPoint]...)
		newCups = append(newCups, cup1, cup2, cup3)
		newCups = append(newCups, cups[insertPoint:]...)
		cups = newCups

		// increment curr and currIdx
		for i, c := range cups {
			if c == curr {
				currIdx = i
			}
		}
		currIdx = (currIdx + 1) % l
	}

	return cups
}

func score(cups []int) int {
	var labels []int
	foundOne := false
	it := 0
	score := 0
	exp := len(cups) - 2
	for len(labels) != len(cups)-1 {
		c := cups[it]
		if foundOne {
			labels = append(labels, c)
			score += c * int(math.Pow10(exp))
			exp--
		}

		if c == 1 {
			foundOne = true
		}
		it = (it + 1) % len(cups)
	}
	return score
}

func printCups(move int, cups []int, currIdx int, pickups []int, dest int) {
	fmt.Printf("--Move %v --\n", move)
	fmt.Printf("cups: ")
	for i, c := range cups {
		if i == currIdx {
			fmt.Printf("(%v)", c)
		} else {
			fmt.Printf("%v", c)
		}
		if i != len(cups)-1 {
			fmt.Printf(" ")
		} else {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("pick up: %v\n", pickups)
	fmt.Printf("destination: %v\n", dest)
	fmt.Println()
}

func part1() {
	cups := ParseCups(input)
	cups = PlayGame(cups, 100)
	labels := score(cups)
	fmt.Printf("Part 1: %v\n", labels)
}

// 75893264
func part2() {
	fmt.Printf("Part 2: %v\n", 2)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
