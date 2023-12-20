package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "14"
	dayTitle = "Parabolic Reflector Dish"
)

type Rocks map[util.Point]struct{}

func parse(filename string) (Rocks, Rocks, int, int) {
	rolling, _ := util.ReadSparseGrid(filename, 'O')
	stable, _ := util.ReadSparseGrid(filename, '#')

	lines, _ := util.ReadInput(filename)
	yMax, xMax := len(lines), len(lines[0])
	return rolling, stable, xMax, yMax
}

var North = util.NewPoint(0, -1)
var West = util.NewPoint(-1, 0)
var South = util.NewPoint(0, 1)
var East = util.NewPoint(1, 0)

func tilt(rolling, stable Rocks, dir util.Point, xMax, yMax int) Rocks {
	next := make(Rocks)
	for p := range rolling {
		new := util.NewPoint(p.X+dir.X, p.Y+dir.Y)
		if new.X < 0 || new.X >= xMax || new.Y < 0 || new.Y >= yMax {
			new = p
		}
		if _, ok := stable[new]; new.Y < 0 || ok {
			new = p
		} else if _, ok := rolling[new]; ok {
			new = p
		}
		next[new] = struct{}{}
	}
	return next
}

func isSame(a, b Rocks) bool {
	for r := range a {
		if _, ok := b[r]; !ok {
			return false
		}
	}
	return true
}

func fullTilt(rolling, stable Rocks, dir util.Point, xMax, yMax int) Rocks {
	for true {
		next := tilt(rolling, stable, dir, xMax, yMax)
		if isSame(rolling, next) {
			break
		}
		rolling = next
	}
	return rolling
}

func spinCycle(rolling, stable Rocks, xMax, yMax int) Rocks {
	for _, dir := range []util.Point{North, West, South, East} {
		rolling = fullTilt(rolling, stable, dir, xMax, yMax)
	}
	return rolling
}

func PrintGrid(rolling, stable Rocks, xMax, yMax int) {
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			p := util.NewPoint(x, y)
			if _, ok := rolling[p]; ok {
				fmt.Printf("O")
			} else if _, ok := stable[p]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func load(rolling Rocks, yMax int) (sum int) {
	for r := range rolling {
		sum += yMax - r.Y
	}
	return
}

func part1() {
	rolling, stable, xMax, yMax := parse("input.txt")
	rolling = fullTilt(rolling, stable, North, xMax, yMax)
	fmt.Printf("Part 1: %v\n", load(rolling, yMax))
}

func part2() {
	rolling, stable, xMax, yMax := parse("input.txt")
	cycle := 1000000000

	var stateList = make([]Rocks, 0)
	cycleOffset, iter := 0, 0
search:
	for true {
		// sequence search stinks! A hashmap lookup would be faster here
		for idx, rocks := range stateList {
			if isSame(rocks, rolling) {
				cycleOffset = idx
				stateList = append(stateList, rolling)
				break search
			}
		}

		stateList = append(stateList, rolling)
		rolling = spinCycle(rolling, stable, xMax, yMax)
		iter++
	}
	cycleLength := iter - cycleOffset
	idx := (cycle-cycleOffset)%cycleLength + cycleOffset
	score := load(stateList[idx], yMax)
	fmt.Printf("Part 2: %v\n", score)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
