package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "02"
	dayTitle = "Rock Paper Scissors"
)

type RPS int

const (
	Rock    RPS = 1
	Paper   RPS = 2
	Scissor RPS = 3
)

type Inst struct {
	ElfCode, HumanCode rune
	ElfRps, HumanRps   RPS
}

var ElfMap map[rune]RPS = map[rune]RPS{
	'A': Rock,
	'B': Paper,
	'C': Scissor,
}

var HumanMap map[rune]RPS = map[rune]RPS{
	'X': Rock,
	'Y': Paper,
	'Z': Scissor,
}

func Play(insts []Inst) int {
	score := 0
	for _, inst := range insts {
		score += PlayRPS(inst.ElfRps, inst.HumanRps)
	}
	return score
}

func PlayRPS(elf, human RPS) int {
	if elf == human {
		return 3 + int(human)
	}

	if elf == Rock && human == Scissor || elf == Scissor && human == Paper ||
		elf == Paper && human == Rock { // elf wins
		return 0 + int(human)
	} else { // human wins
		return 6 + int(human)
	}
}

func humanPlay(elf RPS, outcome rune) RPS {
	switch outcome {
	case 'X': // lose
		switch elf {
		case Rock:
			return Scissor
		case Paper:
			return Rock
		case Scissor:
			return Paper
		}
	case 'Y': // draw
		switch elf {
		case Rock:
			return Rock
		case Paper:
			return Paper
		case Scissor:
			return Scissor
		}
	case 'Z': // win
		switch elf {
		case Rock:
			return Paper
		case Paper:
			return Scissor
		case Scissor:
			return Rock
		}
	}

	panic("Impossible!")
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)

	input, _ := util.ReadRuneput("input.txt")
	insts := make([]Inst, 0, len(input))
	for _, runes := range input {
		insts = append(insts, Inst{runes[0], runes[2], ElfMap[runes[0]], HumanMap[runes[2]]})
	}
	fmt.Printf("Part 1: %v\n", Play(insts))

	cheats := make([]Inst, 0, len(insts))
	for _, inst := range insts {
		move := humanPlay(inst.ElfRps, inst.HumanCode)
		cheats = append(cheats, Inst{inst.ElfCode, inst.HumanCode, inst.ElfRps, move})
	}
	fmt.Printf("Part 2: %v\n", Play(cheats))
}
