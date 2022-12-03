package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "03"
	dayTitle = "Rucksack Reorganization"
)

type Set map[rune]struct{}

type RuckSack struct {
	First, Second Set
	Overlap       rune
}

func (rs RuckSack) ToSet() Set {
	set := make(Set)
	for k := range rs.First {
		set[k] = struct{}{}
	}
	for k := range rs.Second {
		set[k] = struct{}{}
	}
	return set
}

func NewRuckSack(first, second Set) RuckSack {
	var overlap rune
	for r := range first {
		if _, ok := second[r]; ok {
			overlap = r
			break
		}
	}
	return RuckSack{first, second, overlap}
}

func MakeSet(runes []rune) Set {
	s := make(Set)
	for _, r := range runes {
		s[r] = struct{}{}
	}
	return s
}

func parse(filename string) []RuckSack {
	lineRunes, _ := util.ReadRuneput(filename)
	sacks := make([]RuckSack, 0, len(lineRunes))
	for _, line := range lineRunes {
		f := line[:len(line)/2]
		s := line[len(line)/2:]
		sacks = append(sacks, NewRuckSack(MakeSet(f), MakeSet(s)))
	}
	return sacks
}

var ScoreLine string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func score(r rune) int {
	for i, s := range ScoreLine {
		if s == r {
			return i + 1
		}
	}
	return 0
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	sacks := parse("input.txt")

	sum := 0
	for _, s := range sacks {
		sum += score(s.Overlap)
	}
	fmt.Printf("Part 1: %v\n", sum)

	badgeSum := 0
	for i := 0; i < len(sacks); i = i + 3 {
		a, b, c := sacks[i].ToSet(), sacks[i+1].ToSet(), sacks[i+2].ToSet()
		for k := range a {
			_, bok := b[k]
			_, cok := c[k]
			if bok && cok {
				badgeSum += score(k)
			}
		}
	}
	fmt.Printf("Part 2: %v\n", badgeSum)
}
