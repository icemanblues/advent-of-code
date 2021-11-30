package main

import (
	"fmt"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "05"
	dayTitle = "Doesn't He Have Intern-Elves For This?"
)

func isVowel(r rune) bool {
	switch r {
	case 'a':
		return true
	case 'e':
		return true
	case 'i':
		return true
	case 'o':
		return true
	case 'u':
		return true
	}
	return false
}

var badWords []string = []string{
	"ab", "cd", "pq", "xy",
}

func noSubstring(s string) bool {
	for _, bad := range badWords {
		if strings.Index(s, bad) != -1 {
			return false
		}
	}
	return true
}

func nice(s string) bool {
	vowelCount := 0
	runes := []rune(s)
	twiceInRow := false
	for i, r := range runes {
		if isVowel(r) {
			vowelCount++
		}

		if i > 0 && r == runes[i-1] {
			twiceInRow = true
		}
	}

	threeVowels := vowelCount >= 3
	noBadWords := noSubstring(s)
	return twiceInRow && threeVowels && noBadWords
}

func betterNice(s string) bool {
	runes := []rune(s)
	repeat := false
	for i, r := range s {
		if i >= 2 && r == runes[i-2] {
			repeat = true
			break
		}
	}
	pair := false
	for i := 0; i < len(s)-1; i++ {
		sub := s[i : i+2]
		left := s[i+2:]
		if strings.Index(left, sub) != -1 {
			pair = true
			break
		}
	}
	return repeat && pair
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	lines, _ := util.ReadInput("input.txt")
	niceCount := 0
	betterNiceCount := 0
	for _, line := range lines {
		if nice(line) {
			niceCount++
		}
		if betterNice(line) {
			betterNiceCount++
		}
	}
	fmt.Printf("Part 1: %v\n", niceCount)
	fmt.Printf("Part 2: %v\n", betterNiceCount)
}
