package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "08"
	dayTitle = "Matchsticks"
)

func numCharsMemory(s string) (int, int) {
	mem, runes := 0, []rune(s)
	for i := 1; i < len(runes)-1; {
		if runes[i] == '\\' {
			switch runes[i+1] {
			case '\\':
				mem++
				i += 2
			case '"':
				mem++
				i += 2
			case 'x':
				mem++
				i += 4
			}
		} else {
			i++
			mem++
		}
	}
	return len(s), mem
}

func encodedNumCharsMemory(s string) (int, int) {
	mem := 0
	for _, r := range s {
		switch r {
		case '\\':
			mem += 2
		case '"':
			mem += 2
		default:
			mem++
		}
	}
	return len(s), mem + 2
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	lines := util.MustRead("input.txt")
	sumChars, sumMem := 0, 0
	sumEncodedChars, sumEncodedMem := 0, 0
	for _, line := range lines {
		chars, mem := numCharsMemory(line)
		sumChars += chars
		sumMem += mem

		encodedChars, encodedMem := encodedNumCharsMemory(line)
		sumEncodedChars += encodedChars
		sumEncodedMem += encodedMem
	}
	fmt.Printf("Part 1: %v\n", sumChars-sumMem)
	fmt.Printf("Part 2: %v\n", sumEncodedMem-sumEncodedChars)
}
