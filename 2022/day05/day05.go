package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "05"
	dayTitle = "Supply Stacks"
)

type Move struct {
	Quantity, From, To int
}

func parse(filename string) ([][]rune, []Move) {
	var stackLines, moveLines []string
	var numLine string

	input, _ := util.ReadInput(filename)
	isStack := true
	for _, line := range input {
		if line == "" {
			isStack = false
			numLine = stackLines[len(stackLines)-1]
			stackLines = stackLines[:len(stackLines)-1]
			continue
		}

		if isStack {
			stackLines = append(stackLines, line)
			continue
		}

		moveLines = append(moveLines, line)
	}

	// build our stacks
	var indices []int
	for i, r := range numLine {
		if r != ' ' {
			indices = append(indices, i)
		}
	}

	stacks := make([][]rune, len(indices), len(indices))
	for j, idx := range indices {
		for i := len(stackLines) - 1; i >= 0; i-- {
			r := []rune(stackLines[i])[idx]
			if r == ' ' {
				continue
			}
			stacks[j] = append(stacks[j], r)
		}
	}

	// build our move list
	moves := make([]Move, 0, len(moveLines))
	for _, e := range moveLines {
		fields := strings.Fields(e)
		q, _ := strconv.Atoi(fields[1])
		f, _ := strconv.Atoi(fields[3])
		t, _ := strconv.Atoi(fields[5])
		moves = append(moves, Move{q, f - 1, t - 1}) // 0-indexed
	}

	return stacks, moves
}

func STop(stacks [][]rune) string {
	msg := make([]rune, len(stacks), len(stacks))
	for i, stack := range stacks {
		msg[i] = stack[len(stack)-1]
	}
	return string(msg)
}

func part1() {
	stacks, moves := parse("input.txt")
	for _, move := range moves {
		for i := 0; i < move.Quantity; i++ {
			block := stacks[move.From][len(stacks[move.From])-1]
			stacks[move.From] = stacks[move.From][:len(stacks[move.From])-1]
			stacks[move.To] = append(stacks[move.To], block)
		}
	}
	fmt.Printf("Part 1: %v\n", STop(stacks))
}

func part2() {
	stacks, moves := parse("input.txt")
	for _, move := range moves {
		blocks := stacks[move.From][len(stacks[move.From])-move.Quantity : len(stacks[move.From])]
		stacks[move.From] = stacks[move.From][:len(stacks[move.From])-move.Quantity]
		stacks[move.To] = append(stacks[move.To], blocks...)

	}
	fmt.Printf("Part 2: %v\n", STop(stacks))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
