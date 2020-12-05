package main

import (
	"fmt"
	"math"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "05"
	dayTitle = "Title"
)

func seat(line string) (int, int) {
	row, col := -1, -1

	min, max := 0, 127
	minC, maxC := 0, 7

	for i, r := range line {
		// row
		if i == 6 {
			if r == 'F' {
				row = min
			} else {
				row = max
			}
			continue
		}
		if i < 7 {
			rowMid := float64(min+max) / 2.0
			if r == 'F' {
				max = int(math.Floor(rowMid))
			} else {
				min = int(math.Ceil(rowMid))
			}
			continue
		}

		if i == 9 {
			if r == 'L' {
				col = minC
			} else {
				col = maxC
			}
			continue
		}

		colMid := float64(minC+maxC) / 2.0
		if r == 'L' {
			maxC = int(math.Floor(colMid))
		} else {
			minC = int(math.Ceil(colMid))
		}

	}
	return row, col
}

func seatID(r, c int) int {
	return 8*r + c
}

func maxSeatID(boardingPass []string) int {
	max := -1
	for _, line := range boardingPass {
		row, col := seat(line)
		seatID := seatID(row, col)

		if seatID > max {
			max = seatID
		}
	}
	return max
}

func test1(bp string) {
	r, c := seat(bp)
	id := seatID(r, c)
	fmt.Printf("test r: %v c: %v id: %v\n", r, c, id)
}

type Seat struct {
	row, col, seatID int
}

func mySeat(lines []string) int {
	seats := make([]Seat, 0, len(lines))
	seatSet := make(map[Seat]struct{})
	ids := make([]int, 0, len(lines))
	idSet := make(map[int]struct{})
	for _, line := range lines {
		row, col := seat(line)
		seatID := seatID(row, col)
		s := Seat{row, col, seatID}
		seats = append(seats, s)
		seatSet[s] = struct{}{}
		ids = append(ids, seatID)
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
