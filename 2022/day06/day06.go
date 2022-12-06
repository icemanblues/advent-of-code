package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "06"
	dayTitle = "Tuning Trouble"
)

func FindMarker(s string, n int) int {
	runes := []rune(s)
	for i := 0; i < len(runes)-n; i++ {
		j := i + n
		marker := runes[i:j]

		set := make(map[rune]struct{})
		for _, r := range marker {
			set[r] = struct{}{}
		}
		if len(set) == n {
			return j
		}
	}
	return -1
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	input, _ := util.ReadInput("input.txt")
	s := input[0]
	fmt.Printf("Part 1: %v\n", FindMarker(s, 4))
	fmt.Printf("Part 2: %v\n", FindMarker(s, 14))
}
