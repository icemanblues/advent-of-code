package main

import (
	"fmt"
	"sort"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "10"
	dayTitle = "Syntax Scoring"
)

var scoring map[rune]int = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var autocomplete map[rune]int = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

var closeToOpen map[rune]rune = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

func score(chunk []rune) (int, int) {
	var stack []rune
	for _, r := range chunk {
		if _, ok := closeToOpen[r]; !ok { // opening bracket
			stack = append(stack, r)
		} else { // closing bracket
			if len(stack) == 0 {
				return scoring[r], 0
			}

			open := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if open != closeToOpen[r] {
				return scoring[r], 0
			}
		}
	}

	if len(stack) == 0 {
		return 0, 0
	}

	// its incomplete, so we score it differently (part2)
	autoscore := 0
	for i := len(stack) - 1; i >= 0; i-- {
		autoscore = (5 * autoscore) + autocomplete[stack[i]]
	}
	return 0, autoscore
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	runes, _ := util.ReadRuneput("input.txt")
	sum := 0
	var autocomplete []int
	for _, line := range runes {
		s, a := score(line)
		sum += s
		if a != 0 {
			autocomplete = append(autocomplete, a)
		}
	}
	fmt.Printf("Part 1: %v\n", sum)

	sort.Ints(autocomplete)
	fmt.Printf("Part 2: %v\n", autocomplete[len(autocomplete)/2])
}
