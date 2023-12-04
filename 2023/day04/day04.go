package main

import (
	"fmt"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "04"
	dayTitle = "Scratchcards"
)

type IntSet map[int]struct{}

func (is IntSet) intersection(os IntSet) int {
	count := 0
	for k := range os {
		if _, ok := is[k]; ok {
			count++
		}
	}
	return count
}

type Scratchcard struct {
	ID            int
	Winning, Hand IntSet
}

func (sc Scratchcard) matches() int {
	return sc.Winning.intersection(sc.Hand)
}

func (sc Scratchcard) score() int {
	matched := sc.matches()
	if matched > 0 {
		return util.IntPow(2, matched-1)
	}
	return 0
}

func NewIntSet(s []string) IntSet {
	is := make(IntSet)
	for _, r := range s {
		i := util.MustAtoi(string(r))
		is[i] = struct{}{}
	}
	return is
}

func parseScratchcards(filename string) []Scratchcard {
	input, _ := util.ReadInput(filename)
	sc := make([]Scratchcard, 0, len(input))
	for _, line := range input {
		idWinHand := strings.Split(line, ": ")
		id := util.MustAtoi(strings.Fields(idWinHand[0])[1])

		winHands := strings.Split(idWinHand[1], " | ")
		win, hands := NewIntSet(strings.Fields(winHands[0])), NewIntSet(strings.Fields(winHands[1]))
		sc = append(sc, Scratchcard{ID: id, Winning: win, Hand: hands})
	}
	return sc
}

func part1() {
	scratchcards := parseScratchcards("input.txt")
	sum := 0
	for _, sc := range scratchcards {
		sum += sc.score()
	}
	fmt.Printf("Part 1: %v\n", sum)
}

func part2() {
	scratchcards := parseScratchcards("input.txt")
	cardCounts := make([]int, len(scratchcards), len(scratchcards))
	for i := range cardCounts {
		cardCounts[i] = 1
	}
	for i, sc := range scratchcards {
		mult := cardCounts[i]
		for j := 0; j < sc.matches(); j++ {
			if i+j+1 >= len(cardCounts) {
				break
			}
			cardCounts[i+j+1] += mult
		}
	}
	sum := 0
	for _, v := range cardCounts {
		sum += v
	}
	fmt.Printf("Part 2: %v\n", sum)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
