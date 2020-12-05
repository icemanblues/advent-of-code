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

	for i, r := range line {
		if i < 6 { // row search
			rowMid := float64(minR+maxR) / 2.0
			if r == 'F' {
				maxR = int(math.Floor(rowMid))
			} else {
				minR = int(math.Ceil(rowMid))
			}
		} else if i == 6 { // row search is ending
			if r == 'F' {
				row = minR
			} else {
				row = maxR
			}
		} else if i < 9 { // col search
			colMid := float64(minC+maxC) / 2.0
			if r == 'L' {
				maxC = int(math.Floor(colMid))
			} else {
				minC = int(math.Ceil(colMid))
			}
		} else { // col search is ending
			if r == 'L' {
				col = minC
			} else {
				col = maxC
			}
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
			if _, ok := idSet[id]; !ok { // not there
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
