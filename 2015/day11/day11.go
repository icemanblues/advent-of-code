package main

import (
	"fmt"
)

const (
	dayNum   = "11"
	dayTitle = "Corporate Policy"
)

const input = "hepxcrrq"

var badRunes map[rune]struct{} = map[rune]struct{}{
	'i': {}, 'o': {}, 'l': {},
}

var base26 []rune = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l',
	'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
}

var map26 map[rune]int = make(map[rune]int)

func init() {
	for i, r := range base26 {
		map26[r] = i
	}
}

func validPassword(s string) bool {
	runes := []rune(s)
	// one increasing straight
	for i := 0; i < len(runes)-2; i++ {
		r := runes[i]
		rr := runes[i+1]
		rrr := runes[i+2]

	}

	// cannot contain i, o, l
	containsIOL := false
	for _, r := range s {
		if _, ok := badRunes[r]; ok {
			containsIOL = true
			break
		}
	}

	// two different, non-overlapping pairs
	pairs := make(map[rune]struct{})
	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == runes[i+1] {
			pairs[runes[i]] = struct{}{}
		}
	}
	hasTwoPairs := len(pairs) >= 2

	return hasTwoPairs && !containsIOL
}

func part1() {
	fmt.Println("Part 1")
	fmt.Printf("%v\n", len(map26))
}

func part2() {
	fmt.Println("Part 2")
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
