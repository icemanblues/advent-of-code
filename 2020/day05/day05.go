package main

import (
	"fmt"
	"math"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "05"
	dayTitle = "Binary Boarding"
)

func parseSeat(line string) int {
	row, col := -1, -1
	minR, maxR := 0, 127
	minC, maxC := 0, 7

	for _, r := range line {
		rowMid := float64(minR+maxR) / 2.0
		colMid := float64(minC+maxC) / 2.0
		switch r {
		case 'F':
			maxR = int(math.Floor(rowMid))
			row = minR
		case 'B':
			minR = int(math.Ceil(rowMid))
			row = maxR
		case 'L':
			maxC = int(math.Floor(colMid))
			col = minC
		case 'R':
			minC = int(math.Ceil(colMid))
			col = maxC
		default:
			fmt.Printf("Unknown rune in boarding pass: %v\n", r)
		}
	}

	return seatID(row, col)
}

func seatID(r, c int) int {
	return 8*r + c
}

func maxSeatID(boardingPass []string) int {
	max := -1
	for _, line := range boardingPass {
		seatID := parseSeat(line)
		if seatID > max {
			max = seatID
		}
	}
	return max
}

func mySeat(lines []string) int {
	idSet := make(map[int]struct{})
	for _, line := range lines {
		seatID := parseSeat(line)
		idSet[seatID] = struct{}{}
	}

	for r := 0; r <= 127; r++ {
		for c := 0; c <= 7; c++ {
			id := seatID(r, c)
			if _, ok := idSet[id]; !ok {
				_, bOk := idSet[id-1]
				_, aOK := idSet[id+1]
				if bOk && aOK {
					return id
				}
			}
		}
	}
	return -1
}

func part1() {
	lines, _ := util.ReadInput("input.txt")
	fmt.Printf("Part 1: %v\n", maxSeatID(lines))
}

func part2() {
	lines, _ := util.ReadInput("input.txt")
	fmt.Printf("Part 2: %v\n", mySeat(lines))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
