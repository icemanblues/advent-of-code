package main

import (
	"fmt"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "02"
	dayTitle = "Cube Conundrum"
)

type Color int

const (
	Blue Color = iota
	Red
	Green
)

func ColorFromString(s string) Color {
	switch s {
	case "blue":
		return Blue
	case "green":
		return Green
	case "red":
		return Red
	default:
		panic("unknown color")
	}
}

type Set map[Color]int

type Game struct {
	ID   int
	Sets []Set
}

func parseGame(line string) Game {
	idAndSets := strings.Split(line, ": ")
	id := util.MustAtoi(strings.Split(idAndSets[0], " ")[1])
	setLines := strings.Split(idAndSets[1], "; ")
	sets := make([]Set, 0, len(setLines))
	for _, setLine := range setLines {
		set := make(Set)
		numColorLines := strings.Split(setLine, ", ")
		for _, numColorLine := range numColorLines {
			numColor := strings.Fields(numColorLine)
			set[ColorFromString(numColor[1])] = util.MustAtoi(numColor[0])
		}
		sets = append(sets, set)
	}
	return Game{ID: id, Sets: sets}
}

func parseInput(filename string) []Game {
	input, _ := util.ReadInput(filename)
	games := make([]Game, 0, len(input))
	for _, line := range input {
		games = append(games, parseGame(line))
	}
	return games
}

func part1() {
	games := parseInput("input.txt")
	sum := 0
	bag := Set{Red: 12, Green: 13, Blue: 14}
	for _, game := range games {
		possible := true
		for _, set := range game.Sets {
			if set[Blue] > bag[Blue] || set[Red] > bag[Red] || set[Green] > bag[Green] {
				possible = false
				break
			}
		}
		if possible {
			sum += game.ID
		}
	}
	fmt.Printf("Part 1: %v\n", sum)
}

func part2() {
	games := parseInput("input.txt")
	sum := 0
	for _, game := range games {
		max := make(Set)
		for _, set := range game.Sets {
			if max[Blue] < set[Blue] {
				max[Blue] = set[Blue]
			}
			if max[Red] < set[Red] {
				max[Red] = set[Red]
			}
			if max[Green] < set[Green] {
				max[Green] = set[Green]
			}
		}
		sum += max[Red] * max[Green] * max[Blue]
	}
	fmt.Printf("Part 2: %v\n", sum)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
