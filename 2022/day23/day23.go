package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "23"
	dayTitle = "Unstable Diffusion"
)

func Parse(filename string) []util.Point {
	m, _ := util.ReadSparseGrid(filename, '#')
	elves := make([]util.Point, 0, len(m))
	for k := range m {
		elves = append(elves, k)
	}
	return elves
}

func Propose(elf util.Point, sparse map[util.Point]struct{}, rnd int) util.Point {
	n := util.NewPoint(elf.X, elf.Y-1)
	ne := util.NewPoint(elf.X+1, elf.Y-1)
	nw := util.NewPoint(elf.X-1, elf.Y-1)
	s := util.NewPoint(elf.X, elf.Y+1)
	se := util.NewPoint(elf.X+1, elf.Y+1)
	sw := util.NewPoint(elf.X-1, elf.Y+1)
	e := util.NewPoint(elf.X+1, elf.Y)
	w := util.NewPoint(elf.X-1, elf.Y)

	_, nOk := sparse[n]
	_, neOk := sparse[ne]
	_, nwOk := sparse[nw]
	_, sOk := sparse[s]
	_, seOk := sparse[se]
	_, swOk := sparse[sw]
	_, eOk := sparse[e]
	_, wOk := sparse[w]

	if !nOk && !neOk && !nwOk && !sOk && !seOk && !swOk && !eOk && !wOk {
		return elf
	}

	// N, S, W, E
	north := !nOk && !neOk && !nwOk
	south := !sOk && !seOk && !swOk
	west := !wOk && !nwOk && !swOk
	east := !eOk && !neOk && !seOk

	decision := []bool{north, south, west, east}
	step := []util.Point{n, s, w, e}
	for i := 0; i < len(step); i++ {
		if j := (i + rnd) % len(step); decision[j] {
			return step[j]
		}
	}

	// can't move?
	return elf
}

func Score(elves []util.Point) int {
	minX, maxX, minY, maxY := 0, 0, 0, 0
	for _, p := range elves {
		if p.X > maxX {
			maxX = p.X
		}
		if p.X < minX {
			minX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
		if p.Y < minY {
			minY = p.Y
		}
	}
	l, w := maxX-minX+1, maxY-minY+1
	return l*w - len(elves)
}

func Round(round int, elves []util.Point) []util.Point {
	// sparse
	sparse := make(map[util.Point]struct{})
	for _, elf := range elves {
		sparse[elf] = struct{}{}
	}

	// each elf needs to propose a move
	proposed := make([]util.Point, 0, len(elves))
	for _, elf := range elves {
		proposed = append(proposed, Propose(elf, sparse, round))
	}

	// check for conflicts
	counts := make(map[util.Point]int)
	for _, p := range proposed {
		counts[p]++
	}

	// update sparse and elves
	next := make([]util.Point, len(elves), len(elves))
	for i, p := range proposed {
		if counts[p] == 1 {
			next[i] = p
		} else {
			next[i] = elves[i]
		}
	}
	return next
}

func part1() {
	elves := Parse("input.txt")
	for round := 0; round < 10; round++ {
		elves = Round(round, elves)
	}
	fmt.Printf("Part 1: %v\n", Score(elves))
}

func part2() {
	elves := Parse("input.txt")
	for round := 0; true; round++ {
		temp := Round(round, elves)

		noMove := true
		for i, e := range temp {
			if e != elves[i] {
				noMove = false
				break
			}
		}
		elves = temp

		if noMove {
			fmt.Printf("Part 2: %v\n", round+1)
			break
		}
	}
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
