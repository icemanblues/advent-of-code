package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "10"
	dayTitle = "Cathode-Ray Tube"
)

type Inst struct {
	cmd   string
	value int
	count int
}

func parse(filename string) []Inst {
	input, _ := util.ReadInput(filename)
	insts := make([]Inst, 0, len(input))
	for _, line := range input {
		fields := strings.Fields(line)
		if len(fields) == 1 {
			insts = append(insts, Inst{fields[0], 0, 1})
			continue
		}
		i, _ := strconv.Atoi(fields[1])
		insts = append(insts, Inst{fields[0], i, 2})
	}
	return insts
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	insts := parse("input.txt")
	idx, x, cycle, score := 0, 1, 0, 0
	for idx < len(insts) {
		cycle++
		insts[idx].count--
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			signal := cycle * x
			score += signal
		}
		if insts[idx].count == 0 {
			x += insts[idx].value
			idx++
		}
	}
	fmt.Printf("Part 1: %v\n", score)
}
