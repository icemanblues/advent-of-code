package main

import (
	"fmt"
	"strconv"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "03"
	dayTitle = "Binary Diagnostic"
)

func part1() {
	input, _ := util.ReadRuneput("input.txt")

	l := len(input) / 2
	c := make([]int, len(input[0]), len(input[0]))
	for _, runes := range input {
		for i := range runes {
			if runes[i] == '1' {
				c[i]++
			}
		}
	}

	gamma := make([]rune, len(c), len(c))
	epsilon := make([]rune, len(c), len(c))
	for i, v := range c {
		if v >= l {
			gamma[i] = '1'
			epsilon[i] = '0'
		} else {
			gamma[i] = '0'
			epsilon[i] = '1'
		}
	}

	g, _ := strconv.ParseInt(string(gamma), 2, 64)
	e, _ := strconv.ParseInt(string(epsilon), 2, 64)
	power := g * e
	fmt.Printf("Part 1: %v\n", power)
}

func part2() {
	input, _ := util.ReadRuneput("input.txt")
	length := len(input[0])

	oxygen := make([][]rune, len(input), len(input))
	copy(oxygen, input)
	for i := 0; i < length; i++ {
		if len(oxygen) == 1 {
			break
		}

		on, off := 0, 0
		for _, runes := range oxygen {
			if runes[i] == '1' {
				on++
			}
			if runes[i] == '0' {
				off++
			}
		}

		bit := '1'
		if on >= off {
			bit = '1'
		} else {
			bit = '0'
		}

		var iter [][]rune
		for _, v := range oxygen {
			if v[i] == bit {
				iter = append(iter, v)
			}
		}
		oxygen = iter
	}

	co := make([][]rune, len(input), len(input))
	copy(co, input)
	for i := 0; i < length; i++ {
		if len(co) == 1 {
			break
		}

		on, off := 0, 0
		for _, runes := range co {
			if runes[i] == '1' {
				on++
			}
			if runes[i] == '0' {
				off++
			}
		}

		bit := '0'
		if off <= on {
			bit = '0'
		} else {
			bit = '1'
		}

		var iter [][]rune
		for _, v := range co {
			if v[i] == bit {
				iter = append(iter, v)
			}
		}
		co = iter
	}

	oRating, _ := strconv.ParseInt(string(oxygen[0]), 2, 64)
	coRating, _ := strconv.ParseInt(string(co[0]), 2, 64)
	life := oRating * coRating
	fmt.Printf("Part 2: %v\n", life)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
